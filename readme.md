# Golang Unit Test Tutorial
ศึกษาจาก Codebangkok -> [Reference](https://www.youtube.com/watch?v=Wd3O6GcA20w)

## คำสั่งเบื้องต้น
เข้าไปข้างใน Folder => `go test -v`<br>
ไม่ได้อยู่ใน Folder => `go test <module>/<package> (go test gotest/services -v)`<br>
รันเฉพาะเทสที่ต้องการ => `go test gotest/services -v -run=TestHello`<br>
เทสหลาย ๆ Package => `go test ./... -v`<br>
บอก test coverage => `go test gotest/services -cover`<br>
run subtest => `go test gotest/services -run="<ชื่อ function test>/<subtest>" (go test gotest/services -run="TestCheckGrade/success grade a")`<br>
run benchmark =>` go test <module>/<pkg> -bench=. -benchmem(เช็คการใช้ memory)`<br>
testify => `github.com/stretchr/testify ใช้ mock, assert`<br>
AAA => `Arrange เตรียมข้อมูล, Act ทำการเทส, Assert`<br>
httptest => `package ใช้สำหรับการทดสอบยิง request http`<br>
