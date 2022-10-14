package main

import (
	"fmt"
	"strconv"
)

func main() {
	// printZero()
	//returnInteger()
	//displayAlphabet()
	//print0To100Sum()
	//loopDivisibleBy50()
	//rangeDivisibleBy50()
	//promptInteger()
	//reverseUserNumber()
	//length()
	//fizzBuzz()
	findIncomeTax()
}

// resulting in zero 4-1
func printZero() {
	//fmt.Println(5 + 3 % 2 * 9)
	// solution
	fmt.Println((5 + 3) % 2 * 9)
}

// exercice 4-2
func returnInteger() {
	var userFloatNumber float64
	fmt.Println("Please enter a float number")
	fmt.Scan(&userFloatNumber)
	integerPortion := uint(userFloatNumber)
	fmt.Printf("the integer portion is %v\n", integerPortion)
}

// 4-3 current value of deposit

func currentValue() {
	var initialDeposit int
	var interestRate float32
	var interestCalculator int
	var numberOfYearsSinceDeposit int
	fmt.Println("please enter the initial deposit value")
	fmt.Scan(&initialDeposit)
	fmt.Println("please enter the interest rate you wish to apply ")
	fmt.Scan(&interestRate)
	fmt.Println("please enter the number of times the interest is calculated a year ")
	fmt.Scan(&interestCalculator)
	fmt.Println("how many years has it been since the initial Deposit ?")
	fmt.Scan(&numberOfYearsSinceDeposit)
	value := initialDeposit*(1+int(interestRate)/interestCalculator) ^ interestCalculator*numberOfYearsSinceDeposit
	fmt.Printf("An initial deposit of %v, at an interest rate of %v calculated %v per year for %v year(s)\n will give you %v\n ", initialDeposit, interestRate, interestCalculator, numberOfYearsSinceDeposit, value)

}

//Using for loops and string functions, create a computer program that displays the alphabet from A to Z.

func displayAlphabet() {
	alphabet := string("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	for i := 0; i < len(alphabet); i++ {
		fmt.Printf("%v\n", string(alphabet[i])) // important to convert to string first.
	}
}

//Create a program that computes the sum of all numbers between 0 and 100

func print0To100Sum() {
	sum := 0
	for i := 1; i < 100; i++ { // range is ( 1 - 99)
		sum += i
	}
	fmt.Printf("the sum of all numbers between 0 and 100 is %v", sum)
}

//Write two programs, each of which displays all numbers divisible by 50 between 100 and 1,000 (inclusive).
//Both programs should have identical output.

//Use the range keyword with for in one program.
//Use for without range in the other program.

func loopDivisibleBy50() {
	for i := 100; i <= 1000; i++ {
		if i%50 == 0 {
			fmt.Printf("%v is divisible by 50\n", i)
		}
	}
}

func rangeDivisibleBy50() {
	for v := range [1001]int{} {
		if v < 100 {
			continue
		}
		if v%50 == 0 {
			fmt.Printf("%v, is divisible by 50\n", v)
		}
	}

}

//Create a program that prompts the user for an integer and then displays the following information about that number:
//The number of digits in the value entered
//The first and last digits of the number
//The sum of the digits in the number
//The product of the digits in the number
//Whether or not the number is prime
//The factorial of the number

func promptInteger() {
	fmt.Println("please enter a number")
	var userInput string
	fmt.Scan(&userInput)
	var userInputToNumber int
	userInputToNumber, _ = strconv.Atoi(userInput)

	fmt.Printf("the number of digits is %v\n", len(userInput))
	firstDigit := string(userInput[0])
	lastDigit := string(userInput[int(len(userInput))-1]) // get the value as string but the last item converting string to int
	fmt.Printf("the first digit is %v and the last one is %v\n", firstDigit, lastDigit)

	var stringToNumber int
	var sum = 0
	var digitMultiplication = 1

	for i := 0; i < len(userInput); i++ {
		stringToNumber, _ = strconv.Atoi(string(userInput[i]))
		sum += stringToNumber
		digitMultiplication *= stringToNumber
	}

	fmt.Printf("the sum of all the digits is %v\n", sum)
	fmt.Printf("the product of the digits is %v\n", digitMultiplication)

	if userInputToNumber%2 != 0 && sum/3 != 0 {
		fmt.Printf("%v is a prime number\n", userInput)
	} else {
		fmt.Printf("%v is not a prime number\n", userInput)

	}
	// factorial;
	var factorial int = 1
	for i := 1; i < userInputToNumber; i++ {
		j := i + 1

		if j > userInputToNumber {
			break
		}
		factorial *= j

	}

	fmt.Printf("%v factorial is %v", userInputToNumber, factorial)

}

//Write a program that asks the user for an input integer and then computes the reverse of that number

func reverseUserNumber() {
	fmt.Println("please enter a number")
	userInput := string("")
	fmt.Scan(&userInput)
	reversedString := string("")
	for i := int(len(userInput)) - 1; i >= 0; i-- {
		reversedString += string(userInput[i])
	}
	fmt.Printf("the reverse string is %v", reversedString)
}

//Write a program that computes the length of a string without using the len function.

func length() {
	fmt.Println("enter your string")
	stringInput := string("")
	fmt.Scan(&stringInput)
	count := 0
	for index := range stringInput {
		count = index
	}
	fmt.Printf("the length is %v", count)
}

/* Write a program that loops through a series of values and uses those values to determine
the output shown to the user. The program should perform the following steps:
Ask the user for a number.
Output a count starting with 0.
Display the count number if it is not divisible by 3 or 5.
Replace every multiple of 3 with the word “fizz.”
Replace every multiple of 5 with the word “buzz.”
Replace multiples of both 3 and 5 with “fizz buzz.”
Continue counting until the number of integers replaced with “fizz,” “buzz,” or “fizz buzz” reaches the input number.
The last output line should read “TRADITION!!” */

func fizzBuzz() {
	fmt.Println("please enter a number")
	initialNumber := string("")
	fmt.Scan(&initialNumber)
	var intInitialNumber int
	intInitialNumber, _ = strconv.Atoi(initialNumber)
	var initialCountofWords = 0
	for i := 0; i >= 0; i++ {
		if initialCountofWords == intInitialNumber {
			fmt.Println("TRADITION !!")
			break
		}

		if i == 0 {
			fmt.Println(i)
			continue

		} else {
			if i%3 == 0 && i%5 == 0 {
				fmt.Println("fizz buzz")
				initialCountofWords++
				continue
			}
			if i%5 == 0 {
				fmt.Println("Buzz")
				initialCountofWords++
				continue
			}
			if i%3 == 0 {
				fmt.Println("Fizz")
				initialCountofWords++
				continue
			}
		}

		fmt.Println(i)

	}

}

// create a program that calculate the income tax of a
// single filer based on the year 2020 information
// you are using $35,987.65 for the gross income value, with two dependents, except where stated otherwise
// When you test the code with input values, you can use different values to see what happens
// All taxpayers are allowed a $12,200 standard deduction.
// For each dependent, a taxpayer is allowed an additional $2,000 deduction
/* 10%	Up to $9,875
//12%	$9,876 to $40,125
22%	$40,126 to $85,525
24%	$85,526 to $163,300
32%	$163,301 to $207,350
35%	$207,351 to $518,400
37%	$518,401 or more */

// as a courtesy program can stop at any percentage deduction

func findIncomeTax() {
	//finding income tax stopping at 24% bracket 163300
	var grossIncome float64
	var dependents int
	var taxableIncome float64
	var taxDue float64
	var max10 = 9875
	var max12 = 40125
	var max22 = 85525

	fmt.Print("Enter your gross income from your 2020 w2:\n")
	fmt.Scan(&grossIncome)
	fmt.Printf("Your gross income is: $%v\n", grossIncome)
	fmt.Print("How many dependents are you claiming ?\n")
	fmt.Scan(&dependents)
	fmt.Printf("Your claimed number of dependents is : %v\n", dependents)

	// tax calculation
	taxableIncome = grossIncome - 12200 - (2000 * float64(dependents))
	fmt.Printf("Your taxable income is %v\n", taxableIncome)
	if taxableIncome <= 0 {
        taxDue = 0
	}else if taxableIncome <= float64(max10) {
		taxDue = taxableIncome * 0.1
	} else if taxableIncome <= float64(max12) {
		taxDue = taxableIncome*0.1 + (taxableIncome-float64(max10))*0.12
	} else if taxableIncome <= float64(max22) {
		taxDue = taxableIncome*0.1 + (taxableIncome-float64(max10))*0.12 + (taxableIncome-float64(max12))*0.22
	} else {
		taxDue = taxableIncome*0.1 + (taxableIncome-float64(max10))*0.12 + (taxableIncome-float64(max12))*0.22 + (taxableIncome-float64(max22))*0.24
	}
	fmt.Printf("your tax due is $%v\n", taxDue)
}
