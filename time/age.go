// Copyright 2015, Stève Sfartz
// Licensed under the MIT License

package main

import (
	"time"
	"fmt"
)

func main() {
	birthDate := time.Date(1971, 12, 24, 18, 0, 0, 0, time.UTC)

	fmt.Println("you're born on ", birthDate)

	today := time.Now()

	fmt.Println("today is ", today)

	age := today.AddDate(-1971, -12, -24)

	fmt.Println("congrats, you are ", age.Year(), ", ", int(age.Month()), " months, and ", age.Day(), ", days.")

	again := today.AddDate(birthDate.Year() * -1, int(birthDate.Month()) * -1, birthDate.Day() * -1)

	fmt.Println("again, you are ", again.Year(), ", ", int(again.Month()), " months, and ", again.Day(), ", days.")

}
