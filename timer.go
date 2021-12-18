package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func printTime(currentTime float64, msg bool) {
	//this is pretty gross... I should clean this up later
	hours := int(currentTime / (60 * 60))
	minutes := int(currentTime) % (60 * 60) / 60
	seconds := float64(int(currentTime)%60) + (currentTime - float64(int(currentTime)))

	if msg {
		fmt.Printf("\t|Time Passed: %2d : %2d : %5.2f   |\n\r", hours, minutes, seconds)
	} else {
		fmt.Printf("\t|Time Remaining: %2d : %2d : %5.2f|\n\r", hours, minutes, seconds)
	}
	fmt.Printf("\t|_______________________________|\r")
	print("\033[A\r")

}

func printBanner() {
	fmt.Printf("\n\t_____________HH/MM/SS____________\n")
}

func printGoodbye(cT float64, msg bool) {
	printTime(cT, msg)
	print("\033[2B\r")
	fmt.Printf("\n\tTimer Complete!\n\n")
}

func permuteArgs(args []string) {
	//rearrange the args array so that named arguments come first, this allows us to use positionals and named args
	args = args[1:] //remember changing the slice changes the underlying array
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

func main() {

	//param vars
	var timeToCount float64
	var direction bool
	var stopwatch bool
	var help bool
	var hours int
	var minutes int
	var seconds int

	flag.BoolVar(&direction, "u", false, "count up or down, default is down")
	flag.BoolVar(&stopwatch, "S", false, "stopwatch mode, count up until quit is given")
	flag.BoolVar(&help, "h", false, "display usage")
	flag.IntVar(&seconds, "s", 0, "time in seconds")
	flag.IntVar(&minutes, "m", 0, "time in minutes")
	flag.IntVar(&hours, "hr", 0, "time in hours")

	if len(os.Args) > 1 { //if we have args lets parse them for positional
		//by default we'll take whatever was the first position arg to be our time in seconds
		timeToCount, _ = strconv.ParseFloat(os.Args[1], 32)
		if timeToCount != 0 { //we had a positional so lets permute args so we can still use the flags lib
			permuteArgs(os.Args)
		}
		//check our args
		flag.Parse()
	} else { //if no args we'll run in stopwatch mode by default
		stopwatch = true
	}

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
	var msg bool

	printBanner() //banner sits over clock

	if stopwatch { // if stopwatch mode run until told to stop
		cT = 0.0
		msg = true
		for true {
			printTime(cT, msg)
			time.Sleep(incrementT)
			cT += incrementF
		}
	} else {
		if direction { //counting up not down
			cT = 0.0
			msg = true
			for cT < timeToCount {
				printTime(cT, msg)
				time.Sleep(incrementT)
				cT += incrementF
			}
		} else { //else counting down
			cT = timeToCount
			msg = false
			for cT > 0.0 {
				printTime(cT, msg)
				time.Sleep(incrementT)
				cT -= incrementF
			}
		}
	}

	//print goodbye message
	printGoodbye(cT, msg)
	return
}
