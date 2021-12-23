package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var originUnit string
var originValue float64

var shouldConvertAgain string

var err error

var errInvalidArguments = errors.New("invalid arguments")
var errReadingInput = errors.New("error reading input")

func main() {
	if len(os.Args) != 2 {
		printError(errInvalidArguments)
	}

	originUnit = strings.ToUpper(os.Args[1])

	for {
		fmt.Print("What is the current temperature in " + originUnit + " ? ")
		_, err = fmt.Scanln(&originValue)
		if err != nil {
			printError(errReadingInput)
		}
		if originUnit == "C" {
			convertCToFahrenheit(originValue)
			//adding convert to K to this if
			convertCToKelvin(originValue)

		}
		//added a if for F
		if originUnit == "F" {
			convertFToCelsius(originValue)
			convertFToKelvin(originValue)
		}
		//added a if for K
		if originUnit == "K" {
			convertKToCelsius(originValue)
			convertKTofahrenheit(originValue)
		}
		/* Dont think i need a else im just going to have more if's
		 else {
			convertFToCelsius(originValue)
			convertCToKelvin(originValue)
		 }
		*/
		//^	added convertToKelvin(originValue) ^

		fmt.Print("Would you like to convert another temperature ? (y/n) ")
		_, err = fmt.Scanln(&shouldConvertAgain)
		if err != nil {
			printError(errReadingInput)
		}

		if strings.ToUpper(strings.TrimSpace(shouldConvertAgain)) != "Y" {
			fmt.Println("Good bye!")
			break
		}
	}
}

func printError(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

//changed to convertFToCelsius from convertToCelsius
func convertFToCelsius(value float64) {
	convertedValue := (value - 32) * 5 / 9
	fmt.Printf("%v C = %.0f F\n", value, convertedValue)
}

//Addded F-K converter
func convertFToKelvin(value float64) {
	convertedValue := (value-32)*5/9 + 273.15
	fmt.Printf("%v K = %.0f F\n", value, convertedValue)
}

//changed to convertCToFahrenheit from convertTOFahrenheit
func convertCToFahrenheit(value float64) {
	convertedValue := (value * 9 / 5) + 32
	fmt.Printf("%v F = %.0f C\n", value, convertedValue)
}

//  add c-k func
func convertCToKelvin(value float64) {
	convertedValue := (value + 273.15)
	fmt.Printf("%v K = %.0f C\n", value, convertedValue)
}

// added k-f converter
func convertKTofahrenheit(value float64) {
	convertedValue := (value-273.15)*9/5 + 32
	fmt.Printf("%v F = %.0f K\n", value, convertedValue)
}

// added K-C converter
func convertKToCelsius(value float64) {
	convertedValue := (value - 273.15)
	fmt.Printf("%v C = %.0f K\n ", value, convertedValue)
}
