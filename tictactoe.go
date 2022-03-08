package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	TL := "_ _"
	TM := "_ _"
	TR := "_ _"
	ML := "_ _"
	MM := "_ _"
	MR := "_ _"
	BL := "   "
	BM := "   "
	BR := "   "

	grid := "   |   |  " + "\n" + TL + "|" + TM + "|" + TR + "\n" + "   |   |  " + "\n" + ML + "|" + MM + "|" + MR  + "\n" + "   |   |   " + "\n"  + BL + "|" + BM + "|" + BR + "\n"
	fmt.Println(grid)

	win := false
	player := 1

	for !win {
		fmt.Println("Player " + strconv.Itoa(player) + ":")
		scanner := bufio.NewScanner(os.Stdin)

		scanner.Scan()

		loc := scanner.Text()
		
		fmt.Println(loc)

		var dot string

		if player == 1 {
			dot = "X"
		} else {
			dot = "O"
		}

			switch loc {
			case TL :
				TL = "_" + dot + "_"
			case TM :
				TM = "_" + dot + "_"
			case TR :
				TR = "_" + dot + "_"
			case ML :
				ML = "_" + dot + "_"
			case MM :
				MM = "_" + dot + "_"
			case MR :
				MR = "_" + dot + "_"
			case BL :
				BL = " " + dot + " "
			case BM :
				BM = " " + dot + " "
			case BR :
				BR = " " + dot + " "
			}

			if player == 1 {
				player = 2
			} else {
				player = 1
			}

			fmt.Println(grid)



	
	}
}
