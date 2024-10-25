package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var memoryLoad []string
	for i := 0; i < 1000000000000; i++ {
		randomHex, err := GenerateRandomHex()
		if err != nil {
			panic(err)
		}
		memoryLoad = append(memoryLoad, randomHex)
	}
	fmt.Println(memoryLoad)
}
