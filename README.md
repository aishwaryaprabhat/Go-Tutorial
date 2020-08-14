# Go Tutorial
Getting Started with GoLang. Based on [this Udemy course.](https://www.udemy.com/course/go-the-complete-developers-guide)

## Overview

```
package main

import "fmt"

func main(){
	fmt.Println("Hi there!")
}
```

### 1. How do we run the code in our project?
![](images/Screenshot%202020-08-13%20at%202.49.37%20PM.png)

### 2. What does 'package main' mean?

#### What is a package?
1. Package is a collection of common source code files
2. Its like a project or workspace
3. One app typically has one package
4. A apackage can have many files in it
![](images/Screenshot%202020-08-13%20at%203.28.10%20PM.png)
5. Every file in the package must say which package it belongs to `package main`

#### What is main?
There are two types of packages:
1. Executable
   1. Generates a file that we can run
   2. Used for actually doing something (eg: run a server)
   3. An executable package is called main....package main is sacred
   4. Package main must always have a main function
2. Reusable
   1. Code used as helpers
   2. Good place to put reusable logic

### 3. What does 'import fmt mean'?
1. Importing libraries to do stuff main can't do
2. [Golang standard packages](http://golang.org/pkg)

### 4. What 'func'?
1. Functions just like functions in any other programming language

### 5. How is the main.go organized?
![](images/Screenshot%202020-08-13%20at%203.41.33%20PM.png)
   

## Project Overview
A package that can do the following:
![](images/Screenshot%202020-08-13%20at%203.46.27%20PM.png)

### Variable Declaration
- Go is statically typed vs dynamically typed eg Python, JS
- We need to define a type for a variable
```
package main

import "fmt"

func main(){
	var card string = "Ace of spades"
	fmt.Println(card)
}
```
- Basic data types
![](images/Screenshot%202020-08-13%20at%203.52.13%20PM.png)
- Letting compiler figure out the type

```
package main

import "fmt"

func main(){
	// var card string = "Ace of spades"
	var card:= "Ace of spades"
	fmt.Println(card)
}
```
- For reassigning value we don't have to use a colon

### Functions and Return Types
- Need to define 
```
package main

import "fmt"

func main(){
	var card = newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
```

### Slices (re-sizable Array)
- Needs a type defined as well
```
package main

import "fmt"

func main(){
	cards := []string{newCard(), newCard(), "Ace of Diamonds"}
}

func newCard() string {
	return "Five of Diamonds"
}
```
- Appending: returns a new slice
```
package main

import "fmt"

func main(){
	cards := []string{newCard(), newCard(), "Ace of Diamonds"}
	
	cards = append(cards, "Six of spades")
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
```

### Looping (over slice)
```
package main

import "fmt"

func main(){
	cards := []string{newCard(), newCard(), "Ace of Diamonds"}
	
	cards = append(cards, "Six of spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
	
}

func newCard() string {
	return "Five of Diamonds"
}
```
![](images/Screenshot%202020-08-13%20at%205.17.15%20PM.png)

### OO approach vs Go approach
![](images/Screenshot%202020-08-13%20at%205.20.03%20PM.png)
![](images/Screenshot%202020-08-13%20at%205.20.12%20PM.png)

### Defining a new type
```
package main

//Create a new type of 'deck'
//which is a slice of strings

type deck []string 

```
```
package main

import "fmt"

func main(){
	cards := deck{newCard(), newCard(), "Ace of Diamonds"}
	
	cards = append(cards, "Six of spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
	
}

func newCard() string {
	return "Five of Diamonds"
}
```
### Receiver Functions - 'Extending' functionality of new type

```
package main
import "fmt"

//Create a new type of 'deck'
//which is a slice of strings

type deck []string 

func (d deck) print(){
	for i, card := range d{
		fmt.Println(i, card)
	}
}
```
```
package main

// import "fmt"

func main(){
	cards := deck{newCard(), newCard(), "Ace of Diamonds"}
	
	cards = append(cards, "Six of spades")

	cards.print()
	
}

func newCard() string {
	return "Five of Diamonds"
}
```
![](images/Screenshot%202020-08-13%20at%205.29.46%20PM.png)
![](images/Screenshot%202020-08-13%20at%205.30.11%20PM.png)

### Double loop 
- We use "_" to let go know that we created a variable but we don't intend to use it
```
package main
import "fmt"

//Create a new type of 'deck'
//which is a slice of strings

type deck []string 

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits{
		for _, value := range cardValues{
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print(){
	for _, card := range d{
		fmt.Println(card)
	}
}
```

### Slice Range Syntax + Multiple Returns
![](images/Screenshot%202020-08-13%20at%205.46.45%20PM.png)

```
//return two different slices
func deal(d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}
```
```
func main(){
	
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)
	hand.print()
	remainingDeck.print()
	
}
```

### Type conversion (deck to string)
![](images/Screenshot%202020-08-13%20at%206.11.02%20PM.png)
- deck->string->byte slice

![](images/Screenshot%202020-08-13%20at%206.12.03%20PM.png)

```
func (d deck) toString() string{
	// []string(d) //convert to slice of string
	return strings.Join([]string(d), ",")
}
```

- byte slice->string->deck
```
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option 1 - log the error and default to newDeck()
		// Option 2 - log the error and quit program
		fmt.Println("Error:", err)
		os.Exit(1)

	}
	s := strings.Split(string(bs), ",")

	return deck(s)
}
```

## Shuffling cards (random number generation, swapping in a slice)
- Pseudo random
```
func (d deck) shuffle() {

	for i := range d {
		newPosition := rand.Intn(len(d) - 1)

		d[newPosition], d[i] = d[i], d[newPosition]
	}

}
```

## Generating random numbers
- Use time -> use it as seed value -> use new Rand
```
func generateRandomNum(l int) int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return r.Intn(l)
}

```