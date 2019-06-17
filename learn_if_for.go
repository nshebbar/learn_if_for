package main

import "fmt"

func main() {
	a := 10
	if a > 5 {
		fmt.Println("a is bigger than 5")
	} else {
		fmt.Println("a is less than or equal to 5")
	}
	b := 3
	for i:=0; i<10; i++ {
		if i > 7 {
			break
		} 
		if i ==b  {
			continue
		}
		fmt.Println(i)
	}
s := "Hello World!"
for k, v := range s {
	fmt.Println(k, v, string(v))
}
}