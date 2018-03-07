/*
Go / Golang

- Strong-typed
- Compiled to machine code
- Not OO
- Struct
- Interface
- Composition
*/

var i int

i = "fooo" // compile err

var s string




// Go uses minimal punctuation
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// Go assignment
var foo string //without type derivation
foo = "bar"

foo := "bar" //with type derivation


bar = "blah"


// Go export
package apackagename //define package name


func Bar (s string) string {
    //Bar will be exported
}

func bar (s string) string {

}

bool 
int


// Go import
import (
    "fmt" // part of Go’s standard library
    "github.com/foo/foo" // imported directly from repository
)
foo.Bar()


// Go - return multiple values
func bar() (int, error) {
    return 1, nil
}
a, b := bar()


//Go error handling
foo, err := bar()
if err != nil {
    // handle error with defer, panic, recover, or log.fatal, etc...
}


// Variadic function
func foo (args ...int) {
	fmt.Println(len(args))
}

func main() {
	foo() // 0
	foo(1,2,3) // 3
}

