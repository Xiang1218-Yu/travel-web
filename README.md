# travel-web

一个使用 Go 标准库实现的轻量旅行计划与游记服务。它提供旅行计划、游记、图片上传与地图标记点的 JSON API，并将数据保存在本地 JSON 文件中。

## 运行

```bash
go run .
```

默认监听 `http://localhost:8080`。静态页面位于 `static/`，上传的图片位于 `uploads/`，数据文件默认为 `data/db.json`。

## API 概览

- `GET` / `POST` `/api/plans`
- `GET` / `PUT` / `DELETE` `/api/plans/{id}`
- `POST` `/api/plans/{id}/privacy`
- `GET` / `POST` `/api/diaries`
- `GET` / `PUT` / `DELETE` `/api/diaries/{id}`
- `POST` `/api/upload`
- `GET` `/api/markers`

计划列表支持 `destination`、`date` 和 `public=true` 筛选；游记列表支持 `plan_id` 和 `date` 筛选；地图标记支持 `public=true` 筛选。

## 开发与验证

项目需要 Go 1.26.5 或兼容工具链：

```bash
go build ./...
go test ./...
```

## 目录结构

```text
main.go       HTTP 服务与本地 JSON 数据读写实现
static/       浏览器页面
uploads/      已上传图片
 data/        本地 JSON 数据
```
