package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	getInnerContent("2[2[ab]c3[d]]")
}

// Case 1
// Input format: 2[a]3[bc]
// Output format: aabcbcbc

// Case 2
// Input format: 2[2[ab]c3[d]]
// Output format: ababcdddababcdd

// Case 3
// Input format: a3[b]d
// Output format: abbbd

// Posso fazer recursivo. Dessa forma a parte mais interna da string é printada antes.

// Algoritmo
// Pego o index do primeiro middleRegEx encontrado
// Se index -1, printo toda a str original
// Printo de 0 até esse index
// Chamo a função recursiva passando a string do indexInicio até o indexFim dessa middleRegEx encontrada
// Chamo a função recursiva passando do indexFim ao fim da string original

// PROBLEMA
// No caso 2, o algoritmo não encontra o ninho mais externo... Como fazer para que ele reconheça?

func decompress(str string, times int) {
	var middleRegEx = regexp.MustCompile("[0-9]+\\[[0-9A-Za-z]+\\]")
	// var contentRegEx = regexp.MustCompile("\\[[0-9A-Za-z]+\\]")
	var numberRegEx = regexp.MustCompile("[0-9]+\\[")

	// var innerStrings = middleRegEx.FindAllString(str, -1)
	var firstNestIndex = middleRegEx.FindStringIndex(str)
	// fmt.Println(firstNestIndex)

	if firstNestIndex == nil {
		for i := 0; i < times; i++ {
			fmt.Print(str)
		}
	} else {
		var previous = str[0:firstNestIndex[0]]
		fmt.Print(previous)

		var nest = str[firstNestIndex[0]:firstNestIndex[1]]
		var value = numberRegEx.FindString(nest)
		var numberOfTimes, _ = strconv.Atoi(value[:len(value)-1])

		var final = nest[2 : len(nest)-1]
		decompress(final, numberOfTimes)

		decompress(str[firstNestIndex[1]:len(str)], 1)
	}
}

func getInnerContent(str string) {
	var bracketCount = 0
	var start, end = 0, 0
	for i := 0; i < len(str); i++ {
		if str[i] == '[' {
			if bracketCount == 0 {
				start = i
			}
			bracketCount++
		}
		if str[i] == ']' {
			if bracketCount == 1 {
				end = i
			}
			bracketCount--
		}
	}
	fmt.Println(str[start+1 : end])
	getInnerContent(str[start+1 : end])
}
