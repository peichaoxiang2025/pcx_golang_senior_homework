package main

import "fmt"

// ------------------------------
// 1. 定义基础结构体 Person
// ------------------------------
// Person 表示一个人的基础信息（姓名、年龄）
type Person struct {
	Name string // 姓名（导出字段，首字母大写）
	Age  int    // 年龄（导出字段，首字母大写）
}

// ------------------------------
// 2. 定义组合结构体 Employee
// ------------------------------
// Employee 表示员工信息，组合了 Person 结构体，并添加员工 ID 字段
type Employee struct {
	Person            // 嵌入 Person 结构体（组合的核心）
	EmployeeID string // 员工 ID（新增字段）
}

// ------------------------------
// 3. 为 Employee 实现 PrintInfo 方法
// ------------------------------
// PrintInfo 输出员工的完整信息（姓名、年龄、员工 ID）
func (e Employee) PrintInfo() {
	// 直接访问嵌入的 Person 结构体的字段（因为字段是导出的）
	fmt.Printf("员工信息:\n")
	fmt.Printf("  姓名: %s\n", e.Name) // 等价于 e.Person.Name
	fmt.Printf("  年龄: %d\n", e.Age)  // 等价于 e.Person.Age
	fmt.Printf("  员工ID: %s\n", e.EmployeeID)
}

func main() {
	// ------------------------------
	// 4. 创建 Employee 实例并调用方法
	// ------------------------------
	// 初始化 Employee 实例（嵌套初始化 Person）
	emp := Employee{
		Person: Person{
			Name: "李华", // 设置 Person 的 Name 字段
			Age:  28,   // 设置 Person 的 Age 字段
		},
		EmployeeID: "DEV007", // 设置 Employee 的 EmployeeID 字段
	}

	// 调用 Employee 的 PrintInfo 方法
	emp.PrintInfo()
}
