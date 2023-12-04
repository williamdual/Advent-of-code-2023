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
	var filepath string = "q1.txt"

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
	var sumOfIds int = 0
	//cubes
	var redMax int = 12
	var greenMax int = 13
	var blueMax int = 14
	cubeMaxMap := map[string]int{"red": redMax, "green": greenMax, "blue": blueMax}
	cubeColors := [3]string{"red", "green", "blue"}
	//regex
	gameR := regexp.MustCompile("Game \\d+: ")
	gameIdRightR := regexp.MustCompile(": ")
	gameIdLeftR := regexp.MustCompile("Game ")
	handR := regexp.MustCompile("; ")
	cubeR := regexp.MustCompile(", ")

	for i := 0; i < len(lines)-1; i++ { // replace 5 with len(lines)-1 when done
		var gamePossible bool = true
		line := lines[i]
		//Here line = Game 1: 12 blue, 15 red, 2 green; 17 red, 8 green, 5 blue; 8 red, 17 blue; 9 green, 1 blue, 4 red

		var game string = gameR.Split(line, 2)[1]
		gameId, err := strconv.Atoi(gameIdLeftR.Split((gameIdRightR.Split(line, 2)[0]), 2)[1]) //gameId is an int
		checkError(err)

		//Here game = 12 blue, 15 red, 2 green; 17 red, 8 green, 5 blue; 8 red, 17 blue; 9 green, 1 blue, 4 red
		//fmt.Println(game)

		hands := handR.Split(game, 20)
		//Here hands = ["12 blue, 15 red, 2 green", " 17 red, 8 green, 5 blue", " 8 red, 17 blue", " 9 green, 1 blue, 4 red"]
		//fmt.Printf("	Hands: %q\n", hands)
		for j := 0; j < len(hands); j++ {
			cubes := cubeR.Split(hands[j], 10)
			//Here cubes = ["12 blue", "15 red", "2 green"]
			cubeDict := getCubeDict(cubes)
			//Here cubeDict = map[blue:12 green:2 red:15]
			for k := 0; k < len(cubeColors); k++ {
				color := cubeColors[k]
				if cubeDict[color] > cubeMaxMap[color] {
					gamePossible = false
					break
				}
			}
			if !gamePossible {
				break
			}
		}
		if gamePossible {
			sumOfIds += gameId
		}

	}
	return sumOfIds
}

func main() {
	var input string = readFile()
	ans := decode(input)
	fmt.Println(ans)
}
