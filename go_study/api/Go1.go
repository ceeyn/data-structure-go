package main

// 在 Go 语言中，package main 表示当前的代码文件所属的包是 main 包。
//
// main 包比较特殊，当一个 Go 程序的入口函数（func main()）所在的代码文件属于 main 包时，
// 这个程序可以被编译为可执行文件。如果包名不是 main，则通常是一个库包，用于被其他代码引用和使用。
// 在 Go 语言中创建自己的包，您可以按照以下步骤进行操作：
//创建一个目录来存放您的包代码。例如，如果您要创建一个名为 mypackage 的包，可以创建一个名为 mypackage 的文件夹。
//在该文件夹内创建 Go 源文件（例如 mypackage.go ）。
import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unicode/utf8"
)

func main() {
	//testInt()
	//testFor()
	//fmt.Println(testFunc(0, 3))
	// testArray()
	// exampleSlice()
	//testTimer()
	// testLimitSpeed()
	//testChannelRoutineSync()
}

func testInt() {
	// 声明一个整数变量
	var num int
	num = 10
	fmt.Println(num)

	// 声明并初始化一个字符串变量
	var str = "Hello"
	fmt.Println(str)

	// 整数类型
	var intVar int = 10
	fmt.Println(intVar)

	// 浮点数类型
	var floatVar float64 = 3.14
	fmt.Println(floatVar)

	// 字符串类型
	var strVar string = "Hello"
	fmt.Println(strVar)

	// 布尔类型
	var boolVar bool = true
	fmt.Println(boolVar)

	// 短变量声明
	name := "Alice"
	fmt.Println(name)

	// 声明一个常量
	const pi = 3.14
	fmt.Println(pi)

	// 常量可以在声明时不指定类型，由编译器推断
	const num1 = 5
	fmt.Println(num1)
}

func testFor() {
	// 基本的 for 循环
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// 类似 while 的 for 循环
	j := 0
	for j < 3 {
		j++
	}
}

func testFunc(a int, b int) int {
	return a + b
}

func testArray() {
	/**

	数组：
	数组的长度是固定的，在声明时就需要指定。
	声明方式：var arrayName [length]type ，例如：var arr [5]int 。
	可以通过索引访问数组元素，索引从 0 开始。

	切片：

	切片是对数组的一种抽象，可以动态地增加或减少元素的数量。
	声明方式：var sliceName []type ，例如：var s []int 。
	可以使用 make 函数创建切片，指定长度和容量，例如：s := make([]int, 5, 10) ，
	表示创建一个初始长度为 5，容量为 10 的整数切片。
	通过 append 函数向切片添加元素。
	切片可以基于数组创建，例如：arr := [5]int{1, 2, 3, 4, 5} s := arr[1:3] ，
	创建了一个包含 arr 中索引 1 到 2 元素的切片。
	*/
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	fmt.Println(arr)
	s := []int{1, 2, 3, 4}
	fmt.Println("切片:", s)
	s = append(s, 10)
	fmt.Println("添加元素后的切片:", s)
	arr2 := [5]int{10, 11, 12, 13, 14}
	s2 := arr2[2:4]
	fmt.Println("基于数组创建的切片:", s2)
}

func exampleVar() {
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	var a = "initial"
	fmt.Println(a)
	// 你可以一次性声明多个变量。
	var b, c int = 1, 2
	fmt.Println(b, c)
	// Go 会自动推断已经有初始值的变量的类型。
	var d = true
	fmt.Println(d)
	// 声明后却没有给出对应的初始值时，变量将会初始化为 零值 。 例如，int 的零值是 0。
	var e int
	fmt.Println(e)
	// := 语法是声明并初始化变量的简写， 例如 var f string = "short" 可以简写为右边这样。
	f := "short"
	fmt.Println(f)
}

// const 语句可以出现在任何 var 语句可以出现的地方
const s string = "constant"

func exampleConst() {
	fmt.Println(s)
	const n = 500000000
	//常数表达式可以执行任意精度的运算
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	//一个数字可以根据上下文的需要（比如变量赋值、函数调用）自动确定类型。
	//举个例子，这里的 math.Sin 函数需要一个 float64 的参数，n 会自动确定类型
	fmt.Println(math.Sin(n))
	// 在条件语句之前可以有一个声明语句；在这里声明的变量可以在这个语句所有的条件分支中使用。
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
func exampleSwitch() {
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
func exampleArray() {
	var b [5]int = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < 5; i++ {
		fmt.Println(b[i])
	}

	var a [2][1]int = [2][1]int{{1}, {2}}
	for i := 0; i < 2; i++ {
		for j := 0; j < 1; j++ {
			fmt.Println(a[i][j])
		}
	}
}
func exampleSlice() {
	//除了基本操作外，slice 支持比数组更丰富的操作。比如 slice 支持内建函数 append，
	//该函数会返回一个包含了一个或者多个新值的 slice。 注意由于 append 可能返回一个新的 slice，我们需要接收其返回值。
	s := make([]string, 3)
	var s1 []int
	fmt.Println(s1)
	var s2 []int = []int{1, 2, 3, 4, 5}
	fmt.Println(s2)
	s = append(s, "a")
	s = append(s, "b")
	s = append(s, "c")
	fmt.Println(s)
	c := make([]string, len(s))
	//slice 还可以 copy。这里我们创建一个空的和 s 有相同长度的 slice——c， 然后将 s 复制给 c
	copy(c, s)
	// slice 支持通过 slice[low:high] 语法进行“切片”操作。
	//例如，右边的操作可以得到一个包含元素 s[2]、s[3] 和 s[4] 的 slice。
	l := c[2:3]
	fmt.Println("sl:", l)
	twoD := make([][]int, 3)
	// Slice 可以组成多维数据结构。内部的 slice 长度可以不一致，这一点和多维数组不同。
	for j := 0; j < 3; j++ {
		innerLen := j + 1
		twoD[j] = make([]int, innerLen)
		for k := 0; k < innerLen; k++ {
			twoD[j][k] = k + j
		}
	}
	fmt.Println("2d: ", twoD)
}
func exampleMap() {
	// 要创建一个空 map，需要使用内建函数 make：make(map[key-type]val-type)。
	m := make(map[string]int)
	var m1 map[string]int
	var m2 map[string]int = make(map[string]int)
	fmt.Println(m2)
	// 使用典型的 name[key] = val 语法来设置键值对
	// 当从一个 map 中取值时，还有可以选择是否接收的第二个返回值，该值表明了 map 中是否存在这个键。 这可以用来消除 键不存在 和 键的值为零值
	//产生的歧义， 例如 0 和 ""。这里我们不需要值，所以用 空白标识符(blank identifier) _ 将其忽略
	_, prs := m["k2"]
	fmt.Println(prs)
	var m3 map[string]int = map[string]int{"f": 1, "f2": 2}
	for k, v := range m3 {
		fmt.Println(k, v)
	}
	// 内建函数 delete 可以从一个 map 中移除键值对。
	delete(m1, "f")
	fmt.Println(m1)
}
func exampleRange() {
	// range 在数组和 slice 中提供对每项的索引和值的访问。
	// 上面我们不需要索引，所以我们使用 空白标识符 _ 将其忽略。 实际上，我们有时候是需要这个索引的。
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// range 也可以只遍历 map 的键
	m := make(map[string]int)
	for k := range m {
		fmt.Println(k)
	}
	// range 在字符串中迭代 unicode 码点(code point)。
	// 第一个返回值是字符的起始字节位置，然后第二个是字符本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
func vals() (int, int) {
	return 3, 7
}

// Go 原生支持 _多返回值_。 这个特性在 Go 语言中经常用到，例如用来同时返回一个函数的结果和错误信息
func exampleMultiReturn() {
	// 仅仅需要返回值的一部分的话，你可以使用空白标识符 _。
	_, c := vals()
	fmt.Println(c)
}
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
func exampleSum() {
	// 如果你有一个含有多个值的 slice，想把它们作为参数使用， 你需要这样调用 func(slice...)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

// intSeq 函数返回一个在其函数体内定义的匿名函数。
// 返回的函数使用闭包的方式 隐藏 变量 i。 返回的函数 隐藏 变量 i 以形成闭包
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// 我们调用 intSeq 函数，将返回值（一个函数）赋给 nextInt。
// 这个函数的值包含了自己的值 i，这样在每次调用 nextInt 时，都会更新 i 的值
func exampleBi() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	newInts := intSeq()
	fmt.Println(newInts())
}
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func exampleDi() {
	fmt.Println(fact(7))
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
}

/*
*
我们将通过两个函数：zeroval 和 zeroptr 来比较 指针 和 值。
zeroval 有一个 int 型参数，所以使用值传递。 zeroval 将从调用它的那个函数中得到一个实参的拷贝：ival。
*/
func zeroval(ival int) {
	ival = 0
}

/*
*
zeroptr 有一个和上面不同的参数：*int，这意味着它使用了 int 指针。 紧接着，函数体内的 *iptr
会 解引用 这个指针，从它的内存地址得到这个地址当前对应的值。 对解引用的指针赋值，会改变这个指针引用的真实地址的值。
*/
func zeroptr(iptr *int) {
	*iptr = 0
}
func examplePtr() {
	// zeroval 在 main 函数中不能改变 i 的值， 但是 zeroptr 可以，因为它有这个变量的内存地址的引用。
	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)
	// 通过 &i 语法来取得 i 的内存地址，即指向 i 的指针。
	zeroptr(&i)
	fmt.Println("zeroptr:", i)
	// 指针也是可以被打印的。
	fmt.Println("pointer:", &i)
}

/*
*

Go语言中的字符串是一个只读的byte类型的切片。 Go语言和标准库特别对待字符串 - 作为以 UTF-8 为编码的文本容器。 在其他语言当中， 字符串由”字符”组成。
在Go语言当中，字符的概念被称为 rune - 它是一个表示 Unicode 编码的整数。 这个Go博客 很好的介绍了这个主题。
*/
func exampleRune() {
	// s 是一个 string 分配了一个 literal value 表示泰语中的单词 “hello” 。
	// Go 字符串是 UTF-8 编码的文本。
	const s = "สวัสดี"
	// 因为字符串等价于 []byte， 这会产生存储在其中的原始字节的长度。
	fmt.Println("Len:", len(s))
	// 对字符串进行索引会在每个索引处生成原始字节值。 这个循环生成构成s中 Unicode 的所有字节的十六进制值。
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
	// 要计算字符串中有多少rune，我们可以使用utf8包。 注意RuneCountInString的运行时取决于字符串的大小。 因为它必须按顺序解码每个 UTF-8 rune。 一些泰语字符由多个 UTF-8 code point 表示， 所以这个计数的结果可能会令人惊讶。
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
	// 我们可以通过显式使用 utf8.DecodeRuneInString 函数来实现相同的迭代。
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		//这演示了将 rune value 传递给函数。
		examineRune(runeValue)
	}

}
func examineRune(r rune) {
	//用单引号括起来的值是 rune literals. 我们可以直接将 rune value 与 rune literal 进行比较。
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}

type person struct {
	name string
	age  int
}

// newPerson 使用给定的名字构造一个新的 person 结构体.
func newPerson(name string) *person {
	//您可以安全地返回指向局部变量的指针， 因为局部变量将在函数的作用域中继续存在。
	p := person{name: name}
	p.age = 42
	return &p
}
func examplePerson() {
	// 1.结构体初始化方式
	p1 := &person{}
	fmt.Println(p1)

	// 2.new方法
	p2 := new(person)
	fmt.Println(p2)

	// 使用这个语法创建新的结构体元素。
	fmt.Println(person{"Bob", 20})
	// 你可以在初始化一个结构体元素时指定字段名字。
	fmt.Println(person{name: "Alice", age: 30})
	// 省略的字段将被初始化为零值。
	fmt.Println(person{name: "Fred"})
	// & 前缀生成一个结构体指针。
	fmt.Println(&person{name: "Ann", age: 40})
	// 在构造函数中封装创建新的结构实例是一种习惯用法
	fmt.Println(newPerson("Jon"))
	// 使用.来访问结构体字段。
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	// 也可以对结构体指针使用. - 指针会被自动解引用。
	sp := &s
	fmt.Println(sp.age)
	// 结构体是可变(mutable)的。
	sp.age = 51
	fmt.Println(sp.age)
}

type rect1 struct {
	width, height int
}

func (r *rect1) area1() int {
	return r.width * r.height
}

// 在 func (c circle) perim() float64 中，(c circle) 表示这是一个为 circle 类型定义的方法。
// 其中 c 是接收者，circle 是类型。通过这种方式，可以让 circle 类型拥有 perim 这个方法来执行特定的操作。
func (r rect1) perim1() int {
	return 2*r.width + 2*r.height
}
func exampleMethods() {
	r := rect{width: 10, height: 5}
	// 这里我们调用上面为结构体定义的两个方法。
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())
	// 调用方法时，Go 会自动处理值和指针之间的转换。 想要避免在调用方法时产生一个拷贝，
	// 或者想让方法可以修改接受结构体的值， 你都可以使用指针来调用方法。
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}

// 接口 geometry：定义了一个名为 geometry 的接口，
// 包含两个方法：area 和 perim。任何实现了这两个方法的类型，都可以被视为实现了 geometry 接口
// 通过接口实现多态性，可以将不同类型的实例传递给同一个函数 measure，并在函数内部调用相同的方法。
// 这种方式使得代码更加灵活和通用，易于扩展。
type geometry interface {
	area() float64
	perim() float64
}

// 在这个例子中，我们将为 rect 和 circle 实现该接口。
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// 要在 Go 中实现一个接口，我们只需要实现接口中的所有方法。 这里我们为 rect 实现了 geometry 接口。
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// circle 的实现。
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 如果一个变量实现了某个接口，我们就可以调用指定接口中的方法。 这儿有一个通用的 measure 函数，我们可以通过它来使用所有的 geometry。
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
func exampleInterface() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	// 结构体类型 circle 和 rect 都实现了 geometry 接口， 所以我们可以将其实例作为 measure 的参数。
	measure(r)
	measure(c)
}

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func exampleEmb() {
	// 当创建含有嵌入的结构体，必须对嵌入进行显式的初始化； 在这里使用嵌入的类型当作字段的名字
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}
	//我们可以直接在 co 上访问 base 中定义的字段, 例如： co.num.
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	// 或者，我们可以使用嵌入的类型名拼写出完整的路径。
	fmt.Println("also num:", co.base.num)
	// 由于 container 嵌入了 base，因此base的方法 也成为了 container 的方法。
	// 在这里我们直接在 co 上 调用了一个从 base 嵌入的方法。
	fmt.Println("describe:", co.describe())
	type describer interface {
		describe() string
	}
	// 可以使用带有方法的嵌入结构来赋予接口实现到其他结构上。 因为嵌入了 base ，
	// 在这里我们看到 container 也实现了 describer 接口。
	var d describer = co
	fmt.Println("describer:", d.describe())
}

/*
*
默认发送和接收操作是阻塞的，直到发送方和接收方都就绪。
这个特性允许我们，不使用任何其它的同步操作， 就可以在程序结尾处等待消息 "ping"。
*/
func testChannel() {
	messages := make(chan string)
	//使用 channel <- 语法 发送 一个新的值到通道中。 这里我们在一个新的协程中发送 "ping" 到上面创建的 messages 通道中。
	go func() { messages <- "ping" }()
	//使用 <-channel 语法从通道中 接收 一个值。 这里我们会收到在上面发送的 "ping" 消息并将其打印出来。
	msg := <-messages
	fmt.Println(msg)
}

// 默认情况下，通道是 无缓冲 的，这意味着只有对应的接收（<- chan）
// 通道准备好接收时，才允许进行发送（chan <-）。
// 有缓冲通道 允许在没有对应接收者的情况下，缓存一定数量的值。
func testCache() {
	// 这里我们 make 了一个字符串通道，最多允许缓存 2 个值。
	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"
	// 然后我们可以正常接收这两个值。
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

func testChannelFor() {
	//for t := range time.Tick(200 * time.Millisecond)
	//for ... range 循环：用于从通道中接收值。
	//t：在每次循环中，t 是从 time.Tick(200 * time.Millisecond) 通道接收到的时间值。
	//t 的类型是 time.Time，表示当前时间。
	burstyLimiter := make(chan time.Time, 3)
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()
}
func ChildWait(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

/*
*
可以使用通道来同步协程之间的执行状态。
这儿有一个例子，使用阻塞接收的方式，实现了等待另一个协程完成。 如果需要等待多个协程，WaitGroup 是一个更好的选择。
*/
func testChannelWait() {
	// 我们将要在协程中运行这个函数。
	// done 通道将被用于通知其他协程这个函数已经完成工作
	// 运行一个 worker 协程，并给予用于通知的通道
	done := make(chan bool, 1)
	go ChildWait(done)
	// 程序将一直阻塞，直至收到 worker 使用通道发送的通知。
	is_done := <-done
	fmt.Println(is_done)
	// 如果你把 <- done 这行代码从程序中移除，
	// 程序甚至可能在 worker 开始运行前就结束了

}

/**
当使用通道作为函数的参数时，你可以指定这个通道是否为只读或只写。 该特性可以提升程序的类型安全。
*/
// ping 函数定义了一个只能发送数据的（只写）通道。 尝试从这个通道接收数据会是一个编译时错误。
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pong 函数接收两个通道，pings 仅用于接收数据（只读），pongs 仅用于发送数据（只写）。
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
func testChannelD() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

/*
*
Go 的 选择器（select） 让你可以同时等待多个通道操作。 将协程、通道和选择器结合，是 Go 的一个强大特性。
*/
func testSelector() {
	c1 := make(chan string)
	c2 := make(chan string)
	// 各个通道将在一定时间后接收一个值， 通过这种方式来模拟并行的协程执行（例如，RPC 操作）时造成的阻塞（耗时）。
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
	// 我们使用 select 关键字来同时等待这两个值， 并打印各自接收到的值。
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

/*
*
超时 对于一个需要连接外部资源， 或者有耗时较长的操作的程序而言是很重要的。
得益于通道和 select，在 Go 中实现超时操作是简洁而优雅的。
*/
func testTimeOut() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()
	// 这里是使用 select 实现一个超时操作。 res := <- c1 等待结果，<-time.After 等待超时（1秒钟）以后发送的值。 由于 select 默认处理第一个已准备好的接收操作， 因此如果操作耗时超过了允许的 1 秒的话，将会执行超时 case。
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}
	// 如果我们允许一个长一点的超时时间：3 秒， 就可以成功的从 c2 接收到值，并且打印出结果。
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

/*
*
常规的通过通道发送和接收数据是阻塞的。
然而，我们可以使用带一个 default 子句的 select 来实现 非阻塞 的发送、接收，甚至是非阻塞的多路 select
*/
func testNoZuse() {
	messages := make(chan string)
	signals := make(chan bool)
	//这是一个非阻塞接收的例子。 如果在 messages 中存在，然后 select 将这个值带入 <-messages case 中。 否则，就直接到 default 分支中。
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}
	//一个非阻塞发送的例子，代码结构和上面接收的类似。 msg 不能被发送到 message 通道，因为这是 个无缓冲区通道，并且也没有接收者，因此， default 会执行。
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}
	//我们可以在 default 前使用多个 case 子句来实现一个多路的非阻塞的选择器。 这里我们试图在 messages 和 signals 上同时使用非阻塞的接收操作。
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

/*
*
关闭 一个通道意味着不能再向这个通道发送值了。 该特性可以向通道的接收方传达工作已经完成的信息。
*/
func testChannelClos() {
	jobs := make(chan int, 5)
	done := make(chan bool)
	// 在这个例子中，我们将使用一个 jobs 通道，将工作内容， 从 main() 协程传递到一个工作协程中。
	// 当我们没有更多的任务传递给工作协程时，我们将 close 这个 jobs 通道。
	// 这是工作协程。使用 j, more := <- jobs 循环的从 jobs 接收数据。
	// 根据接收的第二个值，如果 jobs 已经关闭了， 并且通道中所有的值都已经接收完毕，那么 more 的值将是 false。
	// 当我们完成所有的任务时，会使用这个特性通过 done 通道通知 main 协程。
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()
	// 使用 jobs 发送 3 个任务到工作协程中，然后关闭 jobs
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	//使用前面学到的通道同步方法等待任务结束。
	<-done
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}
func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

//关闭文件时，进行错误检查是非常重要的， 即使在 defer 函数中也是如此。

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
func testDefer() {
	// defer 关键字用于注册一个延迟执行的函数，该函数会在包含 defer 语句的函数返回之前执行,
	//Defer 用于确保程序在执行完成后，会调用某个函数，一般是执行清理工作。
	//Defer 的用途跟其他语言的 ensure 或 finally 类似。
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
	//假设我们想要创建一个文件、然后写入数据、最后在程序结束时关闭该文件。 这里展示了如何通过 defer 来做到这一切.
	//在 createFile 后立即得到一个文件对象， 我们使用 defer 通过 closeFile 来关闭这个文件。
	//这会在封闭函数（main）结束时执行，即 writeFile 完成以后。
}

/*
*
for 和 range 为基本的数据结构提供了迭代的功能。
我们也可以使用这个语法来遍历的从通道中取值。
*/
func testForChannel() {
	// 我们将遍历在 queue 通道中的两个值。
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	// range 迭代从 queue 中得到每个值。 因为我们在前面 close 了这个通道，
	// 所以，这个迭代会在接收完 2 个值之后结束。
	for elem := range queue {
		fmt.Println(elem)
	}
	// 这个例子也让我们看到，一个非空的通道也是可以关闭的， 并且，通道中剩下的值仍然可以被接收到。
}

/*
*
我们经常需要在未来的某个时间点运行 Go 代码，或者每隔一定时间重复运行代码。
Go 内置的 定时器 和 打点器 特性让这些变得很简单。 我们会先学习定时器，然后再学习打点器。
*/
func testTimer() {
	//定时器表示在未来某一时刻的独立事件。 你告诉定时器需要等待的时间，
	//然后它将提供一个用于通知的通道。 这里的定时器将等待 2 秒。
	timer1 := time.NewTimer(2 * time.Second)
	// 阻塞当前协程，直到从 timer1.C 通道接收到一个值，这表示定时器 timer1 已经失效并触发
	a := <-timer1.C
	fmt.Println("Timer 1 fired %d", a)
	// 如果你需要的仅仅是单纯的等待，使用 time.Sleep 就够了。 使用定时器的原因之一就是，你可以在定时器触发之前将其取消。 例如这样。
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	// 给 timer2 足够的时间来触发它，以证明它实际上已经停止了。
	time.Sleep(2 * time.Second)
	// 第一个定时器将在程序开始后大约 2s 触发， 但是第二个定时器还未触发就停止了。
}

/*
*
定时器 是当你想要在未来某一刻执行一次时使用的 - 打点器
则是为你想要以固定的时间间隔重复执行而准备的。 这里是一个打点器的例子，它将定时的执行，直到我们将它停止。
*/
func testTicker() {
	// 打点器和定时器的机制有点相似：使用一个通道来发送数据。
	//这里我们使用通道内建的 select，等待每 500ms 到达一次的值。
	// ticker := time.NewTicker(500 * time.Millisecond) 创建了一个新的 Ticker，
	// 它会每隔 500 毫秒向其 C 通道发送一个当前时间的值。
	//在 goroutine 中，使用 select 语句等待从 ticker.C 通道接收值，每次接收到一个值时，
	// 都会执行 fmt.Println("Tick at", t) 打印当前时间
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()
	// 打点器可以和定时器一样被停止。 打点器一旦停止，将不能再从它的通道中接收到值。 我们将在运行 1600ms 后停止这个打点器。

	time.Sleep(1600 * time.Millisecond)
	// ticker.Stop() 停止 ticker，一旦停止，就不会再向 ticker.C 通道发送值。
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
func worker(id int, jobs chan int, results chan int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}

}
func testWorkerPool() {
	// 这是 worker 程序，我们会并发的运行多个 worker。 worker 将在 jobs 频道上接收工作，并在 results 上发送相应的结果。 每个 worker 我们都会 sleep 一秒钟，以模拟一项昂贵的（耗时一秒钟的）任务。

	// 为了使用 worker 工作池并且收集其的结果，我们需要 2 个通道。

	// 这里启动了 3 个 worker， 初始是阻塞的，因为还没有传递任务。

	// 这里我们发送 5 个 jobs， 然后 close 这些通道，表示这些就是所有的任务了。
	const numJobs = 5
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	// 最后，我们收集所有这些任务的返回值。 这也确保了所有的 worker 协程都已完成。
	// 另一个等待多个协程的方法是使用WaitGroup。
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

func worker1(id int) {
	fmt.Printf("Worker %d starting\n", id)
	//睡眠一秒钟，以此来模拟耗时的任务。

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}
func testWaitGroup() {
	var wg sync.WaitGroup
	// 启动几个协程，并为其递增 WaitGroup 的计数器。
	// 这个 WaitGroup 用于等待这里启动的所有协程完成。 注意：如果 WaitGroup 显式传递到函数中，则应使用 指针
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		// 避免在每个协程闭包中重复利用相同的 i 值 更多细节可以参考 the FAQ
		// 在 Go 中，闭包会捕获它们引用的外部变量，而不是变量的当前值。
		// 这意味着如果我们直接在 goroutine 中使用循环变量 i，所有的 goroutine 可能都会在循环结束时引用同一个 i 的值
		// 你可能期望输出 1 到 5，但实际上，所有的 goroutine 很可能输出的是相同的值，通常是 6。
		// 这是因为当这些 goroutine 运行时，循环已经结束，i 的值已经变为 6（循环结束时的值）
		i := i
		// 将 worker 调用包装在一个闭包中，可以确保通知 WaitGroup 此工作线程已完成。 这样，worker 线程本身就不必知道其执行中涉及的并发原语。

		go func() {
			// defer wg.Done()：使用 defer 确保在函数返回时调用 wg.Done()，
			// 将 WaitGroup 的计数器减 1，表示当前 goroutine 完成
			defer wg.Done()
			worker1(i)
		}()
	}
	// 阻塞，直到 WaitGroup 计数器恢复为 0； 即所有协程的工作都已经完成。

	wg.Wait()
	// 请注意，WaitGroup 这种使用方式没有直观的办法传递来自 worker 的错误。 更高级的用例，请参见 errgroup package
}

// func Tick(d Duration) <-chan Time,time.Tick 是 Go 语言标准库 time 包中的一个函数，
// 用于创建一个定时器，它会以指定的时间间隔向一个通道发送当前时间,<-chan Time:
// 返回一个只读的时间通道（chan Time），该通道会每隔指定的时间段发送当前时间。
// time.Tick 生成的定时器不能手动停止或清理，它会一直运行直到程序退出。这可能导致资源泄漏。
// 如果需要手动控制定时器的停止，应该使用 time.NewTicker
// ticker := time.NewTicker(200 * time.Millisecond)
// time.NewTicker 创建了一个定时器 ticker，它每隔 200 毫秒会向 ticker.C 通道发送当前时间。
func testLimitSpeed1() {
	//requests := make(chan int, 5)
	//for i := 1; i <= 5; i++ {
	//	requests <- i
	//}
	//close(requests)
	//// limiter 通道每 200ms 接收一个值。 这是我们任务速率限制的调度器。
	//
	//limiter := time.Tick(200 * time.Millisecond)
	//for req := range requests {
	//	<-limiter
	//	fmt.Println("request", req, time.Now())
	//}
	burstyLimiter := make(chan time.Time, 3)
	// 有时候我们可能希望在速率限制方案中允许短暂的并发请求，并同时保留总体速率限制。
	// 我们可以通过缓冲通道来完成此任务。 burstyLimiter 通道允许最多 3 个爆发（bursts）事件。
	//burstyLimiter := time.Tick(200 * time.Millisecond)
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			// 由于 burstyLimiter 的缓冲区大小是 3，当它填满后，发送操作会阻塞，程序停在这一行等待消费者读取数据。
			burstyLimiter <- t
		}
	}()
	// 由于 burstyLimiter 的缓冲区大小是 3，当它填满后，发送操作会阻塞，程序停在这一行等待消费者读取数据
	go func() {
		for t := range burstyLimiter {
			fmt.Println(t)
		}
	}()
	// 让主协程等待一段时间以观察输出
	time.Sleep(2 * time.Second)
	close(burstyLimiter) // 关闭通道以停止 goroutine
}

func testLimitSpeed() {
	//requests := make(chan int, 5)
	//for i := 1; i <= 5; i++ {
	//	requests <- i
	//}
	//close(requests)
	//// limiter 通道每 200ms 接收一个值。 这是我们任务速率限制的调度器。
	//
	//limiter := time.Tick(200 * time.Millisecond)
	//for req := range requests {
	//	<-limiter
	//	fmt.Println("request", req, time.Now())
	//}
	burstyLimiter := make(chan time.Time, 3)
	// 有时候我们可能希望在速率限制方案中允许短暂的并发请求，并同时保留总体速率限制。
	// 我们可以通过缓冲通道来完成此任务。 burstyLimiter 通道允许最多 3 个爆发（bursts）事件。

	// 预填充,每 200ms 我们将尝试添加一个新的值到 burstyLimiter中， 直到达到 3 个的限制。
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.NewTicker(200 * time.Millisecond).C {
			burstyLimiter <- t
		}
	}()

	burstyReq := make(chan int, 5)

	for i := 0; i < 5; i++ {
		burstyReq <- i
	}
	close(burstyReq)
	// 现在，模拟另外 5 个传入请求。 受益于 burstyLimiter 的爆发（bursts）能力，前 3 个请求可以快速完成
	for req := range burstyReq {
		// 第二批请求，由于爆发（burstable）速率控制，我们直接连续处理了 3 个请求，
		// 然后以大约每 200ms 一次的速度，处理了剩余的 2 个请求。
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

func testAtomic() {
	var ops uint64
	// WaitGroup 帮助我们等待所有协程完成它们的工作。

	var wg sync.WaitGroup
	// 我们会启动 50 个协程，并且每个协程会将计数器递增 1000 次。

	for i := 0; i < 50; i++ {
		wg.Add(1)
		// 使用 AddUint64 来让计数器自动增加， 使用 & 语法给定 ops 的内存地址。

		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}
	// 等待，直到所有协程完成。

	wg.Wait()
	// 现在可以安全的访问 ops，因为我们知道，此时没有协程写入 ops， 此外，
	// 还可以使用 atomic.LoadUint64 之类的函数，在原子更新的同时安全地读取它们。

	fmt.Println("ops:", ops)
	// 预计会进行 50,000 次操作。如果我们使用非原子的 ops++ 来增加计数器， 由于多个协程会互相干扰，运行时值会改变，
	// 可能会导致我们得到一个不同的数字。 此外，运行程序时带上 -race 标志，我们可以获取数据竞争失败的详情。
}

// Container 中定义了 counters 的 map ，由于我们希望从多个 goroutine 同时更新它，
// 因此我们添加了一个 互斥锁Mutex 来同步访问。 请注意不能复制互斥锁，如果需要传递这个 struct，应使用指针完成。
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

// 在访问 counters 之前锁定互斥锁； 使用 [defer]（defer） 在函数结束时解锁。

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

// 请注意，互斥量的零值是可用的，因此这里不需要初始化。

func testMutex() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}
	var wg sync.WaitGroup
	// 这个函数在循环中递增对 name 的计数
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}
	// 同时运行多个 goroutines; 请注意，它们都访问相同的 Container，其中两个访问相同的计数器。
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)
	// 等待上面的 goroutines 都执行结束
	wg.Wait()
	fmt.Println(c.counters)
}

type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key   int
	value int
	resp  chan bool
}

// *********************************************************
// 互斥锁方式
//共享状态：使用一个全局或局部变量来保存状态（例如一个 map）。
//互斥锁：使用互斥锁 sync.Mutex 来确保在同一时刻只有一个协程可以访问共享状态。
//锁定与解锁：在访问共享状态之前加锁，访问完毕后解锁。
//优点：使用简单，直观，适用于大多数并发场景。
//缺点：容易引发死锁，特别是在复杂的系统中使用多个互斥锁时。
//协程与通道方式
//单一拥有者：通过创建一个专门的协程来拥有和管理共享状态。
//通信访问：其他协程通过通道 channel 向拥有状态的协程发送读写请求。
//请求与响应：读写请求通过通道发送到拥有状态的协程，拥有状态的协程处理请求并通过通道返回结果。
//优点：避免了使用互斥锁，减少了死锁的风险，适用于复杂的并发系统和多个互斥锁的场景。
//缺点：实现复杂度较高，代码较长，可能不适合简单的并发场景。
//共享内存思想
// 【Go 语言提倡通过通信来共享内存，而不是通过共享内存来通信。】具体来说，这意味着：
//
//避免数据竞争：通过让每个数据仅被一个协程所拥有，确保数据不会在多个协程之间竞争。
//通过通信实现共享内存：其他协程通过通道向拥有数据的协程发送请求，拥有数据的协程处理这些请求并返回结果。
//简化并发控制：这种方法避免了使用互斥锁，减少了死锁和数据竞争的问题，使得并发控制更加简单和安全。
// 1.问题一：为什么不会出现并发问题
//单一拥有者模型：
//
//共享状态（state）仅由一个协程拥有和管理。其他协程无法直接访问 state，只能通过请求-响应的方式间接访问。
//这种方式避免了多个协程同时访问和修改共享状态的问题。
//串行化请求处理：
// 并且通道是
//状态管理协程通过 select 语句串行处理来自 reads 和 writes 通道的请求。
//每次只能处理一个请求，这确保了对共享状态的访问和修改是顺序进行的，避免了并发访问的问题。
//无锁并发：
//
//通过通道来传递请求和响应，而不是使用互斥锁来保护共享状态。这简化了并发控制，避免了死锁和竞争条件等常见的并发问题。
/*
*
在前面的例子中，我们用 互斥锁 进行了明确的锁定， 来让共享的 state 跨多个 Go 协程同步访问。
另一个选择是，使用内建协程和通道的同步特性来达到同样的效果。 Go 共享内存的思想是，
通过通信使每个数据仅被单个协程所拥有，即通过通信实现共享内存。 基于通道的方法与该思想完全一致！
*/
func testChannelRoutineSync() {

	var readOps uint64
	var writeOps uint64
	reads := make(chan readOp)
	writes := make(chan writeOp)
	//在这个例子中，state 将被一个单独的协程拥有。 这能保证数据在并行读取时不会混乱。
	//为了对 state 进行读取或者写入， 其它的协程将发送一条数据到目前拥有数据的协程中，
	//然后等待接收对应的回复。 结构体 readOp 和 writeOp 封装了这些请求，并提供了响应协程的方法。
	go func() {
		//select：在 reads 和 writes 通道上进行选择，处理读取和写入请求。
		//读取请求：从 reads 通道接收一个 readOp，然后在 state 中查找对应的值，并将其发送到 read.resp 通道。
		//写入请求：从 writes 通道接收一个 writeOp，在 state 中更新对应的键值对，并通过 write.resp 通道确认写入成功。
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.value
				write.resp <- true
			}
		}
	}()
	// 和前面的例子一样，我们会计算操作执行的次数。
	// 其他协程将通过 reads 和 writes 通道来发布 读 和 写 请求。
	//    reads := make(chan readOp)
	//    writes := make(chan writeOp)
	// 这就是拥有 state 的那个协程， 和前面例子中的 map 一样，不过这里的 state 是被这个状态协程私有的。
	// 这个协程不断地在 reads 和 writes 通道上进行选择，并在请求到达时做出响应。
	// 首先，执行请求的操作；然后，执行响应，在响应通道 resp 上发送一个值，表明请求成功（reads 的值则为 state 对应的值）。
	for i := 0; i < 100; i++ {
		// 问题二：不会存在多个协程往reads通道发送read的情况吗？不会引起冲突吗？
		// 在 Go 语言中，多个协程可以安全地向同一个通道发送数据。这是因为通道本身是并发安全的，
		// 即它们内部已经实现了必要的同步机制来确保多协程的访问不会产生数据竞争或其他并发问题
		// 通道的设计：
		//
		//Go 的通道设计是并发安全的，这意味着多个协程可以同时向一个通道发送数据，而不需要额外的锁机制。
		//通道内部实现了同步原语，确保在并发场景下的正确性。
		//发送操作：
		//
		//当一个协程执行 reads <- read 时，如果有一个或多个协程在等待接收数据，这个数据会被安全地传递给其中一个等待的接收协程。
		//如果没有接收协程在等待，发送操作会阻塞，直到有接收协程准备好接收数据。
		//为什么不会引起冲突
		//通道是并发安全的：
		//
		//通道内部实现了同步机制，确保多个协程可以安全地向同一个通道发送数据。
		//串行处理请求：
		//
		//尽管有多个协程向 reads 通道发送数据，但接收和处理这些请求的协程会按照顺序处理每个请求。
		//这是通过 select 语句实现的，确保每次只处理一个请求。
		//阻塞与唤醒：
		//
		//当发送操作遇到通道已满的情况时，发送协程会阻塞，直到通道中有空间可以继续发送。这种阻塞机制同样适用于接收操作。
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	//启动 100 个协程通过 reads 通道向拥有 state 的协程发起读取请求。
	//每个读取请求需要构造一个 readOp，发送它到 reads 通道中， 并通过给定的 resp 通道接收结果。
	for i := 0; i < 10; i++ {
		go func() {
			for {
				write := writeOp{
					key:   rand.Intn(5),
					value: rand.Intn(100),
					resp:  make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)
	// 最后，获取并报告 ops 值。
	// 运行这个程序显示这个基于协程的状态管理的例子 达到了每秒大约 80,000 次操作。
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}

func testSort() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)
	//一个 int 排序的例子。

	ints := []int{7, 2, 4}
	sort.Ints(ints)

	fmt.Println("Ints:   ", ints)
	// 我们也可以使用 sort 来检查一个切片是否为有序的。

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}

type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// 为了在 Go 中使用自定义函数进行排序，我们需要一个对应的类型。 我们在这里创建了一个 byLength 类型，它只是内建类型 []string 的别名。
// 我们为该类型实现了 sort.Interface 接口的 Len、Less 和 Swap 方法， 这样我们就可以使用 sort 包的通用 Sort 方法了，
// Len 和 Swap 在各个类型中的实现都差不多， Less 将控制实际的自定义排序逻辑。 在这个的例子中，我们想按字符串长度递增的顺序来排序，
// 所以这里使用了 len(s[i]) 和 len(s[j]) 来实现 Less
func testCustomSort() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}

// panic 意味着有些出乎意料的错误发生。 通常我们用它来表示程序正常运行中不应该出现的错误，
// 或者我们不准备优雅处理的错误。
// 运行程序将会导致 panic： 输出一个错误消息和协程追踪信息，并以非零的状态退出程序。
// 如果你想看到程序尝试创建 /tmp/file 文件，请注释掉第一个panic
func testPanic() {
	panic("a problem")
	//panic 的一种常见用法是：当函数返回我们不知道如何处理（或不想处理）的错误值时，
	//中止操作。 如果创建新文件时遇到意外错误该如何处理？这里有一个很好的 panic 示例。
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}

// 这是一个 panic 函数
// 必须在 defer 函数中调用 recover。 当跳出引发 panic 的函数时，defer 会被激活， 其中的 recover 会捕获 panic
func mayPanic() {
	panic("a problem")
}

// Go 通过使用 recover 内置函数，可以从 panic 中 恢复recover 。
// recover 可以阻止 panic 中止程序，并让它继续执行。
//
// 在这样的例子中很有用：当其中一个客户端连接出现严重错误，服务器不希望崩溃。
// 相反，服务器希望关闭该连接并继续为其他的客户端提供服务。
// 实际上，这就是Go的 net/http 包默认对于 HTTP 服务器的处理。
func testRecover() {
	// recover 的返回值是在调用 panic 时抛出的错误。

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()
	mayPanic()
	// 这行代码不会执行，因为 mayPanic 函数会调用 panic。 main 程序的执行在 panic 点停止，
	// 并在继续处理完 defer 后结束。
	fmt.Println("After mayPanic()")

}

// 我们给 fmt.Println 一个较短的别名， 因为我们随后会大量的使用它。
var p = fmt.Println

/*
*
面的 len 以及索引工作在字节级别上。 Go 使用 UTF-8 编码字符串，因此通常按原样使用。 如果您可能使用多字节的字符，
则需要使用可识别编码的操作。 详情请参考 strings, bytes, runes and characters in Go。
*/
func testString() {
	p("contains: ", strings.Contains("test", "es"))
	p("count", strings.Count("aaaa", "a"))
	p("preFix", strings.HasPrefix("test", "te"))
	p("suffix", strings.HasSuffix("test", "st"))
	p("Index", strings.Index("test", "es"))
	p("Join", strings.Join([]string{"a", "b"}, "-"))
	p("repeat", strings.Repeat("a", 5))
	p("replace", strings.Replace("food", "o", "0", -1))
	p("Split:     ", strings.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", strings.ToLower("TEST"))
	p("ToUpper:   ", strings.ToUpper("test"))
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
	//Contains:   true
	//Count:      2
	//HasPrefix:  true
	//HasSuffix:  true
	//Index:      1
	//Join:       a-b
	//Repeat:     aaaaa
	//Replace:    f00
	//Replace:    f0o
	//Split:      [a b c d e]
	//ToLower:    test
	//ToUpper:    TEST
	//Len:  5
	//Char: 101
}

type point struct {
	x, y int
}

func testType() {
	p := point{1, 2}
	//struct1: {1 2}
	fmt.Printf("struct1: %v\n", p)

	// 如果值是一个结构体，%+v 的格式化输出内容将包括结构体的字段名。
	// struct2: {x:1 y:2}
	fmt.Printf("struct2: %+v\n", p)

	// %#v 根据 Go 语法输出值，即会产生该值的源码片段。
	// struct3: main.point{x:1, y:2}
	fmt.Printf("struct3: %#v\n", p)

	// 需要打印值的类型，使用 %T。
	// type: main.point
	fmt.Printf("type: %T\n", p)

	// 格式化布尔值很简单。
	// bool: true
	fmt.Printf("bool: %t\n", true)

	// 格式化整型数有多种方式，使用 %d 进行标准的十进制格式化。
	// int: 123，Decimal
	fmt.Printf("int: %d\n", 123)

	// 这个输出二进制表示形式。
	// bin: 1110
	fmt.Printf("bin: %b\n", 14)

	// 输出给定整数的对应字符。
	// char: !
	fmt.Printf("char: %c\n", 33)

	// %x 提供了十六进制编码。
	// hex: 1c8
	fmt.Printf("hex: %x\n", 456)

	// 同样的，也为浮点型提供了多种格式化选项。 使用 %f 进行最基本的十进制格式化。
	// float1: 78.900000
	fmt.Printf("float1: %f\n", 78.9)

	// %e 和 %E 将浮点型格式化为（稍微有一点不同的）科学记数法表示形式。
	// float2: 1.234000e+08
	// float3: 1.234000E+08
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)

	// 使用 %s 进行基本的字符串输出。
	// str1: "string"
	fmt.Printf("str1: %s\n", "\"string\"")

	// 像 Go 源代码中那样带有双引号的输出，使用 %q。
	// str2: "\"string\""
	fmt.Printf("str2: %q\n", "\"string\"")

	// 和上面的整型数一样，%x 输出使用 base-16 编码的字符串， 每个字节使用 2 个字符表示。
	// str3: 6865782074686973
	fmt.Printf("str3: %x\n", "hex this")

	// 要输出一个指针的值，使用 %p。
	// pointer: 0xc0000ba000
	fmt.Printf("pointer: %p\n", &p)

	// 格式化数字时，您经常会希望控制输出结果的宽度和精度。 要指定整数的宽度，请在动词 “%” 之后使用数字。
	// 默认情况下，结果会右对齐并用空格填充。
	// width1: |    12|   345|
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	// 你也可以指定浮点型的输出宽度，同时也可以通过 宽度.精度 的语法来指定输出的精度。
	// width2: |  1.20|  3.45|
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	// 要左对齐，使用 - 标志。
	// width3: |1.20  |3.45  |
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// 你也许也想控制字符串输出时的宽度，特别是要确保他们在类表格输出时的对齐。 这是基本的宽度右对齐方法。
	// width4: |   foo|     b|
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	// 要左对齐，和数字一样，使用 - 标志。
	// width5: |foo   |b     |
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	// 到目前为止，我们已经看过 Printf 了， 它通过 os.Stdout 输出格式化的字符串。 Sprintf 则格式化并返回一个字符串而没有任何输出。
	// sprintf: a string
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	// 你可以使用 Fprintf 来格式化并输出到 io.Writers 而不是 os.Stdout。
	// io: an error
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
func testReg() {
	//true

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// 上面我们是直接使用字符串，但是对于一些其他的正则任务，
	// 你需要通过 Compile 得到一个优化过的 Regexp 结构体。

	r, _ := regexp.Compile("p([a-z]+)ch")
	// 该结构体有很多方法。这是一个类似于我们前面看到的匹配测试。
	// true
	fmt.Println(r.MatchString("peach"))

	// 查找匹配的字符串。
	// peach
	fmt.Println(r.FindString("peach punch"))

	// 这个也是查找首次匹配的字符串， 但是它的返回值是，匹配开始和结束位置的索引，而不是匹配的内容。
	// idx: [0 5]
	fmt.Println("idx:", r.FindStringIndex("peach punch"))

	// Submatch 返回完全匹配和局部匹配的字符串。 例如，这里会返回匹配 p([a-z]+)ch 和 ([a-z]+) 的信息。
	// [peach ea]
	fmt.Println(r.FindStringSubmatch("peach punch"))

	//类似的，这个会返回完全匹配和局部匹配位置的索引。
	//[0 5 1 3]
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// 带 All 的这些函数返回全部的匹配项， 而不仅仅是首次匹配项。例如查找匹配表达式全部的项。
	// [peach punch pinch]
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// All 同样可以对应到上面的所有函数。
	// all: [[0 5 1 3] [6 11 7 9] [12 17 13 15]]
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	// 这些函数接收一个非负整数作为第二个参数，来限制匹配次数。
	// [peach punch]
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// 上面的例子中，我们使用了字符串作为参数， 并使用了 MatchString 之类的方法。
	// 我们也可以将 String 从函数名中去掉，并提供 []byte 的参数。
	// true
	fmt.Println(r.Match([]byte("peach")))

	//创建正则表达式的全局变量时，可以使用 Compile 的变体 MustCompile 。
	//MustCompile 用 panic 代替返回一个错误 ，这样使用全局变量更加安全。
	// regexp: p([a-z]+)ch
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	// regexp 包也可以用来将子字符串替换为为其它值。
	// a <fruit>
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
	// Func 变体允许您使用给定的函数来转换匹配的文本。
	// a PEACH
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
