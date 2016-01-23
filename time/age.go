// Copyright 2015, Stève Sfartz
// Licensed under the MIT License

package main

import (
	"time"
	"fmt"
)

// This function highlights a bug with time.AddDate invocation
// Published at http://play.golang.org/p/u6zHYRatdA
func main() {

	today := time.Date(2016, 01, 21, 0, 0, 0, 0, time.UTC)

	fmt.Println("suppose today is  : ", today)

	birthDate := time.Date(1971, 12, 24, 0, 0, 0, 0, time.UTC)

	fmt.Println("if you're born on : ", birthDate)

	// Buggy call, a month in missing with go 1.5 on windows 7 64 bits
	// Reproducible with days from 21 to 31, December any year
	age := today.AddDate(birthDate.Year() * -1, int(birthDate.Month()) * -1, birthDate.Day() * -1)

	// /!\ Should be one month more
	fmt.Println("then you are      : ", age.Year(), " years, ", int(age.Month()), " months, and ", age.Day(), ", days.")
	fmt.Println("/!\\ 44 years, 0 months and 28 days is the correct answer")

	// Testing symetry
	// ? do we recover if we add back ?
	back := age.AddDate(birthDate.Year(), int(birthDate.Month()), birthDate.Day())

	fmt.Println("\nchecking symetry  : ", back == today)
	fmt.Println("confirm this is today: ", back)
}


