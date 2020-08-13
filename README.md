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
### 4. What 'func'?
### 5. How is the main.go organized?
   

