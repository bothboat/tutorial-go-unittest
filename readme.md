# Golang Unit Test Tutorial
ศึกษาจาก Codebangkok -> ##[Reference](https://www.youtube.com/watch?v=Wd3O6GcA20w)

## คำสั่งเบื้องต้น
เข้าไปข้างใน Folder => `go test -v`
ไม่ได้อยู่ใน Folder => `go test <module>/<package> (go test gotest/services -v)`
รันเฉพาะเทสที่ต้องการ => `go test gotest/services -v -run=TestHello`
เทสหลาย ๆ Package => `go test ./... -v`
บอก test coverage => `go test gotest/services -cover`
run subtest => `go test gotest/services -run="<ชื่อ function test>/<subtest>" (go test gotest/services -run="TestCheckGrade/success grade a")`
run benchmark =>` go test <module>/<pkg> -bench=. -benchmem(เช็คการใช้ memory)`
testify => `github.com/stretchr/testify ใช้ mock, assert`
AAA => `Arrange เตรียมข้อมูล, Act ทำการเทส, Assert`
httptest => `package ใช้สำหรับการทดสอบยิง request http`