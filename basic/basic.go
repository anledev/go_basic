package basic

import (
	"fmt"
	"time"
)

func demoBasic() {
	//bien, kieu du lieu, func, cau truc du lieu, oop, go routine, channel
	//simple

	//demo1Variable()
	//demo2Func()
	//demo3DataStructure()
	//demo4OOP()
	//demo5GoRoutine()
	demo6GoChannel()
}

func demo6GoChannel() {
	c := make(chan int)
	go func() {
		c <- 100 // đẩy giá trị 100 vào channel
	}()

	go func() {
		fmt.Println(<-c)
	}()

	time.Sleep(time.Second)
}

/*demo 5 - starts*/

func demo5GoRoutine() {
	//go fun1
	go g1()

	go func() {
		fmt.Println("2")
	}()

	time.Sleep(time.Second)
}

func g1() {
	fmt.Println("1")
}

/*demo 5 - ends*/

/*demo 4 - starts*/
type Student struct {
	//class
	firstName string
	lastName  string
	email     string
}

//value receiver method - bản sao, không thay đổi giá trị bản gốc
func (s Student) getEmail() string {
	s.email = "frankdev@gmail.com"
	return s.email
}

//pointer receiver method - thay đổi giá trị bản gốc
func (s *Student) getEmailPointer() string {
	s.email = "frankdev@gmail.com"
	return s.email
}

func demo4OOP() {
	st := Student{
		firstName: "An",
		lastName:  "Le",
		email:     "anle@gmail.com",
	}
	st.firstName = "Frank"
	fmt.Println(st)
	fmt.Println("=====Value Receiver Method=======")
	e := st.getEmail()
	fmt.Println(e)
	fmt.Println(st.email)

	fmt.Println("=====Pointer Receive Method=======")
	ePointer := st.getEmailPointer()
	fmt.Println(ePointer)
	fmt.Println(st.email)
}

/*demo 4 - ends*/

func demo3DataStructure() {
	//slice - mở rộng array
	//array - fix số lượng phần tử - giống list: cần thì append
	//slice, map, array

	//slice
	fmt.Println("===============Slice===============")
	s := make([]string, 0)
	s = append(s, "Android")
	s = append(s, "iOS")
	s = append(s, "NodeJS")
	s = append(s, "Golang")
	fmt.Println(s)

	//map
	fmt.Println("===============Map===============")
	m := make(map[string]int)
	//key-value
	m["key1"] = 30
	m["key2"] = 10
	fmt.Println(m["key1"])
	fmt.Println(m["key2"])
	fmt.Println(m)

	//array
	fmt.Println("===============Array===============")
	arr := [1]string{}
	arr[0] = "Chelsea"

	fmt.Println(arr)
}

func demo2Func() {
	total := sum(1, 4)
	fmt.Println(total)
}

func sum(a int, b int) int {
	return a + b
}

func demo1Variable() {
	//method 1
	var email string = "abc@gmail.com"
	var name = "Anle"
	var age = 29
	fmt.Println(email)
	fmt.Println(name)
	fmt.Printf("%T", email)
	fmt.Println()
	fmt.Printf("%V", email)
	fmt.Println()
	fmt.Printf("%T", age)
	fmt.Println()

	//method 2
	fullName := "Le Thanh An"
	fmt.Println(fullName)
	fmt.Printf("%T", fullName)
	fmt.Println()

	var number = 10
	var numberFloat = 10.123
	fmt.Println(number)
	fmt.Printf("%T", number)
	fmt.Println()
	fmt.Println(numberFloat)
	fmt.Printf("%T", numberFloat)
	fmt.Println()
}
