**Go For Loops**
The for loop loops through a block of code a specified number of times.
The for loop is the only loop available in Go.
The for loop can take up to three statements:

```go
for statement1; statement2; statement3 {
    // code to be executed for each iteration
}
```
Example:
```go
package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    fmt.Println(i)
  }
}
```

**Continue Statement**
The continue statement is used to skip one or more iterations in the loop. It then continues with the next iteration in the loop.
Example
```go
package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    if i == 3 {
      continue
    }
   fmt.Println(i)
  }
}
```

**Break Statement**
The break statement is used to break/terminate the loop execution.
Example:
```go
package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    if i == 3 {
      break
    }
   fmt.Println(i)
  }
}
```

**Nested Loops**
It is possible to place a loop inside another loop.
Example:
```go
package main
import ("fmt")

func main() {
  adj := [2]string{"big", "tasty"}
  fruits := [3]string{"apple", "orange", "banana"}
  for i:=0; i < len(adj); i++ {
    for j:=0; j < len(fruits); j++ {
      fmt.Println(adj[i],fruits[j])
    }
  }
}
```

**Range Keyword**
The range keyword is used to more easily iterate through the elements of an array, slice or map. It returns both the index and the value.

Example:
```go
package main
import ("fmt")

func main() {
  fruits := [3]string{"apple", "orange", "banana"}
  for idx, val := range fruits {
     fmt.Printf("%v\t%v\n", idx, val)
  }
}
```