package main

import "fmt"

//方法接收者详解[struct相当于类的话，只定义了属性，而方法接受者就相当于java类中的成员方法，]
// 接口和java中的一样
//在 Go 语言中，方法接收者用于[定义方法与某个特定类型的关联。通过方法接收者，
//可以为类型定义行为，使得类型更具表现力和功能性。方法接收者可以是值接收者或指针接收者。]
//下面我们详细介绍这两种接收者及其区别和使用场景。
//
//1. 值接收者
//值接收者是方法接收者的一种，它使用类型的值（即一个实例的副本）来调用方法。值接收者的方法不能修改接收者的实际值，因为它操作的是接收者的副本。

//func (t Type) MethodName() {
//	// 方法体
//}

type rect2 struct {
	width, height int
}

func (r rect2) area1() int {
	return r.width * r.height
}
func test_recv1() {
	r := rect{width: 10, height: 5}
	fmt.Println("Area:", r.area()) // 输出: Area: 50
}

// 在上面的例子中，area 方法的接收者是 rect 类型的值 r。当调用 r.area() 时，r 的一个副本被传递给 area 方法。

//2. 指针接收者
//指针接收者是方法接收者的另一种，它使用类型的指针（即类型实例的地址）来调用方法。
//指针接收者的方法可以修改接收者的实际值，因为它操作的是接收者的地址。

type rect3 struct {
	width, height int
}

func (r *rect2) scale2(factor int) {
	r.width *= factor
	r.height *= factor
}

func test_recv2() {
	r := rect2{width: 10, height: 5}
	r.scale2(2)
	fmt.Println("Scaled Width:", r.width)   // 输出: Scaled Width: 20
	fmt.Println("Scaled Height:", r.height) // 输出: Scaled Height: 10
}

//在上面的例子中，scale 方法的接收者是 rect 类型的指针 r。当调用 r.scale(2) 时，
//r 的地址被传递给 scale 方法，使得方法可以修改 r 的实际值。

//3. 值接收者与指针接收者的区别
//修改接收者的值：指针接收者可以修改接收者的值，而值接收者不能。
//性能：对于大的结构体类型，使用指针接收者可以避免每次方法调用时的值拷贝，提升性能。
//一致性：如果某个方法需要修改接收者的值，那么该类型的所有方法最好都使用指针接收者，以保持一致性。
//4. 自动转换
//在 Go 语言中，当你使用一个类型的值来调用其指针接收者的方法时，Go 会自动将该值的地址传递给方法。
//反之亦然，当你使用一个类型的指针来调用其值接收者的方法时，Go 会自动解引用该指针并传递给方法。

func (r rect2) area3() int {
	return r.width * r.height
}

func (r *rect2) scale3(factor int) {
	r.width *= factor
	r.height *= factor
}

func testRec3() {
	r := rect2{width: 10, height: 5}
	fmt.Println("Area:", r.area3())        // 自动转换：(&r).area()
	r.scale3(2)                            // 自动转换：(&r).scale(2)
	fmt.Println("Scaled Area:", r.area3()) // 自动转换：(&r).area()
}

//5. 选择方法接收者的建议
//不修改接收者的值：如果方法不会修改接收者的值，通常使用值接收者。
//需要修改接收者的值：如果方法需要修改接收者的值，使用指针接收者。
//避免值拷贝：对于大的结构体类型，使用指针接收者以避免值拷贝带来的性能开销。
//总结
//方法接收者是 Go 语言中定义方法与特定类型关联的一种机制。通过值接收者和指针接收者，
//可以分别定义不修改和修改接收者值的方法。理解方法接收者的区别和使用场景，对于编写高效和一致的 Go 代码非常重要。
//通过自动转换机制，Go 提供了灵活性，使得在调用方法时不必过于担心是值还是指针接收者。
