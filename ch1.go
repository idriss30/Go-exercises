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
	reverseUserNumber()
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

func reverseUserNumber(){
  fmt.Println("please enter a number");
  userInput := string("");
  fmt.Scan(&userInput);
  reversedString := string("")
  for i:= int(len(userInput)) -1; i >= 0 ; i-- {
    reversedString += string(userInput[i])
  }
  fmt.Printf("the reverse string is %v", reversedString)
}