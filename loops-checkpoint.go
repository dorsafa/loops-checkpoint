package main

import "fmt"

func main() {
	x :=0

	fmt.Println("Enter a number :")
	fmt.Scan(&x)

	fmt.Println("multiple of 2 and 5:")
	multipleOf2and5(x)
	fmt.Println("multiple of 2 and 3 and 5 :")
	multipleOf2and3and5(x)
}

func multipleOf2and5(n int)  {

	for i := 2 ;i<=n ; i++ {
		b := false
		if i%2 == 0 && i%5 == 0 {
			b = true
		}
		if b == true {
			fmt.Println(i)
		}

}}

func multipleOf2and3and5(n int){
	for i := 2 ;i<n ; i++ {
		b := false
		if i%2 == 0 && i%3==0 && i%5 == 0 {
			b =true
		}
		if b == true {
			fmt.Println(i)
		}

}}