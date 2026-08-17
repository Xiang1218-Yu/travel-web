# 缺陷复现说明

## 缺陷现象
地图上没有显示一个已填写地点名称、但经纬度恰好是 0,0 的旅行计划。文件先不要改，帮我定位为什么这个位置会被静默过滤，并说明同类影响范围。这次只需要定位结论，不用修改或执行测试。

## 触发方式
# 1. 进入项目目录；预期：在旅行服务基线中复现具名 0,0 坐标计划缺失地图标记。
cd /workplace/travel-web__005
# 2. 连续执行目标用例；预期：基线稳定失败，返回列表中没有该计划标记。
go test -v -run TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation -count=20
# 3. 检查整体回归；预期：修复版本完整测试通过。
go test ./...

## 触发后的实际错误输出

```text
=== RUN   TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation
    marker_location_test.go:23: 具名的 0,0 坐标旅行没有出现在地图标记中: []main.Marker{}
--- FAIL: TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation (0.00s)
=== RUN   TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation
    marker_location_test.go:23: 具名的 0,0 坐标旅行没有出现在地图标记中: []main.Marker{}
--- FAIL: TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation (0.00s)
=== RUN   TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation
    marker_location_test.go:23: 具名的 0,0 坐标旅行没有出现在地图标记中: []main.Marker{}
--- FAIL: TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation (0.00s)
=== RUN   TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation
    marker_location_test.go:23: 具名的 0,0 坐标旅行没有出现在地图标记中: []main.Marker{}
--- FAIL: TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation (0.00s)
=== RUN   TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation
    marker_location_test.go:23: 具名的 0,0 坐标旅行没有出现在地图标记中: []main.Marker{}
--- FAIL: TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation (0.00s)
=== RUN   TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation
    marker_location_test.go:23: 具名的 0,0 坐标旅行没有出现在地图标记中: []main.Marker{}
--- FAIL: TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation (0.00s)
=== RUN   TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation
    marker_location_test.go:23: 具名的 0,0 坐标旅行没有出现在地图标记中: []main.Marker{}
--- FAIL: TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation (0.00s)
=== RUN   TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation
    marker_location_test.go:23: 具名的 0,0 坐标旅行没有出现在地图标记中: []main.Marker{}
--- FAIL: TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation (0.00s)
=== RUN   TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation
    marker_location_test.go:23: 具名的 0,0 坐标旅行没有出现在地图标记中: []main.Marker{}
--- FAIL: TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation (0.00s)
=== RUN   TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation
    marker_location_test.go:23: 具名的 0,0 坐标旅行没有出现在地图标记中: []main.Marker{}
--- FAIL: TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation (0.00s)
=== RUN   TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation
    marker_location_test.go:23: 具名的 0,0 坐标旅行没有出现在地图标记中: []main.Marker{}
--- FAIL: TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation (0.00s)
=== RUN   TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation
    marker_location_test.go:23: 具名的 0,0 坐标旅行没有出现在地图标记中: []main.Marker{}
--- FAIL: TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation (0.00s)
=== RUN   TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation
    marker_location_test.go:23: 具名的 0,0 坐标旅行没有出现在地图标记中: []main.Marker{}
--- FAIL: TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation (0.00s)
=== RUN   TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation
    marker_location_test.go:23: 具名的 0,0 坐标旅行没有出现在地图标记中: []main.Marker{}
--- FAIL: TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation (0.00s)
=== RUN   TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation
    marker_location_test.go:23: 具名的 0,0 坐标旅行没有出现在地图标记中: []main.Marker{}
--- FAIL: TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation (0.00s)
=== RUN   TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation
    marker_location_test.go:23: 具名的 0,0 坐标旅行没有出现在地图标记中: []main.Marker{}
--- FAIL: TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation (0.00s)
=== RUN   TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation
    marker_location_test.go:23: 具名的 0,0 坐标旅行没有出现在地图标记中: []main.Marker{}
--- FAIL: TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation (0.00s)
=== RUN   TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation
    marker_location_test.go:23: 具名的 0,0 坐标旅行没有出现在地图标记中: []main.Marker{}
--- FAIL: TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation (0.00s)
=== RUN   TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation
    marker_location_test.go:23: 具名的 0,0 坐标旅行没有出现在地图标记中: []main.Marker{}
--- FAIL: TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation (0.00s)
=== RUN   TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation
    marker_location_test.go:23: 具名的 0,0 坐标旅行没有出现在地图标记中: []main.Marker{}
--- FAIL: TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation (0.00s)
FAIL
exit status 1
FAIL	travel-web	0.464s
```
