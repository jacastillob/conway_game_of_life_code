package main

import (
	"fmt"
	"math/rand"
	"time"
)

//types
type generation [10][10]int

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
func print(m generation) {

	for i := 0; i < len(m); i++ {
		fmt.Println("")
		for j := 0; j < len(m[i]); j++ {
			fmt.Print(m[i][j])
		}
	}

}

//check neighbours

func copy(a, m generation, x, y int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			m[i+x][j+y] = a[i][j]
		}
	}
}

//generation-> generation 0 , generation 1

func main() {

	gSeed := intializeGeneration()

	//gSeed := generation{{1, 2}, {3, 4}}
	print(gSeed)
	//print(5, 5, matrix)

	//var result = math.Sqrt(5)

}
