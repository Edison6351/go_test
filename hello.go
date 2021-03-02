package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// array_test()
	// array_point()
	array_string()
	fmt.Println("Please visit http://127.0.0.1:12345/")
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		s := fmt.Sprintf("你好，世界! -- Time:%s", time.Now().String())
		fmt.Fprintf(w, "%v\n", s)
		log.Printf("%v\n", s)
	})
	if err := http.ListenAndServe(":12345", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

//数组相关练习
func array_test() {
	var a [3]int
	var b = [...]int{1, 2, 3}
	var c = [...]int{2: 3, 1: 2}
	var d = [...]int{1, 2, 4: 5, 6}
	fmt.Println(a, b, c, d)
}

func array_point() {
	//数组指针
	var a = [...]int{1, 2, 3}
	var b = &a
	fmt.Println(a[0], a[1])
	fmt.Println(b[0], b[1])

	for i, v := range b {
		fmt.Println(i, v)
	}
}
<<<<<<< HEAD

//字符串指针
func array_string() {
	var a = "1234567"
	for i := range a {
		fmt.Println(i)
	}
}
=======
>>>>>>> 4bcf30bd3fd995fd11048245ef9be3968fe28faa
