# 缺陷复现说明

## 缺陷现象
旅行计划已经在 9 月 12 日结束，但用户还能新增 9 月 13 日的游记，接口返回 201，导致行程外的内容混进计划。请修复这个问题，保留同一计划内正常新增游记的行为。

## 触发方式
# 1. 进入项目目录；预期：在旅行服务的测试环境中执行后续命令。
cd /workplace/travel-web__001
# 2. 验证行程结束后的游记会被拒绝；预期：修复前连续 20 次失败，修复后连续 20 次通过。
go test -v -run TestCreateDiaryRejectsDateAfterPlanEnd -count=20
# 3. 回归全部行为；预期：修复后没有新增失败。
go test ./...

## 触发后的实际错误输出

```text
=== RUN   TestCreateDiaryRejectsDateAfterPlanEnd
    diary_date_test.go:31: status = 201, want 400; body={"id":"13d8fd3d01e620774e54fb88d6512339","plan_id":"plan-1","date":"2026-09-13","title":"返程后补记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T01:30:39Z","updated_at":"2026-08-17T01:30:39Z"}
--- FAIL: TestCreateDiaryRejectsDateAfterPlanEnd (0.00s)
=== RUN   TestCreateDiaryRejectsDateAfterPlanEnd
    diary_date_test.go:31: status = 201, want 400; body={"id":"4dc49596ae086e3d72043fce1e904b5d","plan_id":"plan-1","date":"2026-09-13","title":"返程后补记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T01:30:39Z","updated_at":"2026-08-17T01:30:39Z"}
--- FAIL: TestCreateDiaryRejectsDateAfterPlanEnd (0.00s)
=== RUN   TestCreateDiaryRejectsDateAfterPlanEnd
    diary_date_test.go:31: status = 201, want 400; body={"id":"082b88cf30a20deba5b561ebc1404776","plan_id":"plan-1","date":"2026-09-13","title":"返程后补记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T01:30:39Z","updated_at":"2026-08-17T01:30:39Z"}
--- FAIL: TestCreateDiaryRejectsDateAfterPlanEnd (0.00s)
=== RUN   TestCreateDiaryRejectsDateAfterPlanEnd
    diary_date_test.go:31: status = 201, want 400; body={"id":"b54aad91dd695e5f5493f535c7a43ee0","plan_id":"plan-1","date":"2026-09-13","title":"返程后补记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T01:30:39Z","updated_at":"2026-08-17T01:30:39Z"}
--- FAIL: TestCreateDiaryRejectsDateAfterPlanEnd (0.00s)
=== RUN   TestCreateDiaryRejectsDateAfterPlanEnd
    diary_date_test.go:31: status = 201, want 400; body={"id":"aa1eb2865df78f72dc68488f8c465a66","plan_id":"plan-1","date":"2026-09-13","title":"返程后补记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T01:30:39Z","updated_at":"2026-08-17T01:30:39Z"}
--- FAIL: TestCreateDiaryRejectsDateAfterPlanEnd (0.00s)
=== RUN   TestCreateDiaryRejectsDateAfterPlanEnd
    diary_date_test.go:31: status = 201, want 400; body={"id":"b68820e7f3d7d5c553da9ae7568bf439","plan_id":"plan-1","date":"2026-09-13","title":"返程后补记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T01:30:39Z","updated_at":"2026-08-17T01:30:39Z"}
--- FAIL: TestCreateDiaryRejectsDateAfterPlanEnd (0.00s)
=== RUN   TestCreateDiaryRejectsDateAfterPlanEnd
    diary_date_test.go:31: status = 201, want 400; body={"id":"02e923fd40da920938058e0a3a7ebc3d","plan_id":"plan-1","date":"2026-09-13","title":"返程后补记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T01:30:39Z","updated_at":"2026-08-17T01:30:39Z"}
--- FAIL: TestCreateDiaryRejectsDateAfterPlanEnd (0.00s)
=== RUN   TestCreateDiaryRejectsDateAfterPlanEnd
    diary_date_test.go:31: status = 201, want 400; body={"id":"c5b3266716e8ab6c8d5c2ffc8b0dc765","plan_id":"plan-1","date":"2026-09-13","title":"返程后补记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T01:30:39Z","updated_at":"2026-08-17T01:30:39Z"}
--- FAIL: TestCreateDiaryRejectsDateAfterPlanEnd (0.00s)
=== RUN   TestCreateDiaryRejectsDateAfterPlanEnd
    diary_date_test.go:31: status = 201, want 400; body={"id":"159eaeac4adc90f173985192a3b7e5a7","plan_id":"plan-1","date":"2026-09-13","title":"返程后补记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T01:30:39Z","updated_at":"2026-08-17T01:30:39Z"}
--- FAIL: TestCreateDiaryRejectsDateAfterPlanEnd (0.00s)
=== RUN   TestCreateDiaryRejectsDateAfterPlanEnd
    diary_date_test.go:31: status = 201, want 400; body={"id":"ee064813968eb78ea699c2fc248f0d36","plan_id":"plan-1","date":"2026-09-13","title":"返程后补记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T01:30:39Z","updated_at":"2026-08-17T01:30:39Z"}
--- FAIL: TestCreateDiaryRejectsDateAfterPlanEnd (0.00s)
=== RUN   TestCreateDiaryRejectsDateAfterPlanEnd
    diary_date_test.go:31: status = 201, want 400; body={"id":"ecf13e36b4d752bd426f65b9c734ff15","plan_id":"plan-1","date":"2026-09-13","title":"返程后补记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T01:30:39Z","updated_at":"2026-08-17T01:30:39Z"}
--- FAIL: TestCreateDiaryRejectsDateAfterPlanEnd (0.00s)
=== RUN   TestCreateDiaryRejectsDateAfterPlanEnd
    diary_date_test.go:31: status = 201, want 400; body={"id":"52fbb8767a841e72b8df4cb7fc25daa5","plan_id":"plan-1","date":"2026-09-13","title":"返程后补记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T01:30:39Z","updated_at":"2026-08-17T01:30:39Z"}
--- FAIL: TestCreateDiaryRejectsDateAfterPlanEnd (0.00s)
=== RUN   TestCreateDiaryRejectsDateAfterPlanEnd
    diary_date_test.go:31: status = 201, want 400; body={"id":"bc6b1277a2a7e427fb2bf369ec1a468a","plan_id":"plan-1","date":"2026-09-13","title":"返程后补记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T01:30:39Z","updated_at":"2026-08-17T01:30:39Z"}
--- FAIL: TestCreateDiaryRejectsDateAfterPlanEnd (0.00s)
=== RUN   TestCreateDiaryRejectsDateAfterPlanEnd
    diary_date_test.go:31: status = 201, want 400; body={"id":"f09c96f5f3eb201bbbfe28fa39d2f3be","plan_id":"plan-1","date":"2026-09-13","title":"返程后补记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T01:30:39Z","updated_at":"2026-08-17T01:30:39Z"}
--- FAIL: TestCreateDiaryRejectsDateAfterPlanEnd (0.00s)
=== RUN   TestCreateDiaryRejectsDateAfterPlanEnd
    diary_date_test.go:31: status = 201, want 400; body={"id":"8fd19905369a697ab6adc809c8c3f7f8","plan_id":"plan-1","date":"2026-09-13","title":"返程后补记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T01:30:39Z","updated_at":"2026-08-17T01:30:39Z"}
--- FAIL: TestCreateDiaryRejectsDateAfterPlanEnd (0.00s)
=== RUN   TestCreateDiaryRejectsDateAfterPlanEnd
    diary_date_test.go:31: status = 201, want 400; body={"id":"8511547708c3f6e0d5abb0662371f711","plan_id":"plan-1","date":"2026-09-13","title":"返程后补记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T01:30:39Z","updated_at":"2026-08-17T01:30:39Z"}
--- FAIL: TestCreateDiaryRejectsDateAfterPlanEnd (0.00s)
=== RUN   TestCreateDiaryRejectsDateAfterPlanEnd
    diary_date_test.go:31: status = 201, want 400; body={"id":"77048be5b7eeb45318a168727caf2946","plan_id":"plan-1","date":"2026-09-13","title":"返程后补记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T01:30:39Z","updated_at":"2026-08-17T01:30:39Z"}
--- FAIL: TestCreateDiaryRejectsDateAfterPlanEnd (0.00s)
=== RUN   TestCreateDiaryRejectsDateAfterPlanEnd
    diary_date_test.go:31: status = 201, want 400; body={"id":"7ef725a98ea160dd2d1555551399a1af","plan_id":"plan-1","date":"2026-09-13","title":"返程后补记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T01:30:39Z","updated_at":"2026-08-17T01:30:39Z"}
--- FAIL: TestCreateDiaryRejectsDateAfterPlanEnd (0.00s)
=== RUN   TestCreateDiaryRejectsDateAfterPlanEnd
    diary_date_test.go:31: status = 201, want 400; body={"id":"ef40e6944bc14fa81cd342034ee6e056","plan_id":"plan-1","date":"2026-09-13","title":"返程后补记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T01:30:39Z","updated_at":"2026-08-17T01:30:39Z"}
--- FAIL: TestCreateDiaryRejectsDateAfterPlanEnd (0.00s)
=== RUN   TestCreateDiaryRejectsDateAfterPlanEnd
    diary_date_test.go:31: status = 201, want 400; body={"id":"d2cd854f20fd19a47fc3fd56fc99c9c6","plan_id":"plan-1","date":"2026-09-13","title":"返程后补记","content":"","images":[],"location":{"name":"","latitude":0,"longitude":0},"created_at":"2026-08-17T01:30:39Z","updated_at":"2026-08-17T01:30:39Z"}
--- FAIL: TestCreateDiaryRejectsDateAfterPlanEnd (0.00s)
FAIL
exit status 1
FAIL	travel-web	0.010s
```
