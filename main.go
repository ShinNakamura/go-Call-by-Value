package main

import (
	"fmt"
)

type Person struct {
	Name       string
	Age        int
	SchoolName *string
}

func main() {
	i := 1
	s := "foo"
	arr := [2]int{1, 2}

	primitiveNoMod(i, s, arr)
	fmt.Println(i, s, arr) // prints 1 foo [1 2]
	primitiveMod(&i, &s, &arr)
	fmt.Println(i, s, arr) // prints 0 bar [0 2]

	schoolName := "GoSchool"
	a := Person{"A", 43, &schoolName}

	structNoMod(a)
	fmt.Println(schoolName)    // GoSchool
	fmt.Println(a)             // {A 43 (memory address)}
	fmt.Println(*a.SchoolName) // GoSchool

	structMod(&a)
	fmt.Println(schoolName)    // GoSchool
	fmt.Println(a)             // {B 15 (memory address)}
	fmt.Println(*a.SchoolName) // GoHighSchool

	m := map[int]string{1: "X", 2: "Y"}
	// map はそもそもポインター型
	mapMod(m)
	fmt.Println(m) // map[1:_X 2:Y 3:Z]

	/*
		スライスは
			pointer value
			length
			capacity
		の3つのフィールドを持つ変数。
		Call-by-Valueな関数呼び出しでそのまま渡すと、
		上記3つのフィールドがコピーされる。

		pointer の値 ＝ メモリー上のアドレスが変わらない変更を
		している間は元の値に変更が加えられる。
		しかし、append などしてメモリーアロケーションが発生すると、
		pointer 値が「コピーの側で」書き換わる。
		そうすると、そこから先の変更は元の値とは独立した値に対して行われる
	*/
	slice := []int{1, 2, 3}

	sliceModCopy(slice)
	fmt.Println(slice) // [99 2 3] // not append 80
}

// コピーを受け取る
func primitiveNoMod(i int, s string, arr [2]int) {
	// 以下は全部コピーを変更している
	i = 0
	s = "bar"
	arr[0] = 0
}

// ポインターを受け取る
func primitiveMod(i *int, s *string, arr *[2]int) {
	// ポインターをデリファレンスし、
	// 呼び出し元の値を変更する
	*i = 0
	*s = "bar"
	arr[0] = 0
}

// コピーを受け取る
func structNoMod(p Person) {
	// 以下は全部コピーを変更している
	p.Name = "B"
	p.Age = 15
	schoolName := "GoHighSchool"
	p.SchoolName = &schoolName
}

// ポインターを受け取る
func structMod(p *Person) {
	// ポインターをデリファレンスし、
	// 呼び出し元の値を変更する
	p.Name = "B"
	p.Age = 15
	schoolName := "GoHighSchool"
	p.SchoolName = &schoolName
}

func mapMod(m map[int]string) {
	m[1] = "_X"
	m[3] = "Z"
}

func sliceModCopy(s []int) {
	s[0] = 99
	s = append(s, 80) // change address
}
