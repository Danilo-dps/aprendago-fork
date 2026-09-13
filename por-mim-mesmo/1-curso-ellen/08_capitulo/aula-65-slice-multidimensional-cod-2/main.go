package main

import (
	"fmt"
)

func main() {
	ss := [][][][]int{
		{
			{
				{1, 2, 3, 4, 5, 6},
			},
			{
				{10, 20, 30, 40, 50},
			},
		},

		{
			{
				{2, 4, 6, 8, 10},
			},
			{
				{3},
			},
		},
	}

	fmt.Println(ss[1][0][0][2])
	fmt.Println(ss)
}
