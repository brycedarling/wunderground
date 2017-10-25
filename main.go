package main

import (
	"fmt"
	"github.com/brycedarling/wunderground"
	"os"
	"path/filepath"
)

func displayUsage() {
	program := filepath.Base(os.Args[0])
	fmt.Printf("usage: %s command \"location\"\n", program)
	fmt.Println("command must be one of tendayforecast or astronomy")
	fmt.Println("location must be a comma separated City, State value")
	fmt.Println("for example:")
	fmt.Printf("%s tendayforecast \"Bozeman, MT\"\n", program)
	fmt.Println("or...")
	fmt.Printf("%s astronomy \"Asheville, NC\"\n", program)
}

func displayTenDayForecast(cityAndState string) {
	url, err := wunderground.GetTenDayForecastUrl(cityAndState)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	forecast, err := wunderground.GetTenDayForecastData(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", forecast)
}

func displayAstronomy(cityAndState string) {
	url, err := wunderground.GetAstronomyUrl(cityAndState)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	astronomy, err := wunderground.GetAstronomyData(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", astronomy)
}

func main() {
	if len(os.Args) != 3 {
		displayUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	cityAndState := os.Args[2]

	if command == "tendayforecast" {
		displayTenDayForecast(cityAndState)
	} else if command == "astronomy" {
		displayAstronomy(cityAndState)
	} else {
		displayUsage()
	}
}
