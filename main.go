package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

// ==================== 数据模型 ====================

type Location struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type ItineraryItem struct {
	Time        string `json:"time"`
	Activity    string `json:"activity"`
	Description string `json:"description"`
}

type TravelPlan struct {
	ID          string          `json:"id"`
	Title       string          `json:"title"`
	Destination string          `json:"destination"`
	StartDate   string          `json:"start_date"`
	EndDate     string          `json:"end_date"`
	Location    Location        `json:"location"`
	Itinerary   []ItineraryItem `json:"itinerary"`
	IsPublic    bool            `json:"is_public"`
	CreatedAt   string          `json:"created_at"`
	UpdatedAt   string          `json:"updated_at"`
}

type DiaryEntry struct {
	ID        string   `json:"id"`
	PlanID    string   `json:"plan_id"`
	Date      string   `json:"date"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Images    []string `json:"images"`
	Location  Location `json:"location"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type Database struct {
	Plans   []TravelPlan `json:"plans"`
	Diaries []DiaryEntry `json:"diaries"`
}

// ==================== 全局变量 ====================

var (
	db        Database
	dbMutex   sync.RWMutex
	dataPath  = "data/db.json"
	uploadDir = "uploads"
)

// ==================== 工具函数 ====================

func generateID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func loadDB() error {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		db = Database{Plans: []TravelPlan{}, Diaries: []DiaryEntry{}}
		return saveDBLocked()
	}

	data, err := os.ReadFile(dataPath)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &db)
}

func saveDBLocked() error {
	data, err := json.MarshalIndent(db, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataPath, data, 0644)
}

func saveDB() error {
	dbMutex.Lock()
	defer dbMutex.Unlock()
	return saveDBLocked()
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

func parseJSON(r *http.Request, v interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.Unmarshal(body, v)
}

// ==================== 旅行计划 API ====================

func createPlan(w http.ResponseWriter, r *http.Request) {
	var plan TravelPlan
	if err := parseJSON(r, &plan); err != nil {
		writeError(w, http.StatusBadRequest, "无效的请求数据")
		return
	}

	if plan.Title == "" || plan.Destination == "" || plan.StartDate == "" {
		writeError(w, http.StatusBadRequest, "标题、目的地和开始日期为必填项")
		return
	}

	plan.ID = generateID()
	plan.CreatedAt = time.Now().Format(time.RFC3339)
	plan.UpdatedAt = plan.CreatedAt
	if plan.Itinerary == nil {
		plan.Itinerary = []ItineraryItem{}
	}

	dbMutex.Lock()
	db.Plans = append(db.Plans, plan)
	saveDBLocked()
	dbMutex.Unlock()

	writeJSON(w, http.StatusCreated, plan)
}

func listPlans(w http.ResponseWriter, r *http.Request) {
	dbMutex.RLock()
	defer dbMutex.RUnlock()

	plans := make([]TravelPlan, len(db.Plans))
	copy(plans, db.Plans)

	// 搜索筛选
	dest := strings.ToLower(r.URL.Query().Get("destination"))
	date := r.URL.Query().Get("date")
	onlyPublic := r.URL.Query().Get("public") == "true"

	filtered := []TravelPlan{}
	for _, p := range plans {
		if onlyPublic && !p.IsPublic {
			continue
		}
		if dest != "" && !strings.Contains(strings.ToLower(p.Destination), dest) {
			continue
		}
		if date != "" {
			if p.StartDate > date || (p.EndDate != "" && p.EndDate < date) {
				continue
			}
			if p.EndDate == "" && p.StartDate != date {
				continue
			}
		}
		filtered = append(filtered, p)
	}

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].CreatedAt > filtered[j].CreatedAt
	})

	writeJSON(w, http.StatusOK, filtered)
}

func getPlan(w http.ResponseWriter, r *http.Request, id string) {
	dbMutex.RLock()
	defer dbMutex.RUnlock()

	for _, p := range db.Plans {
		if p.ID == id {
			writeJSON(w, http.StatusOK, p)
			return
		}
	}
	writeError(w, http.StatusNotFound, "旅行计划不存在")
}

func updatePlan(w http.ResponseWriter, r *http.Request, id string) {
	patch, err := decodePlanPatch(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "无效的请求数据")
		return
	}
	dbMutex.Lock()
	defer dbMutex.Unlock()
	for i := range db.Plans {
		if db.Plans[i].ID != id {
			continue
		}
		db.Plans[i] = applyPlanPatch(db.Plans[i], patch, time.Now())
		if err := saveDBLocked(); err != nil {
			writeError(w, http.StatusInternalServerError, "保存旅行计划失败")
			return
		}
		writeJSON(w, http.StatusOK, db.Plans[i])
		return
	}
	writeError(w, http.StatusNotFound, "旅行计划不存在")
}

func togglePlanPrivacy(w http.ResponseWriter, r *http.Request, id string) {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	for i, p := range db.Plans {
		if p.ID == id {
			db.Plans[i].IsPublic = !p.IsPublic
			db.Plans[i].UpdatedAt = time.Now().Format(time.RFC3339)
			saveDBLocked()
			writeJSON(w, http.StatusOK, db.Plans[i])
			return
		}
	}
	writeError(w, http.StatusNotFound, "旅行计划不存在")
}

func deletePlan(w http.ResponseWriter, r *http.Request, id string) {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	planIdx := -1
	for i, p := range db.Plans {
		if p.ID == id {
			planIdx = i
			break
		}
	}
	if planIdx == -1 {
		writeError(w, http.StatusNotFound, "旅行计划不存在")
		return
	}

	// 删除关联的游记和图片
	remainingDiaries := []DiaryEntry{}
	for _, d := range db.Diaries {
		if d.PlanID == id {
			for _, img := range d.Images {
				os.Remove(filepath.Join(uploadDir, img))
			}
		} else {
			remainingDiaries = append(remainingDiaries, d)
		}
	}
	db.Diaries = remainingDiaries
	db.Plans = append(db.Plans[:planIdx], db.Plans[planIdx+1:]...)
	saveDBLocked()

	writeJSON(w, http.StatusOK, map[string]string{"message": "删除成功"})
}

// ==================== 游记 API ====================

func createDiary(w http.ResponseWriter, r *http.Request) {
	var diary DiaryEntry
	if err := parseJSON(r, &diary); err != nil {
		writeError(w, http.StatusBadRequest, "无效的请求数据")
		return
	}

	if diary.PlanID == "" || diary.Date == "" || diary.Title == "" {
		writeError(w, http.StatusBadRequest, "计划ID、日期和标题为必填项")
		return
	}

	diary.ID = generateID()
	diary.CreatedAt = time.Now().Format(time.RFC3339)
	diary.UpdatedAt = diary.CreatedAt
	if diary.Images == nil {
		diary.Images = []string{}
	}

	dbMutex.Lock()
	db.Diaries = append(db.Diaries, diary)
	saveDBLocked()
	dbMutex.Unlock()

	writeJSON(w, http.StatusCreated, diary)
}

func listDiaries(w http.ResponseWriter, r *http.Request) {
	dbMutex.RLock()
	defer dbMutex.RUnlock()

	planID := r.URL.Query().Get("plan_id")
	date := r.URL.Query().Get("date")

	diaries := make([]DiaryEntry, 0)
	for _, d := range db.Diaries {
		if planID != "" && d.PlanID != planID {
			continue
		}
		if date != "" && !strings.HasPrefix(d.Date, date) {
			continue
		}
		diaries = append(diaries, d)
	}

	sort.Slice(diaries, func(i, j int) bool {
		return diaries[i].Date > diaries[j].Date
	})

	writeJSON(w, http.StatusOK, diaries)
}

func getDiary(w http.ResponseWriter, r *http.Request, id string) {
	dbMutex.RLock()
	defer dbMutex.RUnlock()

	for _, d := range db.Diaries {
		if d.ID == id {
			writeJSON(w, http.StatusOK, d)
			return
		}
	}
	writeError(w, http.StatusNotFound, "游记不存在")
}

func updateDiary(w http.ResponseWriter, r *http.Request, id string) {
	var updates DiaryEntry
	if err := parseJSON(r, &updates); err != nil {
		writeError(w, http.StatusBadRequest, "无效的请求数据")
		return
	}

	dbMutex.Lock()
	defer dbMutex.Unlock()

	for i, d := range db.Diaries {
		if d.ID == id {
			if updates.Date != "" {
				db.Diaries[i].Date = updates.Date
			}
			if updates.Title != "" {
				db.Diaries[i].Title = updates.Title
			}
			if updates.Content != "" {
				db.Diaries[i].Content = updates.Content
			}
			if updates.Images != nil {
				db.Diaries[i].Images = updates.Images
			}
			if updates.Location.Name != "" || updates.Location.Latitude != 0 || updates.Location.Longitude != 0 {
				db.Diaries[i].Location = updates.Location
			}
			db.Diaries[i].UpdatedAt = time.Now().Format(time.RFC3339)
			saveDBLocked()
			writeJSON(w, http.StatusOK, db.Diaries[i])
			return
		}
	}
	writeError(w, http.StatusNotFound, "游记不存在")
}

func deleteDiary(w http.ResponseWriter, r *http.Request, id string) {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	for i, d := range db.Diaries {
		if d.ID == id {
			for _, img := range d.Images {
				os.Remove(filepath.Join(uploadDir, img))
			}
			db.Diaries = append(db.Diaries[:i], db.Diaries[i+1:]...)
			saveDBLocked()
			writeJSON(w, http.StatusOK, map[string]string{"message": "删除成功"})
			return
		}
	}
	writeError(w, http.StatusNotFound, "游记不存在")
}

// ==================== 图片上传 API ====================

func uploadImage(w http.ResponseWriter, r *http.Request) {
	const maxUpload = 50 << 20
	if r.ContentLength > maxUpload {
		writeError(w, http.StatusBadRequest, "文件过大，最大50MB")
		return
	}
	if err := r.ParseMultipartForm(maxUpload); err != nil {
		writeError(w, http.StatusBadRequest, "解析上传表单失败: "+err.Error())
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		writeError(w, http.StatusBadRequest, "获取上传文件失败(字段名应为image): "+err.Error())
		return
	}
	defer file.Close()

	ext := strings.ToLower(filepath.Ext(header.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" && ext != ".webp" {
		writeError(w, http.StatusBadRequest, "仅支持 jpg/jpeg/png/gif/webp 格式")
		return
	}

	filename := generateID() + ext
	savePath := filepath.Join(uploadDir, filename)

	out, err := os.Create(savePath)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "保存文件失败")
		return
	}
	defer out.Close()

	if _, err := io.Copy(out, file); err != nil {
		writeError(w, http.StatusInternalServerError, "保存文件失败")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"filename": filename,
		"url":      "/uploads/" + filename,
	})
}

// ==================== 地图标记点 API ====================

func getMapMarkers(w http.ResponseWriter, r *http.Request) {
	dbMutex.RLock()
	defer dbMutex.RUnlock()

	onlyPublic := r.URL.Query().Get("public") == "true"
	type Marker struct {
		ID        string  `json:"id"`
		Type      string  `json:"type"`
		PlanID    string  `json:"plan_id"`
		Title     string  `json:"title"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Date      string  `json:"date"`
	}

	markers := []Marker{}
	for _, p := range db.Plans {
		if onlyPublic && !p.IsPublic {
			continue
		}
		if p.Location.Latitude != 0 || p.Location.Longitude != 0 {
			markers = append(markers, Marker{
				ID:        "plan-" + p.ID,
				Type:      "plan",
				PlanID:    p.ID,
				Title:     p.Title,
				Latitude:  p.Location.Latitude,
				Longitude: p.Location.Longitude,
				Date:      p.StartDate,
			})
		}
	}

	planPublicMap := map[string]bool{}
	for _, p := range db.Plans {
		planPublicMap[p.ID] = p.IsPublic
	}

	for _, d := range db.Diaries {
		if onlyPublic && !planPublicMap[d.PlanID] {
			continue
		}
		if d.Location.Latitude != 0 || d.Location.Longitude != 0 {
			markers = append(markers, Marker{
				ID:        "diary-" + d.ID,
				Type:      "diary",
				PlanID:    d.PlanID,
				Title:     d.Title,
				Latitude:  d.Location.Latitude,
				Longitude: d.Location.Longitude,
				Date:      d.Date,
			})
		}
	}

	writeJSON(w, http.StatusOK, markers)
}

// ==================== 路由分发 ====================

func handleAPI(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api")
	path = strings.Trim(path, "/")
	parts := strings.Split(path, "/")

	if len(parts) == 0 {
		writeError(w, http.StatusNotFound, "API不存在")
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	switch parts[0] {
	case "plans":
		switch {
		case len(parts) == 1 && r.Method == "POST":
			createPlan(w, r)
		case len(parts) == 1 && r.Method == "GET":
			listPlans(w, r)
		case len(parts) == 2 && r.Method == "GET":
			getPlan(w, r, parts[1])
		case len(parts) == 2 && r.Method == "PUT":
			updatePlan(w, r, parts[1])
		case len(parts) == 2 && r.Method == "DELETE":
			deletePlan(w, r, parts[1])
		case len(parts) == 3 && parts[2] == "privacy" && r.Method == "POST":
			togglePlanPrivacy(w, r, parts[1])
		default:
			writeError(w, http.StatusMethodNotAllowed, "方法不支持")
		}

	case "diaries":
		switch {
		case len(parts) == 1 && r.Method == "POST":
			createDiary(w, r)
		case len(parts) == 1 && r.Method == "GET":
			listDiaries(w, r)
		case len(parts) == 2 && r.Method == "GET":
			getDiary(w, r, parts[1])
		case len(parts) == 2 && r.Method == "PUT":
			updateDiary(w, r, parts[1])
		case len(parts) == 2 && r.Method == "DELETE":
			deleteDiary(w, r, parts[1])
		default:
			writeError(w, http.StatusMethodNotAllowed, "方法不支持")
		}

	case "upload":
		if r.Method == "POST" {
			uploadImage(w, r)
		} else {
			writeError(w, http.StatusMethodNotAllowed, "方法不支持")
		}

	case "markers":
		if r.Method == "GET" {
			getMapMarkers(w, r)
		} else {
			writeError(w, http.StatusMethodNotAllowed, "方法不支持")
		}

	default:
		writeError(w, http.StatusNotFound, "API不存在")
	}
}

// ==================== 主函数 ====================

func main() {
	if err := os.MkdirAll(filepath.Dir(dataPath), 0755); err != nil {
		log.Fatal("创建数据目录失败:", err)
	}
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		log.Fatal("创建上传目录失败:", err)
	}

	if err := loadDB(); err != nil {
		log.Fatal("加载数据失败:", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/api/", handleAPI)
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir(uploadDir))))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "static/index.html")
		} else if r.URL.Path == "/plans" {
			http.ServeFile(w, r, "static/plans.html")
		} else if r.URL.Path == "/diary" {
			http.ServeFile(w, r, "static/diary.html")
		} else if r.URL.Path == "/map" {
			http.ServeFile(w, r, "static/map.html")
		} else {
			http.NotFound(w, r)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}
	addr := host + ":" + port
	fmt.Println("🚀 旅行日记服务启动成功!")
	fmt.Println("📁 数据存储:", dataPath)
	fmt.Println("🖼️  图片存储:", uploadDir)
	fmt.Println("🌐 访问地址: http://localhost:" + port)
	fmt.Println("📋 首页:     http://localhost:" + port + "/")
	fmt.Println("🗺️  地图:     http://localhost:" + port + "/map")
	fmt.Println("🔗 监听地址:", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal("服务启动失败:", err)
	}
}
