package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var TL, TM, TR, ML, MM, MR string = "_1_", "_2_", "_3_", "_4_", "_5_", "_6_"
	var BL, BM, BR string = " 7 ", " 8 ", " 9 "

	grid := "   |   |  " + "\n" + TL + "|" + TM + "|" + TR + "\n" + "   |   |  " + "\n" + ML + "|" + MM + "|" + MR + "\n" + "   |   |   " + "\n" + BL + "|" + BM + "|" + BR + "\n"
	

	//When a player gets two in a row
	win := false
	plays := 1
	used := []string{}

	scanner1 := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter player 1's name:")
	scanner1.Scan()
	player1 := scanner1.Text()

	scanner2 := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter player 2's name:")
	scanner2.Scan()
	player2 := scanner2.Text()

	fmt.Println(grid)
	fmt.Println("Let's Go!")

	name := player1

	for !win {

		//Let Users know which player should go
		fmt.Println(name + ":")

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

		if name == player1 {
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
			//Middle Row
			fmt.Println("Congratulations " + name + " YOU WIN!")
			return
		} else if TL == TM && TM == TR && TR != "_ _" {
			//Top Row
			fmt.Println("Congratulations " + name + " YOU WIN!")
			return
		} else if BL == BM && BM == BR && BR != "   " {
			//Bottom Row
			fmt.Println("Congratulations " + name + " YOU WIN!")
			return
		} else if TR == MM && []rune(MM)[1] == []rune(BL)[1] && []rune(MM)[1] != ' ' {
			//right to left diag
			fmt.Println("Congratulations " + name + " YOU WIN!")
			return
		} else if TL == MM && []rune(MM)[1] == []rune(BR)[1] && []rune(MM)[1] != ' ' {
			//Left to right diag
			fmt.Println("Congratulations " + name + " YOU WIN!")
			return
		} else if TL == ML && []rune(ML)[1] == []rune(BL)[1] && []rune(ML)[1] != ' ' {
			//Left Verticle
			fmt.Println("Congratulations " + name + " YOU WIN!")
			return
		} else if TM == MM && []rune(MM)[1] == []rune(BM)[1] && []rune(MM)[1] != ' ' {
			//Middle Verticle
			fmt.Println("Congratulations " + name + " YOU WIN!")
			return
		} else if TR == MR && []rune(MR)[1] == []rune(BR)[1] && []rune(MR)[1] != ' ' {
			//Right Verticle
			fmt.Println("Congratulations " + name + " YOU WIN!")
			return
		}

		// Change player to next user
		if name == player1 {
			name = player2
		} else {
			name = player1
		}

		if plays >= 9 {
			fmt.Println("This game was a draw! Try again")
			return
		}
		plays++
	}
}
