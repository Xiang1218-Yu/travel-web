# 缺陷复现说明

## 缺陷现象
运营在后台把旅行目的地从厦门改成泉州后，接口返回成功但页面仍显示厦门。文件先不要改，帮我定位为什么这个字段会被悄悄忽略，并说明影响范围。这次只需要定位结论，不用修改或执行测试。

## 触发方式
# 1. 进入项目目录；预期：在旅行服务的基线环境中复现目的地更新问题。
cd /workplace/travel-web__003
# 2. 验证目的地字段更新；预期：含缺陷的基线连续 20 次显示目的地仍为厦门。
go test -v -run TestUpdatePlanChangesDestination -count=20
# 3. 检查整体回归；预期：修复版本的完整测试通过。
go test ./...

## 触发后的实际错误输出

```text
=== RUN   TestUpdatePlanChangesDestination
    plan_destination_test.go:28: destination="厦门", want 泉州
--- FAIL: TestUpdatePlanChangesDestination (0.00s)
=== RUN   TestUpdatePlanChangesDestination
    plan_destination_test.go:28: destination="厦门", want 泉州
--- FAIL: TestUpdatePlanChangesDestination (0.00s)
=== RUN   TestUpdatePlanChangesDestination
    plan_destination_test.go:28: destination="厦门", want 泉州
--- FAIL: TestUpdatePlanChangesDestination (0.00s)
=== RUN   TestUpdatePlanChangesDestination
    plan_destination_test.go:28: destination="厦门", want 泉州
--- FAIL: TestUpdatePlanChangesDestination (0.00s)
=== RUN   TestUpdatePlanChangesDestination
    plan_destination_test.go:28: destination="厦门", want 泉州
--- FAIL: TestUpdatePlanChangesDestination (0.00s)
=== RUN   TestUpdatePlanChangesDestination
    plan_destination_test.go:28: destination="厦门", want 泉州
--- FAIL: TestUpdatePlanChangesDestination (0.00s)
=== RUN   TestUpdatePlanChangesDestination
    plan_destination_test.go:28: destination="厦门", want 泉州
--- FAIL: TestUpdatePlanChangesDestination (0.00s)
=== RUN   TestUpdatePlanChangesDestination
    plan_destination_test.go:28: destination="厦门", want 泉州
--- FAIL: TestUpdatePlanChangesDestination (0.00s)
=== RUN   TestUpdatePlanChangesDestination
    plan_destination_test.go:28: destination="厦门", want 泉州
--- FAIL: TestUpdatePlanChangesDestination (0.00s)
=== RUN   TestUpdatePlanChangesDestination
    plan_destination_test.go:28: destination="厦门", want 泉州
--- FAIL: TestUpdatePlanChangesDestination (0.00s)
=== RUN   TestUpdatePlanChangesDestination
    plan_destination_test.go:28: destination="厦门", want 泉州
--- FAIL: TestUpdatePlanChangesDestination (0.00s)
=== RUN   TestUpdatePlanChangesDestination
    plan_destination_test.go:28: destination="厦门", want 泉州
--- FAIL: TestUpdatePlanChangesDestination (0.00s)
=== RUN   TestUpdatePlanChangesDestination
    plan_destination_test.go:28: destination="厦门", want 泉州
--- FAIL: TestUpdatePlanChangesDestination (0.00s)
=== RUN   TestUpdatePlanChangesDestination
    plan_destination_test.go:28: destination="厦门", want 泉州
--- FAIL: TestUpdatePlanChangesDestination (0.00s)
=== RUN   TestUpdatePlanChangesDestination
    plan_destination_test.go:28: destination="厦门", want 泉州
--- FAIL: TestUpdatePlanChangesDestination (0.00s)
=== RUN   TestUpdatePlanChangesDestination
    plan_destination_test.go:28: destination="厦门", want 泉州
--- FAIL: TestUpdatePlanChangesDestination (0.00s)
=== RUN   TestUpdatePlanChangesDestination
    plan_destination_test.go:28: destination="厦门", want 泉州
--- FAIL: TestUpdatePlanChangesDestination (0.00s)
=== RUN   TestUpdatePlanChangesDestination
    plan_destination_test.go:28: destination="厦门", want 泉州
--- FAIL: TestUpdatePlanChangesDestination (0.00s)
=== RUN   TestUpdatePlanChangesDestination
    plan_destination_test.go:28: destination="厦门", want 泉州
--- FAIL: TestUpdatePlanChangesDestination (0.00s)
=== RUN   TestUpdatePlanChangesDestination
    plan_destination_test.go:28: destination="厦门", want 泉州
--- FAIL: TestUpdatePlanChangesDestination (0.00s)
FAIL
exit status 1
FAIL	travel-web	0.851s
```
