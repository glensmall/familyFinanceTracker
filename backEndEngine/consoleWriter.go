package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

/*
	Function: printInfo
	Author	: Glen Small
	Date	: 30th December 2020

	Params	: string : The text to render

	Returns : nothing

	Purpose : will output the given text with [INFO] in Grey at the end.

*/
func printInfo(data string) {

	currentTime := time.Now()

	txt := color.New(color.FgHiWhite).SprintfFunc()

	fmt.Printf("[%s] - %s - %s\n", txt("INFO"), currentTime.Format("2006-01-02 15:04:05"), data)

}

/*
	Function: printSuccess
	Author	: Glen Small
	Date	: 30th December 2020

	Params	: string : The text to render

	Returns : nothing

	Purpose : will output the given text with date-time at the start and [OK] in green at the end.

*/
func printSuccess(data string) {

	currentTime := time.Now()

	txt := color.New(color.FgHiGreen).SprintfFunc()

	fmt.Printf("[%s] - %s - %s\n", txt("OK"), currentTime.Format("2006-01-02 15:04:05"), data)

}

/*
	Function: printWarning
	Author	: Glen Small
	Date	: 30th December 2020

	Params	: string : The text to render

	Returns : nothing

	Purpose : will output the given text with date-time at the start and [WARNING] in yellow at the end.

*/
func printWarning(data string) {

	currentTime := time.Now()

	txt := color.New(color.FgHiYellow).SprintfFunc()

	fmt.Printf("[%s] - %s - %s\n", txt("WARNING"), currentTime.Format("2006-01-02 15:04:05"), data)
}

/*
	Function: printError
	Author	: Glen Small
	Date	: 30th December 2020

	Params	: string : The text to render

	Returns : nothing

	Purpose : will output the given text with date-time at the start and [FAILED] in red at the end.

*/
func printError(data string) {

	currentTime := time.Now()

	txt := color.New(color.FgHiRed).SprintfFunc()

	fmt.Printf("[%s] - %s - %s\n", txt("FAILED"), currentTime.Format("2006-01-02 15:04:05"), data)

}
