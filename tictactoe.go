package main

import "fmt"

func main() {
	var cross = [3][3]string{}
	player := "x"
	var row, column int
	var p int

	for {
		p = 0
		fmt.Println("Enter row from 0 to 2: ")
		fmt.Scanln(&row)

		if row > 2 || row < 0 {
			fmt.Println("Invalid row entered")
			continue
		}

		fmt.Println("Enter column from 0 to 2: ")
		fmt.Scanln(&column)

		if column > 2 || column < 0 {
			fmt.Println("Invalid column entered")
			continue
		}

		if cross[row][column] == "" {
			cross[row][column] = player
		} else {
			fmt.Println("Row: ", row, " and column: ", column, " is filled.")
			continue
		}

		fmt.Println(cross[0])
		fmt.Println(cross[1])
		fmt.Println(cross[2])

		if (cross[0][0] == cross[0][1] && cross[0][1] == cross[0][2] && cross[0][2] == player) || (cross[1][0] == cross[1][1] && cross[1][1] == cross[1][2] && cross[1][2] == player) || (cross[2][0] == cross[2][1] && cross[2][1] == cross[2][2] && cross[2][2] == player) {
			fmt.Println("Player: ", player, " is winner")
			return
		} else if (cross[0][0] == cross[1][0] && cross[1][0] == cross[2][0] && cross[2][0] == player) || (cross[0][1] == cross[1][1] && cross[1][1] == cross[2][1] && cross[2][1] == player) || (cross[0][2] == cross[1][2] && cross[1][2] == cross[2][2] && cross[2][2] == player) {
			fmt.Println("Player: ", player, " is winner")
			return
		} else if (cross[0][0] == cross[1][1] && cross[1][1] == cross[2][2] && cross[2][2] == player) || (cross[0][2] == cross[1][1] && cross[1][1] == cross[2][0] && cross[2][0] == player) {
			fmt.Println("Player: ", player, " is winner")
			return
		}

		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if cross[i][j] != "" {
					p++
				}
			}
		}

		if p == 9 {
			fmt.Println("Match drawn")
			return
		}

		if player == "x" {
			player = "o"
		} else {
			player = "x"
		}
	}

}
