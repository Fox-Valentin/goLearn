package point_demo

import "fmt"

func TestPoint() {
	var count int
	count = 20
	var countPoint *int
	countPoint = &count
	fmt.Printf("count value = %d\n", count)
	fmt.Printf("countPoint value = %x\n", countPoint)
	fmt.Printf("*countPoint value = %d\n", *countPoint)
	fmt.Printf("count address = %x\n", &count)
}

func TestPointArr() {
	a, b := 1, 2
	pointArr := [...]*int{&a, &b}
	arr := [...]int{1, 2, 3, 4}
	fmt.Println("pointArr = ", pointArr)
	fmt.Println("arr point =", &arr)

}
