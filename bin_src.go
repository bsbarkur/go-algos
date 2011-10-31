/*
 * Binary search
 * run it as ./bin_src <input_list>
 * where <input_list> is the array over which we do binary search
 */
package main

import ( 
	"fmt"
	"flag"
	"strconv"
)

func binSearch (list []int, key int) int {
	low := 0
	high := len(list)-1
	
	for low <= high {
		mid := low + (high - low) / 2
		if key == list[mid] {
		     return mid
		} else if key < list[mid] {
		     high = mid-1
	     	} else {
		     low = mid+1
	     	}
	}
     	return -1 
}

func main() {
	flag.Parse()
	fmt.Println("This is the list over which we conduct binary search \n", 
				flag.Args())
	input_list := make([]int, flag.NArg())
	var list_i, i int
	for j := 0; j < flag.NArg(); j++ { 	
		list_i = j
		input_list[list_i], _ = strconv.Atoi(flag.Arg(j))
	}
	fmt.Println("Please enter the key you want to search for")
    	fmt.Scan(&i)
	if index := binSearch(input_list,i); index == -1 {
    		fmt.Println("Key", i, "was not found in the list")
	} else {
    		fmt.Println("Key", i, "found in the list with index", index)
	}
}
