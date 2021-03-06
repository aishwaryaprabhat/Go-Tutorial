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

### Testing with Go
![](images/Screenshot%202020-08-14%20at%2011.25.37%20AM.png)
![](images/Screenshot%202020-08-14%20at%2011.28.22%20AM.png)

```
package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// Check if length of new deck is correct
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	// Check the last card is 4 of clubs
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()

	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

	// loadedDeck
}


```

## Organizing Data with Structs
![](images/Screenshot%202020-08-14%20at%2012.00.45%20PM.png)
![](images/Screenshot%202020-08-14%20at%2012.01.09%20PM.png)

### Create new struct
```
package main

type person struct {
	firstName string
	lastName  string
}

func main() {
	
}

```
### Create a new 'instance' of struct (nasty)
- Nasty
```
package main

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{"Alex", "Anderson"}
}

```
- Okay
```

package main

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
}
```
- Empty struct - Go assigns 0 values
![](images/Screenshot%202020-08-14%20at%2012.07.47%20PM.png)
```
package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var alex person
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
}

$ {firstName: lastName:}
```

### Assigning Values in Structs
```
package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var alex person
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Println(alex)
	fmt.Printf("%+v", alex)
	fmt.Println("")
}

```

### Structception: Embedding one struct into another
![](images/Screenshot%202020-08-14%20at%2012.17.12%20PM.png)

```
package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode string
}

func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: "120514",
		},
	}

	fmt.Printf("%+v", jim)

}

```
- Also valid

```
package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode string
}

func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: "120514",
		},
	}

	fmt.Printf("%+v", jim)

}


```


### Structs with Receiver Functions
```
package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode string
}

func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: "120514",
		},
	}

	
	jim.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}
```
## Pointers
### Pointers in Go (Why p.updateName doesn't work)
```
package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode string
}

func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: "120514",
		},
	}
	jim.updateName("Aish")
	jim.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

```
![](images/Screenshot%202020-08-14%20at%202.00.16%20PM.png)
![](images/Screenshot%202020-08-14%20at%202.01.14%20PM.png)
- Go does pass by value
- How to do it correctly
```
package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode string
}

func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: "120514",
		},
	}

	jimPointer := &jim
	jimPointer.updateName("Aish")
	jim.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

```

### Pointer Operations
![](images/Screenshot%202020-08-14%20at%202.06.57%20PM.png)
![](images/Screenshot%202020-08-14%20at%202.08.44%20PM.png)
![](images/Screenshot%202020-08-14%20at%202.12.26%20PM.png)

### Pointer Shortcuts
```
package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode string
}

func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: "120514",
		},
	}

	
	jim.updateName("Aish")
	jim.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

```
![](images/Screenshot%202020-08-14%20at%202.16.23%20PM.png)

### Gotchas With Pointers
![](images/Screenshot%202020-08-14%20at%202.22.25%20PM.png)
- Seems like there is no need for pointers to do a 'deep copy' sort of thing
- Arrays vs slices
![](images/Screenshot%202020-08-14%20at%202.23.28%20PM.png)
![](images/Screenshot%202020-08-14%20at%202.24.03%20PM.png)
![](images/Screenshot%202020-08-14%20at%202.25.17%20PM.png)
- Reference Types do not require pointers for deep copy sort of functionality, you can alter them directly
![](images/Screenshot%202020-08-14%20at%202.25.55%20PM.png)

## Maps
![](images/Screenshot%202020-08-14%20at%203.19.08%20PM.png)

### Declaration ways
- Simple
```
package main


import "fmt"

func main() {
	colors := map[string]string{
		"red": "12039i10239i01923",
		"green": "12039uijndkcjn",

	}
}
```
- 0 value initialization
```
package main

import "fmt"

func main() {

	var colors map[string]string

	// colors := map[string]string{
	// 	"red":   "12039i10239i01923",
	// 	"green": "12039uijndkcjn",
	// }

	fmt.Printf("%+v", colors)
}

```
- Another 0 init
```
package main

import "fmt"

func main() {

	colors := make(map[string]string)
	// var colors map[string]string

	// colors := map[string]string{
	// 	"red":   "12039i10239i01923",
	// 	"green": "12039uijndkcjn",
	// }

	fmt.Printf("%+v", colors)
}

```
```
package main

import "fmt"

func main() {

	colors := make(map[string]string)
	// var colors map[string]string

	// colors := map[string]string{
	// 	"red":   "12039i10239i01923",
	// 	"green": "12039uijndkcjn",
	// }

	colors["white"] = ";wdkjnaksjdnfajsndf;ijns"

	fmt.Printf("%+v", colors)
}

```
- Maps dont have . syntax like structs
- Delete 

### Iterating over a map
```
package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "wkdna;djnf;kasdjn",
		"white": "asijdnpaijsdnfaijsdnf",
		"green": "aksdjnc;kajsnd;kajn",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}

```

### Maps vs Structs
![](images/Screenshot%202020-08-14%20at%203.28.36%20PM.png)

## Interfaces
![](images/Screenshot%202020-08-20%20at%201.45.39%20PM.png)
- Interfaces make it easier to re-use code for different types without having to re-type

![](images/Screenshot%202020-08-20%20at%201.52.18%20PM.png)

- Bad way of writing since need to repeat the printGreeting function
```
package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

func (spanishBot) getGreeting() string {
	return "Hola Amigo!"
}

```

-  Refactored to use interfaces 

```
package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}
type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (spanishBot) getGreeting() string {
	return "Hola Amigo!"
}

```
- Explanation
![](images/Screenshot%202020-08-20%20at%202.07.47%20PM.png)
- In other words interfaces allow englishBot and spanishBot to have another type called bot, as long as they have a getGteeting function that returns string

### Rules for Interfaces
- We can define both the input and output type of the functions we want to be interfaced
![](images/Screenshot%202020-08-20%20at%202.57.39%20PM.png)
- We can define multiple functions in the interface
```
type bot interface {
	getGreeting() string
	getVersion() string
	respondToUser(user) string
}
```
- We cannot createa an 'object'/value from an interface type. We can only do that using concrete types
![](images/Screenshot%202020-08-20%20at%203.01.10%20PM.png)
- ![](images/Screenshot%202020-08-20%20at%203.02.35%20PM.png)

### Simple project to demo interfaces
```
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(resp)
}

```
![](images/Screenshot%202020-08-20%20at%203.14.04%20PM.png)
-  We can 'nest'/embed interfaces within one another

### Using Read Function to read the google home page html

```
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	bs := make([]byte, 999999)
	resp.Body.Read(bs)

	fmt.Println(string(bs))
}

```
- Doing the above in one line using io.Copy
```
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, resp.Body)
}

```
![](images/Screenshot%202020-08-21%20at%2010.05.30%20AM.png)
![](images/Screenshot%202020-08-21%20at%2010.06.11%20AM.png)


## Channels & Go Routines

### Project
![](images/Screenshot%202020-08-21%20at%2011.10.24%20AM.png)

- Without routines and channels
```
package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
}

```
- What is happening
![](images/Screenshot%202020-08-21%20at%2011.16.53%20AM.png)

- What we want instead is 
![](images/Screenshot%202020-08-21%20at%2011.17.56%20AM.png)

- How we are going to do it using go routines
![](images/Screenshot%202020-08-21%20at%2011.22.31%20AM.png)
- If a call is blocking, control is passed back to main routine so that some other go routine can be spun up and do its thing
- ![](images/Screenshot%202020-08-21%20at%2011.24.11%20AM.png)
- ![](images/Screenshot%202020-08-21%20at%2011.24.47%20AM.png)

- Go routines on one CPU 
![](images/Screenshot%202020-08-21%20at%2011.26.16%20AM.png)

- Go routines on multiple cores
![](images/Screenshot%202020-08-21%20at%2011.28.05%20AM.png)

- Concurrency is not parallelism
![](images/Screenshot%202020-08-21%20at%2011.30.14%20AM.png)
![](images/Screenshot%202020-08-21%20at%2011.31.27%20AM.png)

- Default Go routine (bug!)
![](images/Screenshot%202020-08-21%20at%2011.32.30%20AM.png)

```
package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		go checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
}

```
- This doesn't work because the main go routine ends before the child routines complete their work
- We need to use channels - way to communicate between go routines

### Channels
- Allows communication between go routines
![](images/Screenshot%202020-08-21%20at%2011.36.15%20AM.png)
- Channels are typed so all messages sent to a channel must be of the same type
![](images/Screenshot%202020-08-21%20at%2011.36.45%20AM.png)

#### How to send data over channels
![](images/Screenshot%202020-08-21%20at%2011.39.33%20AM.png)

- Receiving messages from a channel is a blocking call
```
package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string) //create a blank channel

	for _, link := range links {
		go checkLink(link, c)
	}

	fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think"
		return
	}

	fmt.Println(link, "is up!")
	c <- "Yup its up"
}

```

Output:
```
http://google.com is up!
Yup its up
```
- The reason why this is buggy and stops after Google is because
![](images/Screenshot%202020-08-21%20at%2011.50.51%20AM.png)

### Receiving messages over channels
```
package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string) //create a blank channel

	for _, link := range links {
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think"
		return
	}

	fmt.Println(link, "is up!")
	c <- "Yup its up"
}

```

- Keep checking
```
package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string) //create a blank channel

	for _, link := range links {
		go checkLink(link, c)
	}

	for {
		go checkLink(<-c, c)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}

```

### Alternate for loop
- Infinite for loop
```
for {
		go checkLink(<-c, c)
	}
```
same as 
```
for l := range c{
		go checkLink(l, c)
	}
```
- for i in range
```
for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
```

### Function Literal (similar to lambda function in Python)
```
for l := range c {

		go func() {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}()
	}
```
- This will not work because l is defined outside the scope of the functional literal 
- Correct way of doing it
```
for l := range c {

		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(l)
	}
```
- We never try to share variables between go routines (especially between main and child routines)
- We always either pass them as variables or we use channels