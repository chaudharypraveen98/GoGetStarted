package main

import (
	"fmt"

	"rsc.io/quote"
)

func sum_sub(x int, y int) (z1 int, z2 int) {
	defer fmt.Println("block successfully executed")
	z1 = x + y
	z2 = x - y
	fmt.Println("before block")
	return
}

func my_function(test5 func(int) int) int {
	return test5(6) * 7
}
func return_func(x string) func() {
	return func() { fmt.Println(x) }
}
func change_first(slice []int) {
	slice[0] = 1000
}

// it changes the value by accessing address
func change_value1(str *string) {
	*str = "changes"
}

//it creates the copy not altered the stored value
func change_value2(str string) {
	str = "still same"
}

// struct
type Point struct {
	x int32
	y int32
}
type Circle1 struct {
	radius float64
	center *Point
}

// Interface
// It means any struct that have area method and return type is float 64
type Shape interface {
	area() float64
}
type Circle struct {
	radius float64
	center Point
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

type Rect struct {
	l float64
	b float64
}

func (r Rect) area() float64 {
	return r.b * r.l
}

// cool thing about the interface is that we can use it as param, variable or return type
func get_area(s Shape) float64 {
	return s.area()
}

// change value of struct by pointer, we don't need to defrence it *
func change_value_point(pt *Point) {
	pt.x = 100
}

// student struct
type Student struct {
	name  string
	marks []int
	age   uint32
}

func (s Student) get_average() float32 {
	sum := 0
	for i := 0; i < len(s.marks); i++ {
		sum += s.marks[i]
	}
	return float32(sum) / float32(len(s.marks))
}

// unused variable or modules are errors
func main() {
	var name string
	name = "Ram"
	var number int8
	number = 67
	var test = 67
	fmt.Println("Hello World", name, number)
	fmt.Println(quote.Go())
	// %T refers to the type of the variable
	fmt.Printf("%T\n", test)
	// shorthand for intializing the varibale
	raju := "Raj"
	fmt.Printf("%T\n", raju)
	// every variable in go has a default value
	var ram bool
	fmt.Println(ram)
	// formatting
	// boolean and decimal
	fmt.Printf("%d\n", 10)
	fmt.Printf("%t\n", true)
	fmt.Printf("%v\n", 10.9990)
	// q for printing with quotes
	fmt.Printf("%q\n", "time")
	// text justify
	fmt.Printf("%-9q is cool\n", "ram")
	fmt.Printf("%.3f is cool\n", 9.45678)
	var text = fmt.Sprintf("%07.3f is cool", 9.45678)
	fmt.Println(text)
	// operators 0r and not
	test1 := (true || false) && !false
	fmt.Printf("%t\n", test1)
	// loops if AND for
	age := 10
	if age > 5 && age < 8 {
		fmt.Println("and greater than 5")
	} else if age > 8 && age < 10 {
		fmt.Println("i am in between 8 and 10")
	} else if age >= 10 {
		fmt.Println("iam greater than10")
	} else {
		fmt.Println("i am less than 5")
	}
	for x := 0; x < 3; x++ {
		fmt.Println(x)
	}
	// switch
	// we can't check string if its integer value
	ans := -1
	switch ans {
	case 1, -1:
		fmt.Println("i am case 1 and -1")
	case 2:
		fmt.Println("i am case 1")
	default:
		fmt.Println("i am default case")
	}
	ques := 0
	switch {
	case ques < 0:
		fmt.Println("less than 0")
	case ques > 0:
		fmt.Println("greater than zero")
	default:
		fmt.Println("autaully zeri")
	}
	// arrays and slicing
	var a [5]int
	fmt.Println(a)
	// intializig with values
	b := [3]int{1, 2, 3}
	fmt.Println(b)

	fmt.Println(b[:1])
	// []int type is for slice and [3]int for array
	fmt.Printf("%T %T\n", b[:1], b)

	//withput inrializig array
	var c []int = []int{1, 2, 3, 4, 5, 6, 5, 4, 4}
	fmt.Println(c, len(c))

	// can't append list directly
	// d := append(c, 6)
	c = append(c, 6)
	fmt.Println(c)

	// printing unique elements
	for i, element := range c {
		for j := i + 1; j < len(c); j++ {
			element1 := c[j]
			if element == element1 {
				fmt.Println("same found")
			}
		}
	}

	// dict or map in go
	var mp map[string]int = map[string]int{
		"apple": 1,
		"mango": 2,
	}
	// another way to describe the map
	kp := make(map[string]int)
	delete(mp, "apple")
	fmt.Println(len(mp), kp)

	// checking if key exists or not
	val, ok := mp["apple"]
	fmt.Println(val, ok)
	// updating map
	mp["lichi"] = 1
	fmt.Println(len(mp), kp)
	// it prints the default value 0
	fmt.Println(mp["grapes"])

	//calling function
	sum, sub := sum_sub(5, 8)
	fmt.Printf("sum is %d and subtract is %d\n", sum, sub)
	fmt.Println("i am coool", sum)

	//inline function and passing parameter in single line
	test_func := func(x int) int {
		return x * 7
	}(8)
	fmt.Println(test_func)

	//it will pass test_func1 to myfunction or myfunction parameter function test5=test_func1
	test_func1 := func(x int) int {
		return 7 * x
	}
	fmt.Println(my_function(test_func1))

	// returning function and calling at the same time
	return_func("hello")()

	// list follows call by reference by address
	// data types like string , bol, number and float makes a copy
	var x [2]int = [2]int{3, 4}
	y := x
	y[0] = 100
	fmt.Println(x, y)
	// above output [3 4] [100 4]

	//slice and map are mutable
	var x1 []int = []int{2, 3, 4, 5}
	fmt.Println(x1)
	change_first(x1)
	fmt.Println(x1)

	// pointers and changing value by accessing address

	x5 := 7
	y5 := &x5
	fmt.Println(x5, y5)
	*y5 = 8
	fmt.Println(x5, y5)
	to_change := "my name"
	change_value1(&to_change)
	fmt.Println(to_change)
	change_value2(to_change)
	fmt.Println(to_change)

	var p1 Point = Point{1, 2}
	fmt.Println(p1, p1.x)

	p2 := Point{x: 5}
	fmt.Println(p2)

	// we can also define pointer to point
	p3 := &Point{y: 5}
	fmt.Println(p3)
	change_value_point(p3)
	fmt.Println(p3)
	c1 := Circle1{5.65, &Point{4, 5}}
	fmt.Println(c1.center.x)

	student1 := Student{"ram", []int{23, 34, 99, 78, 65}, 34}
	fmt.Println(student1.get_average())

	c3 := Circle{5.65, Point{4, 5}}
	r3 := Rect{4, 6}
	shapes := []Shape{c3, r3}
	for _, shape := range shapes {
		fmt.Println("Area")
		fmt.Println(shape.area())
		// let use interface as a param type
		fmt.Println("generating area by using interface as a param")
		fmt.Println(get_area(shape))
	}
}
