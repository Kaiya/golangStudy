package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human eat")
}
func (this *Human) Walk() {
	fmt.Println("Human walk")
}

type SuperMan struct {
	Human //SuperMan类继承了Human类的方法
	level int
}

//重定义父类的方法Eat
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan Eat...")
}

//子类的新方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan Fly...")
}
func main() {
	h := Human{"zs", "femal"}
	h.Eat()
	h.Walk()
	//定义一个子类对象
	s := SuperMan{Human{"li4", "female"}, 88}
	//或者
	var s2 SuperMan
	s2.name = "f"
	s2.sex = "male"
	s2.level = "66"

	s.Walk() //父类的方法
	s.Fly()
	s.Eat()
}
