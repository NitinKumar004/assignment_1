package main

import (
	"fmt"
	"strings"
	"time"
)

//currency converter

func validator(code string) bool {
	code = strings.ToUpper(code)
	return code == "USD" || code == "INR" || code == "EUR"
}
func converter(amount int, source, target string, mp map[string]float64) float64 {
	key := source + target
	if rate, found := mp[key]; found {
		converted := float64(amount) * rate
		return converted

	} else {
		fmt.Println("Conversion source or target not found")
		return -1
	}

}
func greet(name string) {
	currenthour := time.Now().Hour()
	if currenthour < 12 {
		fmt.Println("Good Morning ", name)
	} else if currenthour < 17 {
		fmt.Println("Good Afternoon", name)
	} else {
		fmt.Println("Good Evening", name)

	}

}
func main() {
	mp := map[string]float64{
		"USDINR": 83.12,
		"USDEUR": 0.93,
		"USDJPY": 156.82,
		"INRUSD": 0.012,
		"EURUSD": 1.07,
		"JPYUSD": 0.064,
	}
	var name string
	fmt.Println("Enter Your Name")
	fmt.Scan(&name)
	greet(name)

	fmt.Println("Current Rate of Currency Converter ")
	for key, value := range mp {
		from := key[:3]
		to := key[3:]
		fmt.Println("From:", from, "To:", to, "Value:", value)

	}

	var amount int
	var source string
	var target string
	fmt.Println("Enter the amount and followed by source and target")
	fmt.Scan(&amount)
	fmt.Println("Enter the source like USD INR ")
	fmt.Scan(&source)
	fmt.Println("Enter the target")
	fmt.Scan(&target)

	if validator(source) && validator(target) {
		source = strings.ToUpper(source)
		target = strings.ToUpper(target)
		floatvalue := fmt.Sprintf("%.2f", converter(amount, source, target, mp))
		fmt.Println("Amount after converting is : %.2f ", floatvalue)
	} else {
		fmt.Println("Sorry, the amount is not convertable : ", amount)
	}

}
