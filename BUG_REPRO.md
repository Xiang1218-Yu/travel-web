# 缺陷复现说明

## 缺陷现象
创建游记时如果传入一个不存在的旅行计划 ID，接口仍返回 201 并写入一条无法归属的游记。请修复这个校验问题：不存在的计划必须被拒绝，已有计划下正常创建游记的行为不能受影响。请实现修复并运行相关测试。

## 触发方式
# 1. 进入项目目录；预期：在含缺陷基线中复现未知计划仍可创建游记。
cd /workplace/travel-web__004
# 2. 连续执行目标用例；预期：基线稳定失败，因为接口错误返回 201。
go test -v -run TestCreateDiaryRejectsUnknownPlan -count=20
# 3. 执行完整回归；预期：修复后没有新增失败。
go test ./...

## 触发后的实际错误输出

```text
=== RUN   TestCreateDiaryRejectsUnknownPlan
    diary_plan_test.go:19: 未知计划仍被创建，status=201 body={"id":"f2e831f10463c92310d2d4b3f9024fd4","plan_id":"missing-plan","date":"2026-09-01","title":"孤儿游记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T02:16:56Z","updated_at":"2026-08-17T02:16:56Z"}
--- FAIL: TestCreateDiaryRejectsUnknownPlan (0.00s)
=== RUN   TestCreateDiaryRejectsUnknownPlan
    diary_plan_test.go:19: 未知计划仍被创建，status=201 body={"id":"92f26f7c77375c0b7823db91bf42d667","plan_id":"missing-plan","date":"2026-09-01","title":"孤儿游记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T02:16:56Z","updated_at":"2026-08-17T02:16:56Z"}
--- FAIL: TestCreateDiaryRejectsUnknownPlan (0.00s)
=== RUN   TestCreateDiaryRejectsUnknownPlan
    diary_plan_test.go:19: 未知计划仍被创建，status=201 body={"id":"c362404babc524692fdf1a8cc326f50b","plan_id":"missing-plan","date":"2026-09-01","title":"孤儿游记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T02:16:56Z","updated_at":"2026-08-17T02:16:56Z"}
--- FAIL: TestCreateDiaryRejectsUnknownPlan (0.00s)
=== RUN   TestCreateDiaryRejectsUnknownPlan
    diary_plan_test.go:19: 未知计划仍被创建，status=201 body={"id":"77d848710e7d2de6025724a19d625449","plan_id":"missing-plan","date":"2026-09-01","title":"孤儿游记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T02:16:56Z","updated_at":"2026-08-17T02:16:56Z"}
--- FAIL: TestCreateDiaryRejectsUnknownPlan (0.00s)
=== RUN   TestCreateDiaryRejectsUnknownPlan
    diary_plan_test.go:19: 未知计划仍被创建，status=201 body={"id":"173588e526ee2acd2763ebb4497df57b","plan_id":"missing-plan","date":"2026-09-01","title":"孤儿游记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T02:16:56Z","updated_at":"2026-08-17T02:16:56Z"}
--- FAIL: TestCreateDiaryRejectsUnknownPlan (0.00s)
=== RUN   TestCreateDiaryRejectsUnknownPlan
    diary_plan_test.go:19: 未知计划仍被创建，status=201 body={"id":"af999888ccfc8a14f206e40effa0412d","plan_id":"missing-plan","date":"2026-09-01","title":"孤儿游记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T02:16:56Z","updated_at":"2026-08-17T02:16:56Z"}
--- FAIL: TestCreateDiaryRejectsUnknownPlan (0.00s)
=== RUN   TestCreateDiaryRejectsUnknownPlan
    diary_plan_test.go:19: 未知计划仍被创建，status=201 body={"id":"d03c8a3b19a25c905c68e82c281188c1","plan_id":"missing-plan","date":"2026-09-01","title":"孤儿游记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T02:16:56Z","updated_at":"2026-08-17T02:16:56Z"}
--- FAIL: TestCreateDiaryRejectsUnknownPlan (0.00s)
=== RUN   TestCreateDiaryRejectsUnknownPlan
    diary_plan_test.go:19: 未知计划仍被创建，status=201 body={"id":"a0406d0e586d6a000292e3c5454cb96c","plan_id":"missing-plan","date":"2026-09-01","title":"孤儿游记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T02:16:56Z","updated_at":"2026-08-17T02:16:56Z"}
--- FAIL: TestCreateDiaryRejectsUnknownPlan (0.00s)
=== RUN   TestCreateDiaryRejectsUnknownPlan
    diary_plan_test.go:19: 未知计划仍被创建，status=201 body={"id":"e379c08666bb13aa12ffe7399400997a","plan_id":"missing-plan","date":"2026-09-01","title":"孤儿游记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T02:16:56Z","updated_at":"2026-08-17T02:16:56Z"}
--- FAIL: TestCreateDiaryRejectsUnknownPlan (0.00s)
=== RUN   TestCreateDiaryRejectsUnknownPlan
    diary_plan_test.go:19: 未知计划仍被创建，status=201 body={"id":"9352022285e241360f9a27c0dbb42f2b","plan_id":"missing-plan","date":"2026-09-01","title":"孤儿游记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T02:16:56Z","updated_at":"2026-08-17T02:16:56Z"}
--- FAIL: TestCreateDiaryRejectsUnknownPlan (0.00s)
=== RUN   TestCreateDiaryRejectsUnknownPlan
    diary_plan_test.go:19: 未知计划仍被创建，status=201 body={"id":"1ed5469b1390e5b438945253a3b53b1f","plan_id":"missing-plan","date":"2026-09-01","title":"孤儿游记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T02:16:56Z","updated_at":"2026-08-17T02:16:56Z"}
--- FAIL: TestCreateDiaryRejectsUnknownPlan (0.00s)
=== RUN   TestCreateDiaryRejectsUnknownPlan
    diary_plan_test.go:19: 未知计划仍被创建，status=201 body={"id":"a6ff2913479051eeae3f8b78c6818934","plan_id":"missing-plan","date":"2026-09-01","title":"孤儿游记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T02:16:56Z","updated_at":"2026-08-17T02:16:56Z"}
--- FAIL: TestCreateDiaryRejectsUnknownPlan (0.00s)
=== RUN   TestCreateDiaryRejectsUnknownPlan
    diary_plan_test.go:19: 未知计划仍被创建，status=201 body={"id":"1a3482f8ebd36031fad80a7b5a7ccf07","plan_id":"missing-plan","date":"2026-09-01","title":"孤儿游记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T02:16:56Z","updated_at":"2026-08-17T02:16:56Z"}
--- FAIL: TestCreateDiaryRejectsUnknownPlan (0.00s)
=== RUN   TestCreateDiaryRejectsUnknownPlan
    diary_plan_test.go:19: 未知计划仍被创建，status=201 body={"id":"0efcf4d72cbaf4e8b925110070429e72","plan_id":"missing-plan","date":"2026-09-01","title":"孤儿游记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T02:16:56Z","updated_at":"2026-08-17T02:16:56Z"}
--- FAIL: TestCreateDiaryRejectsUnknownPlan (0.00s)
=== RUN   TestCreateDiaryRejectsUnknownPlan
    diary_plan_test.go:19: 未知计划仍被创建，status=201 body={"id":"79ce6de487e5f18e05b257c39c590c80","plan_id":"missing-plan","date":"2026-09-01","title":"孤儿游记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T02:16:56Z","updated_at":"2026-08-17T02:16:56Z"}
--- FAIL: TestCreateDiaryRejectsUnknownPlan (0.00s)
=== RUN   TestCreateDiaryRejectsUnknownPlan
    diary_plan_test.go:19: 未知计划仍被创建，status=201 body={"id":"6ad04fc75ec75bf2233e078fadbb769a","plan_id":"missing-plan","date":"2026-09-01","title":"孤儿游记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T02:16:56Z","updated_at":"2026-08-17T02:16:56Z"}
--- FAIL: TestCreateDiaryRejectsUnknownPlan (0.00s)
=== RUN   TestCreateDiaryRejectsUnknownPlan
    diary_plan_test.go:19: 未知计划仍被创建，status=201 body={"id":"88db44740340c1f85c81f8e723f213f3","plan_id":"missing-plan","date":"2026-09-01","title":"孤儿游记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T02:16:56Z","updated_at":"2026-08-17T02:16:56Z"}
--- FAIL: TestCreateDiaryRejectsUnknownPlan (0.00s)
=== RUN   TestCreateDiaryRejectsUnknownPlan
    diary_plan_test.go:19: 未知计划仍被创建，status=201 body={"id":"7ca64ed90d9314105348d2150da3bd51","plan_id":"missing-plan","date":"2026-09-01","title":"孤儿游记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T02:16:56Z","updated_at":"2026-08-17T02:16:56Z"}
--- FAIL: TestCreateDiaryRejectsUnknownPlan (0.00s)
=== RUN   TestCreateDiaryRejectsUnknownPlan
    diary_plan_test.go:19: 未知计划仍被创建，status=201 body={"id":"3a907a79fabd6dffb4837bab77a441c2","plan_id":"missing-plan","date":"2026-09-01","title":"孤儿游记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T02:16:56Z","updated_at":"2026-08-17T02:16:56Z"}
--- FAIL: TestCreateDiaryRejectsUnknownPlan (0.00s)
=== RUN   TestCreateDiaryRejectsUnknownPlan
    diary_plan_test.go:19: 未知计划仍被创建，status=201 body={"id":"d82d8da35dc4d39eab69c39289bc6bcf","plan_id":"missing-plan","date":"2026-09-01","title":"孤儿游记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T02:16:56Z","updated_at":"2026-08-17T02:16:56Z"}
--- FAIL: TestCreateDiaryRejectsUnknownPlan (0.00s)
FAIL
exit status 1
FAIL	travel-web	0.011s
```
