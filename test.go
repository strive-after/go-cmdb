package main

import "fmt"

func T(a interface{}) interface{}{
	user := a.([]string)
	user[0] = "1"
	return user
}

func main() {
	c := T([]string{"a","b"})
	fmt.Printf("%#v",c)
}
