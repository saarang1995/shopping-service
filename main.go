package main

import (
	"fmt"
	"shopping-service/shopping"
)
type Student struct {
	name string
	marks float32
}

func (s *Student) getAnnualMarks() float32 {
	return s.marks * 12;
}

func main() {
	// var numbers = [3]int{1, 2, 3}
	// fmt.Println(numbers)
	// // append(numbers, 2)
	
	// var numbers2 = make([]int,4)
	// fmt.Println(numbers2)

	// numbers2 = append(numbers2, 2)
	// fmt.Println(numbers2)
	// fmt.Println(numbers2[0])
	// fmt.Println(numbers2[0:3])

	// var slice []int = []int{}
	// slice = append(slice, 10)
	// fmt.Println(slice)

	// for i,v := range numbers2{
	// 	fmt.Println(i, v)
	// }

	
	// var studentIndex map[int]Student = make(map[int]Student)

	// studentIndex[1] = Student{
	// 	name: "test",
	// 	marks: 20.3,
	// }

	// studentIndex[2] = Student{
	// 	name: "test2",
	// 	marks: 20.3,
	// }

	// for i, student:= range studentIndex {
	// 	fmt.Println(i, studentIndex[i], student.getAnnualMarks())
	// }

	circle := Circle{x: 0, y: 2, radius: 20}
	rectangle := Rectangle{width: 2,height: 4}
	for i := 0; i < 10; i++ {
		printArea(&circle)
		printArea(&rectangle)
	}
	
	// fmt.Println(circle.area())
	// fmt.Println(rectangle.area())
}

func printArea(shape Shape) {
	fmt.Println(shape.area())
}

type Shape interface {
	area() float64
}

type Circle struct {
	x,y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (circle *Circle) area() float64 {
	return 200
}

func (rectangle *Rectangle) area() float64 {
	return 300
}


func matrices() {
	// var matrix [][]int = make([][]int, 4)

	// for i := 0; i< len(matrix); i++ {
	// 	matrix[i] = make([]int, 4)
	// }
	// matrix[0][0] = 200

	// for i, row := range matrix {
	// 	fmt.Println(row, i)
	// 	// for j, _ := range row {
	// 	// 	fmt.Println(matrix[i][j])
	// 	// }
	// }
}

func dealWithShopping() {
	addNewUser("User1")
	addNewUser("User2")
	addNewUser("User3")
	for _, v := range shopping.GetAllUsers() {
		fmt.Println(*v.Id, v.Name)
		fmt.Println("Orders:")
		for _, order := range *v.Orders {
			fmt.Println(*order.Id, order.Name, order.Amount)
		}
	}
	
}

func addNewUser(name string) {
	newUser := shopping.AddUser(shopping.User{Name: name})

	shopping.AssignOrderToUserById(*newUser.Id,"Milk" ,2)
	shopping.AssignOrderToUserById(*newUser.Id, "Coffee", 1)
}