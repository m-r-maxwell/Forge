package main

import (
	"fmt"
	forge "forge/cmd"
)

func main() {
	fmt.Println(`
	_______   ______   .______        _______  _______
	|   ____| /  __  \  |   _  \      /  _____||   ____|
	|  |__   |  |  |  | |  |_)  |    |  |  __  |  |__
	|   __|  |  |  |  | |      /     |  | |_ | |   __|
	|  |     |  '--'  | |  |\  \----.|  |__| | |  |____
	|__|      \______/  | _| '._____| \______| |_______|
	`)
	fmt.Println("Starting to forge....")
	forge.Execute()
}
