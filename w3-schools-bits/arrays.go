**Go Arrays**

Arrays are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value.
Declare an Array

In Go, there are two ways to declare an array:
1. With the var keyword:
Syntax
```go
var array_name = [length]datatype{values} // here length is defined

or

var array_name = [...]datatype{values} // here length is inferred
2. With the := sign:
Syntax
array_name := [length]datatype{values} // here length is defined

or

array_name := [...]datatype{values} // here length is inferred
```

Note: The length specifies the number of elements to store in the array. In Go, arrays have a fixed length. The length of the array is either defined by a number or is inferred (means that the compiler decides the length of the array, based on the number of values).
Array Examples
Example

This example declares two arrays (arr1 and arr2) with defined lengths:

```go
package main
import ("fmt")

func main() {
  var arr1 = [3]int{1,2,3}
  arr2 := [5]int{4,5,6,7,8}

  fmt.Println(arr1)
  fmt.Println(arr2)
}
```

Result:
[1 2 3]
[4 5 6 7 8] 