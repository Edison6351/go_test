package main

import (
	"fmt"
)

func main() {
	fmt.Println(Hello())
	// fmt.Println("Please visit http://127.0.0.1:12345/")
	// http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	// 	s := fmt.Sprintf("你好，世界! -- Time:%s", time.Now().String())
	// 	fmt.Fprintf(w, "%v\n", s)
	// 	log.Printf("%v\n", s)
	// })
	// if err := http.ListenAndServe(":12345", nil); err != nil {
	// 	log.Fatal("ListenAndServe:", err)
	// }
}

// //数组相关练习
// func array_test() {
// 	var a [3]int
// 	var b = [...]int{1, 2, 3}
// 	var c = [...]int{2: 3, 1: 2}
// 	var d = [...]int{1, 2, 4: 5, 6}
// 	fmt.Println(a, b, c, d)
// }

// func array_point() {
// 	//数组指针
// 	var a = [...]int{1, 2, 3}
// 	var b = &a
// 	fmt.Println(a[0], a[1])
// 	fmt.Println(b[0], b[1])

// 	for i, v := range b {
// 		fmt.Println(i, v)
// 	}
// }

// func values_test() {
// 	fmt.Println("go" + "lang")
// 	fmt.Println("1+1=", 1+1)
// 	fmt.Println(true && false)
// 	fmt.Println(true || false)
// 	fmt.Println(!true)
// }

// func variables_test() {
// 	var a = "initial"
// 	fmt.Println(a)
// 	var b, c int = 1, 2
// 	fmt.Println(b, c)
// 	var d = true
// 	fmt.Println(d)
// 	var e int
// 	fmt.Println(e)

// 	f := "apple"
// 	fmt.Println(f)
// }

func Hello() string {
	return "Hello, go"
}
