package main

import "fmt"

func main() {
	var num, count int
	var done bool
	for !done {
		fmt.Print("Input number: ")
		fmt.Scan(&num)
		if num >= 12307 {
			fmt.Println("Number should be <12307")
		} else {
			for num < 12307 {
				count++
				if num < 0 {
					num *= -1
				} else if num%7 == 0 {
					num *= 39
				} else if num%9 == 0 {
					num = num*13 + 1
					continue
				} else {
					num = (num + 2) * 3
				}
				if num%13 == 0 && num%9 == 0 {
					fmt.Println("service error")
					done = true
					break
				} else {
					num++
				}
			}
			if !done {
				fmt.Println("Result:", num)
				done = true
			}
			fmt.Println("Iteration:", count)
		}
	}
}
