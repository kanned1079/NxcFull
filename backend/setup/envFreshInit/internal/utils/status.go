package utils

import "fmt"

func ShowOK() {
	fmt.Printf("\033[1;32m[OK]\033[0m\033[0m ")
}

func ShowFailure() {
	fmt.Printf("\033[31m[FAILURE]\033[0m\033[0m ")
}

func ShowAttention() {
	fmt.Printf("\033[1;38;5;214m[ATTENTION]\033[0m ")
}
