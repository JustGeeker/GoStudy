package main

import "fmt"

//批量申请常量时，如果某一行没有赋值，默认与上一行一致
const (
	n1 = 100
	n2
	n3 = "sdsd"
	n4
)

//iota 初始为0，每新增一行常量数值加1
const (
	s1 = iota
	s2
	s3
	s4
)

const (
	a1 = iota
	a2
	_
	a3
)
const (
	b1 = iota
	b2
	b3 = 100
	b4
	b5 = iota
	b6
)

func main() {
	fmt.Println(n1, n2, n3, n4)         //100 100 sdsd sdsd
	fmt.Println(s1, s2, s3, s4)         //0 1 2 3
	fmt.Println(a1, a2, a3)             //0 1 3
	fmt.Println(b1, b2, b3, b4, b5, b6) //0 1 100 100 4 5
	for i := 0; i < 4; i++ {
		fmt.Println(i * i)
	}

}
