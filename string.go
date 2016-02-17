package main

import "fmt"

func main() {
	str1 := "hello world my friend!"
	chars := []byte(str1)
	for _, c := range chars {
		fmt.Println(string(c))
	}

	fmt.Println("pos 0 =>", string(chars[0]))
	fmt.Println("pos 1 =>", string(chars[1]))
	fmt.Println("end   =>", string(chars[len(chars)-1]))
}
