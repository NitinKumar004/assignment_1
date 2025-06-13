package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func calculatepercent(count, total int) float64 {
	if total == 0 {
		return 0.0
	}
	return (float64(count) / float64(total)) * 100
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("enter please two arguments like:  go run  index.go and then the log file")
		return
	}

	filename := os.Args[1]         //fetching the first command line argument
	file, err := os.Open(filename) // this attempt to open the file for reading
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	//this will ensure file will get close when main() ends

	//create variable to count the data
	var infocount, warningcount, errcount, totallines int

	scanner := bufio.NewScanner(file)
	//use bufio to read file line by line

	//loop through each line by line
	for scanner.Scan() {
		// if text find increase the count
		line := scanner.Text()
		totallines++
		//use switch to count what type of log is
		switch {
		// if it contains INFO  IN  given log file
		case strings.Contains(line, "[INFO]"):
			infocount++
		// if it contains Warning
		case strings.Contains(line, "[WARNING]"):
			warningcount++
		case strings.Contains(line, "[ERROR]"):
			errcount++
		}
	}
	//if something error find in any file it will show some error in it and return from here
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	// print all data here
	fmt.Printf("\nLog Analysis of file: %s\n\n", filename)
	fmt.Printf("INFO: %d entries (%.2f%%)\n", infocount, calculatepercent(infocount, totallines))
	fmt.Printf("WARNING: %d entries (%.2f%%)\n", warningcount, calculatepercent(warningcount, totallines))
	fmt.Printf("ERROR: %d entries (%.2f%%)\n", errcount, calculatepercent(errcount, totallines))
	fmt.Printf("\nTotal log lines: %d\n", totallines)
	fmt.Printf("Analyzed at: %s\n", time.Now().Format("2006-01-02 15:04:05"))

}
