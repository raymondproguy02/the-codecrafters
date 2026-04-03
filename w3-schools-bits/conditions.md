**Go Conditions**
Conditional statements are used to perform different actions based on different conditions.

A condition can be either true or false.

Go supports the usual comparison operators from mathematics:

    Less than <
    Less than or equal <=
    Greater than >
    Greater than or equal >=
    Equal to ==
    Not equal to !=

Additionally, Go supports the usual logical operators:

    Logical AND &&
    Logical OR ||
    Logical NOT !

You can use these operators or their combinations to create conditions for different decisions.

Go has the following conditional statements:

4. Use switch to specify many alternative blocks of code to be executed

**If Statement**
Use if to specify a block of code to be executed, if a specified condition is true

```go
package main
import ("fmt")

func main() {
  if 20 > 18 {
    fmt.Println("20 is greater than 18")
  }
}
```

**Else Statement**
Use else to specify a block of code to be executed, if the same condition is false

```go
package main
import ("fmt")

func main() {
  time := 20
  if (time < 18) {
    fmt.Println("Good day.")
  } else {
    fmt.Println("Good evening.")
  }
}
```

**Else If Statement**
Use else if to specify a new condition to test, if the first condition is false

```go
package main
import ("fmt")

func main() {
  time := 22
  if time < 10 {
    fmt.Println("Good morning.")
  } else if time < 20 {
    fmt.Println("Good day.")
  } else {
    fmt.Println("Good evening.")
  }
}
```

**Nested If Statement**
You can have if statements inside if statements, this is called a nested if.

```go
package main
import ("fmt")

func main() {
  num := 20
  if num >= 10 {
    fmt.Println("Num is more than 10.")
    if num > 15 {
      fmt.Println("Num is also more than 15.")
     }
  } else {
    fmt.Println("Num is less than 10.")
  }
}
```

**Go Switch Statement**
Use the switch statement to select one of many code blocks to be executed.

```go
package main
import ("fmt")

func main() {
  day := 4

  switch day {
  case 1:
    fmt.Println("Monday")
  case 2:
    fmt.Println("Tuesday")
  case 3:
    fmt.Println("Wednesday")
  case 4:
    fmt.Println("Thursday")
  case 5:
    fmt.Println("Friday")
  case 6:
    fmt.Println("Saturday")
  case 7:
    fmt.Println("Sunday")
  }
}
```

**Multi-case Switch Statement**
It is possible to have multiple values for each case in the switch statement.

```go
package main
import ("fmt")

func main() {
   day := 5

   switch day {
   case 1,3,5:
    fmt.Println("Odd weekday")
   case 2,4:
     fmt.Println("Even weekday")
   case 6,7:
    fmt.Println("Weekend")
  default:
    fmt.Println("Invalid day of day number")
  }
}
```