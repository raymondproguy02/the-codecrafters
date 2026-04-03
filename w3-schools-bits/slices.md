**Go Slices**

Slices are similar to arrays, but are more powerful and flexible.

Like arrays, slices are also used to store multiple values of the same type in a single variable.

However, unlike arrays, the length of a slice can grow and shrink as you see fit.

In Go, there are several ways to create a slice:

    Using the []datatype{values} format
    Create a slice from an array
    Using the make() function

Create a Slice With []datatype{values}
Syntax
slice_name := []datatype{values}

A common way of declaring a slice is like this:
myslice := []int{}

The code above declares an empty slice of 0 length and 0 capacity.

To initialize the slice during declaration, use this:
myslice := []int{1,2,3}

The code above declares a slice of integers of length 3 and also the capacity of 3.

In Go, there are two functions that can be used to return the length and capacity of a slice:

    len() function - returns the length of the slice (the number of elements in the slice)
    cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)

Example

This example shows how to create slices using the []datatype{values} format:

```go
package main
import ("fmt")

func main() {
  myslice1 := []int{}
  fmt.Println(len(myslice1))
  fmt.Println(cap(myslice1))
  fmt.Println(myslice1)

  myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
  fmt.Println(len(myslice2))
  fmt.Println(cap(myslice2))
  fmt.Println(myslice2)
}
```

Result:
0
0
[]
4
4
[Go Slices Are Powerful]