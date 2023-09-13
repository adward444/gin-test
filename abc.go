package main

import "fmt"

type EatWhat interface {
	EatMeat(data interface{}) error
	LikeSleep() bool
}

//定义类型结构1
type me struct {
}

//定义类型结构2
type he struct {
}

//定义类型结构3
type she struct {
}

//定义的me类型实现接口
func (I *me) EatMeat(data interface{}) error {
	fmt.Println("I like eat meat:!!!!: data：", data)
	return nil
}
func (I *me) LikeSleep() bool {
	return true
}

//定义的he类型实现接口
func (H *he) EatMeat(data interface{}) error {
	fmt.Println("he does not like meat!!!data:", data)
	return nil
}
func (H *he) LikeSleep() bool {
	return true
}

//定义的she类型实现接口
func (S *she) EatMeat(data interface{}) error {
	fmt.Println("she also likes meat!!!Data：", data)
	return nil
}
func (S *she) LikeSleep() bool {
	return false
}

func main() { //实例化me结构体
	fmew := new(me)
	fhew := new(he)
	fsew := new(she)
	//声明一个EatWhat的接口
	var ew EatWhat //将接口赋值结构体的实例化，即me类型
	//ew = fmew.EatMeat("dataaaaaa")
	//fmt.Println(ew.LikeSleep())
	//ew = fhew.EatMeat("hhhhhhh")
	//fmt.Println(ew.LikeSleep())
	//ew = fsew.EatMeat("ssssssss")
	//fmt.Println(ew.LikeSleep())

	ew = fmew
	ew.EatMeat("aaaaaa")
	fmt.Println(ew.LikeSleep())


	ew = fhew
	ew.EatMeat("bbbbb")
	fmt.Println(ew.LikeSleep())

	ew = fsew
	ew.EatMeat("ccccc")
	fmt.Println(ew.LikeSleep())
}
