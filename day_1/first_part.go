package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func firstFunction(x int) int {

	value := x / 3

	return value - 2
}

func secondFunction(x int, totalFuel int) int {

	addedFuel := firstFunction(x)
	totalFuel = addedFuel + totalFuel

	if addedFuel > 8 {
		return secondFunction(addedFuel, totalFuel)
	}

	return totalFuel
}

func main() {

	dataFile, err := ioutil.ReadFile("input.data")
	check(err)

	dataPoints := string(dataFile)

	inputValues := strings.Split(dataPoints, "\n")

	// fmt.Println(inputValues)
	// fmt.Println(reflect.TypeOf(dataPoints).String())
	var sum int
	for _, element := range inputValues {

		intgerValue, err := strconv.Atoi(element)
		check(err)
		// sum = sum + firstFunction(intgerValue)
		sum = sum + secondFunction(intgerValue, 0)
	}
	fmt.Println(sum)
}
