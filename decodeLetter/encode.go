package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func decodePassword(s string) string {
	pass := make([]int, len(s)+1)
	for i, v := range s {
		if v == '=' {
			pass[i+1] = pass[i]
		} else if v == 'L' {
			if pass[i] != 0 {
				pass[i+1] = 0
			} else {
				pass[i+1] = -1
			}
		} else if v == 'R' {
			pass[i+1] = pass[i] + 1
		}

		if pass[i+1] < 0 {
			pass[i+1]++
			for j := i; j >= 0; j-- {
				if s[j] == 'L' {
					if pass[j] > pass[j+1] {
						break
					} else {
						pass[j]++
					}
				} else if s[j] == 'R' {
					if pass[j] < pass[j+1] {
						break
					}
				} else if s[j] == '=' {
					if pass[j] == pass[j+1] {
						break
					} else {
						pass[j] = pass[j+1]
					}
				}
			}
		}
	}

	result := ""
	for _, v := range pass {
		result += strconv.Itoa(v)
	}

	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter input: ")
	input, _ := reader.ReadString('\n')
	input = strings.ToUpper(strings.TrimSpace(input))
	resp := decodePassword(input)
	fmt.Println(resp)
}
