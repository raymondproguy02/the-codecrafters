**Go Functions**
A function is a block of statements that can be used repeatedly in a program.

A function will not execute automatically when a page loads.

A function will be executed by a call to the function.

Example:
```go
package main
import ("fmt")

func myMessage() {
  fmt.Println("I just got executed!")
}

func main() {
  myMessage() // call the function
}
```

**Parameters and Arguents**
Information can be passed to functions as a parameter. Parameters act as variables inside the function.

Parameters and their types are specified after the function name, inside the parentheses. You can add as many parameters as you want, just separate them with a comma.

Example:

```go
package main
import ("fmt")

func familyName(fname string) {
  fmt.Println("Hello", fname, "Refsnes")
}

func main() {
  familyName("Liam")
  familyName("Jenny")
  familyName("Anja")
}
```

**Return Values**
If you want the function to return a value, you need to define the data type of the return value (such as int, string, etc), and also use the return keyword inside the function.

Examples:
```go
package main
import ("fmt")

func myFunction(x int, y int) int {
  return x + y
}

func main() {
  fmt.Println(myFunction(1, 2))
}
```

**Multiple Return Values**
Go functions can also return values.

Examples
```go
package main
import ("fmt")

func myFunction(x int, y string) (result int, txt1 string) {
  result = x + x
  txt1 = y + " World!"
  return
}

func main() {
  fmt.Println(myFunction(5, "Hello"))
}
```

**Recursion Functions**
Go accepts recursion functions. A function is recursive if it calls itself and reaches a stop condition.

Example:
```go
package main
import ("fmt")

func testcount(x int) int {
  if x == 11 {
    return 0
  }
  fmt.Println(x)
  return testcount(x + 1)
}

func main(){
  testcount(1)
} 
```