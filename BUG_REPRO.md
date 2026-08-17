# 缺陷复现说明

## 缺陷现象
公开旅行计划只修改标题后会从公开地图中消失。请修复更新接口，保证未提交公开状态字段时保留原有公开状态，并确保显式设置公开或私密仍然生效。请实现修复并运行相关测试。

## 触发方式
# 1. 进入项目目录；预期：在含缺陷基线中复现只改标题导致公开状态丢失。
cd /workplace/travel-web__002
# 2. 连续执行目标用例；预期：基线稳定失败，公开计划被意外改为私密。
go test -v -run TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges -count=20
# 3. 执行完整回归；预期：修复后没有新增失败。
go test ./...

## 触发后的实际错误输出

```text
=== RUN   TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges
    plan_visibility_test.go:28: 只更新标题后公开状态被改成 false: main.TravelPlan{ID:"plan-public", Title:"新标题", Destination:"厦门", StartDate:"2026-09-01", EndDate:"", Location:main.Location{Name:"", Latitude:0, Longitude:0}, Itinerary:[]main.ItineraryItem(nil), IsPublic:false, CreatedAt:"", UpdatedAt:"2026-08-17T02:14:31Z"}
--- FAIL: TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges (0.01s)
=== RUN   TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges
    plan_visibility_test.go:28: 只更新标题后公开状态被改成 false: main.TravelPlan{ID:"plan-public", Title:"新标题", Destination:"厦门", StartDate:"2026-09-01", EndDate:"", Location:main.Location{Name:"", Latitude:0, Longitude:0}, Itinerary:[]main.ItineraryItem(nil), IsPublic:false, CreatedAt:"", UpdatedAt:"2026-08-17T02:14:31Z"}
--- FAIL: TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges (0.00s)
=== RUN   TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges
    plan_visibility_test.go:28: 只更新标题后公开状态被改成 false: main.TravelPlan{ID:"plan-public", Title:"新标题", Destination:"厦门", StartDate:"2026-09-01", EndDate:"", Location:main.Location{Name:"", Latitude:0, Longitude:0}, Itinerary:[]main.ItineraryItem(nil), IsPublic:false, CreatedAt:"", UpdatedAt:"2026-08-17T02:14:31Z"}
--- FAIL: TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges (0.00s)
=== RUN   TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges
    plan_visibility_test.go:28: 只更新标题后公开状态被改成 false: main.TravelPlan{ID:"plan-public", Title:"新标题", Destination:"厦门", StartDate:"2026-09-01", EndDate:"", Location:main.Location{Name:"", Latitude:0, Longitude:0}, Itinerary:[]main.ItineraryItem(nil), IsPublic:false, CreatedAt:"", UpdatedAt:"2026-08-17T02:14:31Z"}
--- FAIL: TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges (0.00s)
=== RUN   TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges
    plan_visibility_test.go:28: 只更新标题后公开状态被改成 false: main.TravelPlan{ID:"plan-public", Title:"新标题", Destination:"厦门", StartDate:"2026-09-01", EndDate:"", Location:main.Location{Name:"", Latitude:0, Longitude:0}, Itinerary:[]main.ItineraryItem(nil), IsPublic:false, CreatedAt:"", UpdatedAt:"2026-08-17T02:14:31Z"}
--- FAIL: TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges (0.00s)
=== RUN   TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges
    plan_visibility_test.go:28: 只更新标题后公开状态被改成 false: main.TravelPlan{ID:"plan-public", Title:"新标题", Destination:"厦门", StartDate:"2026-09-01", EndDate:"", Location:main.Location{Name:"", Latitude:0, Longitude:0}, Itinerary:[]main.ItineraryItem(nil), IsPublic:false, CreatedAt:"", UpdatedAt:"2026-08-17T02:14:31Z"}
--- FAIL: TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges (0.00s)
=== RUN   TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges
    plan_visibility_test.go:28: 只更新标题后公开状态被改成 false: main.TravelPlan{ID:"plan-public", Title:"新标题", Destination:"厦门", StartDate:"2026-09-01", EndDate:"", Location:main.Location{Name:"", Latitude:0, Longitude:0}, Itinerary:[]main.ItineraryItem(nil), IsPublic:false, CreatedAt:"", UpdatedAt:"2026-08-17T02:14:31Z"}
--- FAIL: TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges (0.00s)
=== RUN   TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges
    plan_visibility_test.go:28: 只更新标题后公开状态被改成 false: main.TravelPlan{ID:"plan-public", Title:"新标题", Destination:"厦门", StartDate:"2026-09-01", EndDate:"", Location:main.Location{Name:"", Latitude:0, Longitude:0}, Itinerary:[]main.ItineraryItem(nil), IsPublic:false, CreatedAt:"", UpdatedAt:"2026-08-17T02:14:31Z"}
--- FAIL: TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges (0.00s)
=== RUN   TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges
    plan_visibility_test.go:28: 只更新标题后公开状态被改成 false: main.TravelPlan{ID:"plan-public", Title:"新标题", Destination:"厦门", StartDate:"2026-09-01", EndDate:"", Location:main.Location{Name:"", Latitude:0, Longitude:0}, Itinerary:[]main.ItineraryItem(nil), IsPublic:false, CreatedAt:"", UpdatedAt:"2026-08-17T02:14:31Z"}
--- FAIL: TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges (0.00s)
=== RUN   TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges
    plan_visibility_test.go:28: 只更新标题后公开状态被改成 false: main.TravelPlan{ID:"plan-public", Title:"新标题", Destination:"厦门", StartDate:"2026-09-01", EndDate:"", Location:main.Location{Name:"", Latitude:0, Longitude:0}, Itinerary:[]main.ItineraryItem(nil), IsPublic:false, CreatedAt:"", UpdatedAt:"2026-08-17T02:14:31Z"}
--- FAIL: TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges (0.00s)
=== RUN   TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges
    plan_visibility_test.go:28: 只更新标题后公开状态被改成 false: main.TravelPlan{ID:"plan-public", Title:"新标题", Destination:"厦门", StartDate:"2026-09-01", EndDate:"", Location:main.Location{Name:"", Latitude:0, Longitude:0}, Itinerary:[]main.ItineraryItem(nil), IsPublic:false, CreatedAt:"", UpdatedAt:"2026-08-17T02:14:31Z"}
--- FAIL: TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges (0.00s)
=== RUN   TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges
    plan_visibility_test.go:28: 只更新标题后公开状态被改成 false: main.TravelPlan{ID:"plan-public", Title:"新标题", Destination:"厦门", StartDate:"2026-09-01", EndDate:"", Location:main.Location{Name:"", Latitude:0, Longitude:0}, Itinerary:[]main.ItineraryItem(nil), IsPublic:false, CreatedAt:"", UpdatedAt:"2026-08-17T02:14:31Z"}
--- FAIL: TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges (0.00s)
=== RUN   TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges
    plan_visibility_test.go:28: 只更新标题后公开状态被改成 false: main.TravelPlan{ID:"plan-public", Title:"新标题", Destination:"厦门", StartDate:"2026-09-01", EndDate:"", Location:main.Location{Name:"", Latitude:0, Longitude:0}, Itinerary:[]main.ItineraryItem(nil), IsPublic:false, CreatedAt:"", UpdatedAt:"2026-08-17T02:14:31Z"}
--- FAIL: TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges (0.00s)
=== RUN   TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges
    plan_visibility_test.go:28: 只更新标题后公开状态被改成 false: main.TravelPlan{ID:"plan-public", Title:"新标题", Destination:"厦门", StartDate:"2026-09-01", EndDate:"", Location:main.Location{Name:"", Latitude:0, Longitude:0}, Itinerary:[]main.ItineraryItem(nil), IsPublic:false, CreatedAt:"", UpdatedAt:"2026-08-17T02:14:31Z"}
--- FAIL: TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges (0.00s)
=== RUN   TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges
    plan_visibility_test.go:28: 只更新标题后公开状态被改成 false: main.TravelPlan{ID:"plan-public", Title:"新标题", Destination:"厦门", StartDate:"2026-09-01", EndDate:"", Location:main.Location{Name:"", Latitude:0, Longitude:0}, Itinerary:[]main.ItineraryItem(nil), IsPublic:false, CreatedAt:"", UpdatedAt:"2026-08-17T02:14:31Z"}
--- FAIL: TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges (0.00s)
=== RUN   TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges
    plan_visibility_test.go:28: 只更新标题后公开状态被改成 false: main.TravelPlan{ID:"plan-public", Title:"新标题", Destination:"厦门", StartDate:"2026-09-01", EndDate:"", Location:main.Location{Name:"", Latitude:0, Longitude:0}, Itinerary:[]main.ItineraryItem(nil), IsPublic:false, CreatedAt:"", UpdatedAt:"2026-08-17T02:14:31Z"}
--- FAIL: TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges (0.00s)
=== RUN   TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges
    plan_visibility_test.go:28: 只更新标题后公开状态被改成 false: main.TravelPlan{ID:"plan-public", Title:"新标题", Destination:"厦门", StartDate:"2026-09-01", EndDate:"", Location:main.Location{Name:"", Latitude:0, Longitude:0}, Itinerary:[]main.ItineraryItem(nil), IsPublic:false, CreatedAt:"", UpdatedAt:"2026-08-17T02:14:31Z"}
--- FAIL: TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges (0.00s)
=== RUN   TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges
    plan_visibility_test.go:28: 只更新标题后公开状态被改成 false: main.TravelPlan{ID:"plan-public", Title:"新标题", Destination:"厦门", StartDate:"2026-09-01", EndDate:"", Location:main.Location{Name:"", Latitude:0, Longitude:0}, Itinerary:[]main.ItineraryItem(nil), IsPublic:false, CreatedAt:"", UpdatedAt:"2026-08-17T02:14:31Z"}
--- FAIL: TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges (0.00s)
=== RUN   TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges
    plan_visibility_test.go:28: 只更新标题后公开状态被改成 false: main.TravelPlan{ID:"plan-public", Title:"新标题", Destination:"厦门", StartDate:"2026-09-01", EndDate:"", Location:main.Location{Name:"", Latitude:0, Longitude:0}, Itinerary:[]main.ItineraryItem(nil), IsPublic:false, CreatedAt:"", UpdatedAt:"2026-08-17T02:14:31Z"}
--- FAIL: TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges (0.00s)
=== RUN   TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges
    plan_visibility_test.go:28: 只更新标题后公开状态被改成 false: main.TravelPlan{ID:"plan-public", Title:"新标题", Destination:"厦门", StartDate:"2026-09-01", EndDate:"", Location:main.Location{Name:"", Latitude:0, Longitude:0}, Itinerary:[]main.ItineraryItem(nil), IsPublic:false, CreatedAt:"", UpdatedAt:"2026-08-17T02:14:31Z"}
--- FAIL: TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges (0.00s)
FAIL
exit status 1
FAIL	travel-web	0.011s
```
