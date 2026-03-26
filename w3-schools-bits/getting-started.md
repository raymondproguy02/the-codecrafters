**Getting started with Go**

To start writing go or work with it, you need two things:
1. A text editor, eg VS code
2. A compiler, like GCC, to translate the Go code into a language the computer understands

**Installing Go**

Check out the officail Go website at https://golang.org

Download then check version by go version

**Go Quickstart**
Enter VS code download go extention then download.

Create simple program in a file hello.go, then write this inside the file: 
```
go

package main
import ("fmt")

func main() {
  fmt.Println("Hello, World!")
} 
```
Now save control + s then control + j open terminal locate the file and type go run hello.go.

Output "Hello, World!"
Tadda! your first program.