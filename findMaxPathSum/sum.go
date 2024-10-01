package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func FindMaxPathSum(input [][]int) int {
	if len(input) == 0 {
		return 0
	}
	for i := len(input) - 2; i >= 0; i-- {
		for j := 0; j < len(input[i]); j++ {
			if input[i+1][j] > input[i+1][j+1] {
				input[i][j] += input[i+1][j]
			} else {
				input[i][j] += input[i+1][j+1]
			}
		}
	}
	return input[0][0]
}

func main() {
	byteValue, err := os.ReadFile("../files/hard.json")
	if err != nil {
		panic(err)
	}
	input := make([][]int, 0)
	json.Unmarshal(byteValue, &input)
	fmt.Println(FindMaxPathSum(input))
}
