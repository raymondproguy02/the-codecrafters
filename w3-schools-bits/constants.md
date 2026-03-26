**Go Constans**

If a variable should have a fixed value that cannot be changed, you can use the const keyword.

The const keyword declares the variable as "constant", which means that it is unchangeable and read-only.

Syntax:
const CONSTNAME type = value


**Declaring a Constant**

Here is an example of declaring a constant in Go:

```
go
package main
import ("fmt")

const PI = 3.14

func main() {
  fmt.Println(PI)
}
```
**Constant Rules**

1. Constant names follow the same naming rules as variables
2. Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
3. Constants can be declared both inside and outside of a function

**Constant Types**
There are two types of constants:

1. Typed constants
2. Untyped constants