package main

import "fmt"

func main() {
	total := 270
	bonus := 5
	countBonus := 0
	count := 0
	for countBonus < total{
		count ++
		countBonus = countBonus + bonus
		fmt.Println(count," ")
	}
}