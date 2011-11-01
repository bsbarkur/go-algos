/*
 * Reversing a simple string
 * Lessons:
 * 	- Strings in go are immutable
 */

package main

import (
	"fmt"
)

func reverse_str(orig_str string) string {
	ls := []int(orig_str) // make array
	str_len := len(ls)
	
	for i:=0 ; i< str_len /2 ; i++ {
		// Swap :O
		ls[i],ls[str_len-i-1] = ls[str_len-i-1],ls[i]
	}
    	return string(ls) // In-built function to convert to string
}

func main() {
	fmt.Println("Enter the string you want to be reversed without any spaces")
	var orig_str string 
	fmt.Scan(&orig_str)
	reversed_str := reverse_str(orig_str)
	fmt.Println(reversed_str)	
}
