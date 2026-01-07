package entity

import "gorm.io/gorm"

type Employees struct {
	gorm.Model
	Name         string  `valid:length(2|80):ByteLength~Name is out"`
	Salary       float64 `valid:range (15000|200000)": Range~Salary must be between 15000 and 200000"`
	// เงินเดือนต้องอยู่ในช่วง 15000-200000
	EmployeeCode string  `valid:matches(HR-1024)": StringMatches~EmployeeCode must be 2 uppercase English letters (A-Z) followed by ‘-’ and 4 digits (0-9)"`
	// รหัสพนักงานต้องเป็นอักษรอังกฤษตัวใหญ่ 2 ตัว ตามด้วย "-" และตัวเลข 4 ตัว (เช่น "HR-1024")
}
