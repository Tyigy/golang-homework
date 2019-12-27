package main

import "fmt"

func Multiply(a, b int) (int){
	var s int = 0
	for i:=b; i>0; i--{
		s = s + a
	}
	return s
}

func Fibonacci1(n int) int{
	fib1:=1
	fib2:=1
	fib_sum:=0
	i:=0
	for (i<n-2){
		fib_sum = fib1+fib2
		fib1=fib2
		fib2=fib_sum
		i=i+1
	}
	return fib2	
}

func Fibonacci2(n int) int{
	if n == 1{
		return 1
	}
	if n == 2{
		return 1
	}else{
		return Fibonacci2(n-1)+Fibonacci2(n-2)
	}
}

func Bubble_sort(a[] float64){
	var buff float64
	for i:=1; i<len(a); i++{
		for j:=0; j<len(a)-i; j++{
			if a[j]>a[j+1]{
				buff=a[j]
				a[j]=a[j+1]
				a[j+1]=buff
			}
		}
	}
}

func main() {
	fmt.Println(Multiply(3, 2))
	fmt.Println(Fibonacci1(8))
	fmt.Println(Fibonacci2(6))

	var a = [] float64 {12.9, 5.7, 5.6, 2.1, 7.9}
	Bubble_sort(a)
	fmt.Println(a)
}
