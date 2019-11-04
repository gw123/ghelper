package main

import "fmt"

func main() {
	prefix := "20191101-AIK0200"

	for i := 1; i <= 97; i++ {
		if i< 10{
			fmt.Printf("%s0%d\n", prefix, i)
		}else{
			fmt.Printf("%s%d\n", prefix, i)
		}

	}

}
