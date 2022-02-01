package main

import (
	"fmt"
	"strconv"
)

func greet(name, year string) {
	fmt.Println("Hello! My name is " + name + ".")
	fmt.Println("I was created in " + year + ".")
}

func showName() {
	var name string
	fmt.Println("Please, remind me your name.")
	fmt.Scan(&name)
	fmt.Println("What a great name you have, " + name + "!")
}

func guessAge() {
	var rem3, rem5, rem7, age int

	fmt.Println("Let me guess your age.")
	fmt.Println("Enter remainders of dividing your age by 3, 5 and 7.")
	fmt.Scan(&rem3, &rem5, &rem7)

	age = (rem3*70 + rem5*21 + rem7*15) % 105
	fmt.Println("Your age is " + strconv.Itoa(age) + "; that's a good time to start programming!")
}

func count() {
	var n int

	fmt.Println("Now I will prove to you that I can count to any number you want.")
	fmt.Scan(&n)
	for i := 0; i <= n; i++ {
		fmt.Printf("%d!\n", i)
	}
}

func startQuiz() {
	fmt.Println("Let's test your programming knowledge.")
	// write your code here
	fmt.Println("How many bits are there in a byte?")
	fmt.Println("1. 6 bits.")
	fmt.Println("2. 8 bits.")
	fmt.Println("3. 9 bits.")
	fmt.Println("4. any of the above.")
	var reply int
	gameloop := 1
	for gameloop > 0 {
		fmt.Scan(&reply)
		if reply == 4 {
			gameloop = 0
		} else {
			fmt.Println("Please, try again.")
		}
	}

	fmt.Println("Completed, have a nice day!")
}

func sayGoodbye() {
	fmt.Println("Congratulations, have a nice day!")
}

func main() {
	greet("Alice", "2022") // change it as you need
	showName()
	guessAge()
	count()
	// call something here
	startQuiz()
	sayGoodbye()
}
