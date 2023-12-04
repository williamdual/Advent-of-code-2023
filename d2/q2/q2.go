package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile() string {
	var filepath string = "q2.txt"

	file, err := os.Open(filepath)
	checkError(err)
	scanner := bufio.NewScanner(file)
	var text string
	for scanner.Scan() {
		text += scanner.Text()
		text += "\n"
	}
	return text
}

func getCubeDict(cubes []string) map[string]int {
	colorR := regexp.MustCompile(" ")
	cubeMap := make(map[string]int)
	for i := 0; i < len(cubes); i++ {
		pair := colorR.Split(cubes[i], 2)
		temp, err := strconv.Atoi(pair[0])
		checkError(err)
		cubeMap[pair[1]] = temp

	}
	return cubeMap
}

func decode(gamesRaw string) int {
	lines := strings.Split(gamesRaw, "\n")
	var sumOfPowers int = 0
	//cubes

	cubeColors := [3]string{"red", "green", "blue"}
	//regex
	gameR := regexp.MustCompile("Game \\d+: ")
	gameIdRightR := regexp.MustCompile(": ")
	gameIdLeftR := regexp.MustCompile("Game ")
	handR := regexp.MustCompile("; ")
	cubeR := regexp.MustCompile(", ")

	for i := 0; i < len(lines)-1; i++ {
		line := lines[i]

		var game string = gameR.Split(line, 2)[1]
		gameId, err := strconv.Atoi(gameIdLeftR.Split((gameIdRightR.Split(line, 2)[0]), 2)[1]) //gameId is an int
		checkError(err)

		hands := handR.Split(game, 20)

		cubeDict := map[string]int{"red": 0, "blue": 0, "green": 0}
		for j := 0; j < len(hands); j++ {
			cubes := cubeR.Split(hands[j], 10)
			tempDict := getCubeDict(cubes)
			for k := 0; k < len(cubeColors); k++ {
				color := cubeColors[k]
				if tempDict[color] > cubeDict[color] && tempDict[color] != 0 {
					cubeDict[color] = tempDict[color]
				}
			}

		}
		tempPower := 1
		for j := 0; j < len(cubeColors); j++ {
			color := cubeColors[j]
			tempPower *= cubeDict[color]

		}
		fmt.Printf("%d: %d\n", gameId, tempPower)
		sumOfPowers += tempPower

	}
	return sumOfPowers
}

func main() {
	var input string = readFile()
	ans := decode(input)
	fmt.Println(ans)
}
