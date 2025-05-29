// my very first go program
package main

import (
	"fmt"

	"github.com/scaroph/go-learning-fiesta/file2"
)

// Exported variable
var ExportedVariable = "Hello, World!"

func main() {
	// Accessing exported identifier in the same file
	fmt.Println(ExportedVariable)

	// Accessing exported identifier from another package
	fmt.Println(file2.AnotherExportedVariable)
}

/*
// Valid identifiers:
_geeks23
geeks
gek23sd
Geeks
geeKs
geeks_geeks

// Invalid identifiers:
212geeks
if
default
*/
