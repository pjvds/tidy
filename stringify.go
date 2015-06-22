package tidy

import "fmt"

// Stringify prints the value in a default format. When printing
// structs or pointers to structs it prints the field names.
func Stringify(object interface{}) string {
	return fmt.Sprintf("%+v", object)
}
