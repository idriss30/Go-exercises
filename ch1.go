package main

import (
	"fmt"
)

func main() {
	// printZero()
	//returnInteger()
	
	

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

func displayAlphabet(){
	 alphabet := string("ABCDEFGHIJKLMNOPQRSTUVWXYZ");
	 for i := 0 ; i < len(alphabet); i++ {
		fmt.Printf("%v", alphabet[i])
	 }
}
