package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func printTime(currentTime float64) {
	//this is pretty gross... I should clean this up later
	hours := int(currentTime / (60 * 60))
	minutes := int(currentTime) % (60 * 60) / 60
	seconds := float64(int(currentTime)%60) + (currentTime - float64(int(currentTime)))

	fmt.Printf("\t|Time Passed: %2d : %2d : %5.2f|\n\r", hours, minutes, seconds)
	fmt.Printf("\t|____________________________|\r")
	print("\033[A\r")

}

func printBanner() {
	fmt.Printf("\n\t___________HH/MM/SS___________\n")
}

func printGoodbye() {
	printTime(0)
	print("\033[2B\r")
	fmt.Printf("\n\tTimer Complete!\n\n")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Error: expected time argument\n")
		return
	}
	var direction bool
	var help bool
	var timeToCount float64 = 0.0
	var hours int
	var minutes int
	var seconds int

	//by default we'll take whatever was the first position arg to be our time in seconds
	timeToCount, _ = strconv.ParseFloat(os.Args[1], 32)
	if timeToCount != 0 { //we had a positional so lets permute args so we can still use the flags lib
		args := os.Args[1:] //remember changing the slice changes the underlying array
		optind := 0
		for i := range args {
			if args[i][0] == '-' {
				tmp := args[i]
				args[i] = args[optind]
				args[optind] = tmp
				optind++
			}
		}
	}

	flag.BoolVar(&direction, "u", false, "count up or down, default is down")
	flag.BoolVar(&help, "h", false, "display usage")
	flag.IntVar(&seconds, "s", 0, "time in seconds")
	flag.IntVar(&minutes, "m", 0, "time in minutes")
	flag.IntVar(&hours, "hr", 0, "time in hours")

	//check our args
	flag.Parse()

	if help { //print usage and exit
		flag.Usage()
		return
	}

	if timeToCount == 0.0 { //no positional was given so we'll parse our others
		timeToCount = float64(hours*3600 + minutes*60 + seconds)
	}

	var cT float64                      //current time either starts at max or 0
	var incrementF float64              //increment to count by
	incrementT := 1 * time.Second / 100 //increment to sleep by
	incrementF = 1.0 / 100

	printBanner() //banner sits over clock

	if direction { //counting up not down
		cT = 0.0
		for cT < timeToCount {
			printTime(cT)
			time.Sleep(incrementT)
			cT += incrementF
		}
	} else { //else counting down
		cT = timeToCount
		for cT > 0.0 {
			printTime(cT)
			time.Sleep(incrementT)
			cT -= incrementF
		}
	}

	//print goodbye message
	printGoodbye()
	return
}
