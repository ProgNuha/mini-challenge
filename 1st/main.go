package main

import "fmt"

func main(){
	var n int
	fmt.Print("Input number : ")
    fmt.Scanln(&n)
	
	for i := 1; i <= n; i++ {
		if(i % 5 == 0 && i % 3 == 0) {
			fmt.Println("fizz buzz") 
		} else if(i % 5 == 0) {
			fmt.Println("buzz") 
		} else if(i % 3 == 0) {
			fmt.Println("fizz") 
		} else {
			fmt.Println(i) 
		}
	}

}