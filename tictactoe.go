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
	used := []string{}

	for !win {
		//LEt Users know which player should go
		fmt.Println("Player " + strconv.Itoa(player) + ":")

		//Make a slice to check if the location is free

		freeChoice := false

		var location string

		//Read the players location and check if it is free
		for !freeChoice {
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

			if chooseAgain {
				continue
			} else {
				used = append(used, loc)
				location = loc
				freeChoice = true
			}
		}

		var dot string

		if player == 1 {
			dot = "X"
		} else {
			dot = "O"
		}

		switch location {
		case "1":
			TL = "_" + dot + "_"
		case "2":
			TM = "_" + dot + "_"
		case "3":
			TR = "_" + dot + "_"
		case "4":
			ML = "_" + dot + "_"
		case "5":
			MM = "_" + dot + "_"
		case "6":
			MR = "_" + dot + "_"
		case "7":
			BL = " " + dot + " "
		case "8":
			BM = " " + dot + " "
		case "9":
			BR = " " + dot + " "
		default:
			fmt.Println("Please choose a valid location on the grid")
		}

		grid := "   |   |  " + "\n" + TL + "|" + TM + "|" + TR + "\n" + "   |   |  " + "\n" + ML + "|" + MM + "|" + MR + "\n" + "   |   |   " + "\n" + BL + "|" + BM + "|" + BR + "\n"
		fmt.Println(grid)

		//Make conditions if the game is a win
		if ML == MM && MM == MR && MR != "_ _" {
			fmt.Println("Congratulations Player " + strconv.Itoa(player) + " YOU WIN!")
			return
		} else if TL == TM && TM == TR && TR != "_ _" {
			fmt.Println("Congratulations Player " + strconv.Itoa(player) + " YOU WIN!")
			return
		} else if BL == BM && BM == BR && BR != "   " {
			fmt.Println("Congratulations Player " + strconv.Itoa(player) + " YOU WIN!")
			return
		} else if 

		// Change player to next user
		if player == 1 {
			player = 2
		} else {
			player = 1
		}
	}
}
