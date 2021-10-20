package main

import (
	"fmt"
	"math/rand"
	"time"
)

//types
type generation [20][20]int

//intialize
func intializeGeneration() generation {

	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 1

	newGen := generation{}

	for i := 0; i < len(newGen); i++ {
		for j := 0; j < len(newGen[i]); j++ {
			newGen[i][j] = rand.Intn(max-min+1) + min
		}
	}
	return newGen
}

// print
func print(g generation) {
	fmt.Println("")
	for i := 0; i < len(g); i++ {
		fmt.Println("")
		for j := 0; j < len(g[i]); j++ {
			fmt.Print(g[i][j])
		}
	}
}

//check neighbours
func countNeighbors(x, y int, g generation) int {
	count := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {

			col := (x + i + len(g)) % len(g)
			row := (x + j + len(g[i])) % len(g[i])

			count += g[col][row]

		}
	}
	return count - g[x][y]
}

//compute generation
func computeGeneration(g generation) generation {
	nextGen := generation{}
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {

			neighbors := countNeighbors(i, j, g)
			state := g[i][j]
			//RULE 4
			if state == 0 && neighbors == 3 {
				nextGen[i][j] = 1
				//RULE 1,3
			} else if state == 1 && (neighbors < 2 || neighbors > 3) {
				nextGen[i][j] = 0
			} else {
				nextGen[i][j] = state
			}

		}
	}
	return nextGen
}

//compute generations
func computeGenerations() {
	gSeed := intializeGeneration()
	print(gSeed)
	i := 0
	for {
		print(gSeed)
		newGeneration := computeGeneration(gSeed)
		gSeed = newGeneration

		time.Sleep(time.Second * 1)
		i++
	}
}

func main() {

	go computeGenerations()
	fmt.Println("Press the enter key to stop the game of life")
	fmt.Scanln()
}
