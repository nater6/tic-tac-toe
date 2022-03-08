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

	grid := "   |   |  " + "\n" + TL + "|" + TM + "|" + TR + "\n" + "   |   |  " + "\n" + ML + "|" + MM + "|" + MR + "\n" + "   |   |   " + "\n" + BL + "|" + BM + "|" + BR + "\n"
	fmt.Println(grid)

	win := false
	player := 1

	for !win {
		//LEt Users know which player should go
		fmt.Println("Player " + strconv.Itoa(player) + ":")

		//Make a slice to check if the location is free
		used := []string{}
		freeChoice := false

		var location string

		//Read the players location and check if it is free
		for !freeChoice {
			fmt.Println(used)
			//Create a scanner to read the users input from the terminal
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			loc := scanner.Text()

			chooseAgain := false

			//Create a for loop to check if the location is used
			for i := 0; i < len(used); i++ {
				if used[i] == loc {
					fmt.Println("Please choose a free space on the grid")
					chooseAgain = true
					continue
				}
			}

			

			if !chooseAgain {
				used = append(used, loc)
				location = loc
				freeChoice = true
				fmt.Println(used)
			}
		}

		var dot string

		if player == 1 {
			dot = "X"
		} else {
			dot = "O"
		}

		switch location {
		case "TL":
			TL = "_" + dot + "_"
		case "TM":
			TM = "_" + dot + "_"
		case "TR":
			TR = "_" + dot + "_"
		case "ML":
			ML = "_" + dot + "_"
		case "MM":
			MM = "_" + dot + "_"
		case "MR":
			MR = "_" + dot + "_"
		case "BL":
			BL = " " + dot + " "
		case "BM":
			BM = " " + dot + " "
		case "BR":
			BR = " " + dot + " "
		default:
			fmt.Println("Please choose a valid location on the grid")
		}

		if player == 1 {
			player = 2
		} else {
			player = 1
		}

		grid := "   |   |  " + "\n" + TL + "|" + TM + "|" + TR + "\n" + "   |   |  " + "\n" + ML + "|" + MM + "|" + MR + "\n" + "   |   |   " + "\n" + BL + "|" + BM + "|" + BR + "\n"
		fmt.Println(grid)

	}
}
