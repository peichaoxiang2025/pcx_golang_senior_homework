package main

import (
	"fmt"
	"math"
)

// Shape 定义几何形状的接口，包含面积和周长计算方法
type Shape interface {
	Area() float64      // 计算面积
	Perimeter() float64 // 计算周长
}

// Rectangle 矩形结构体，包含宽和高两个属性
type Rectangle struct {
	Width  float64 // 宽度
	Height float64 // 高度
}

// Area 计算矩形面积（宽 × 高）
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter 计算矩形周长（2 × (宽 + 高)）
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle 圆形结构体，包含半径属性
type Circle struct {
	Radius float64 // 半径
}

// Area 计算圆形面积（π × 半径²）
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter 计算圆形周长（2 × π × 半径）
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	// 创建矩形实例（宽3，高4）
	rect := Rectangle{
		Width:  3.0,
		Height: 4.0,
	}

	// 创建圆形实例（半径5）
	circle := Circle{
		Radius: 5.0,
	}

	// 调用接口方法计算面积和周长
	fmt.Printf("矩形信息：\n")
	fmt.Printf("  面积: %.2f\n", rect.Area())      // 输出：12.00
	fmt.Printf("  周长: %.2f\n", rect.Perimeter()) // 输出：14.00

	fmt.Printf("\n圆形信息：\n")
	fmt.Printf("  面积: %.2f\n", circle.Area())      // 输出：78.54
	fmt.Printf("  周长: %.2f\n", circle.Perimeter()) // 输出：31.42
}
