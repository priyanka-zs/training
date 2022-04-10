package main

import "fmt"

type ReadWriter interface {
	Read() string
	Write(string)
}
type Person struct {
	name string
}

func (p Person) Read() string {
	return p.name
}
func (p *Person) Write(s string) {
	fmt.Println(s)
}
func main() {
	p := &Person{"ram"}
	var rw ReadWriter = p
	var rw1 ReadWriter = &Person{"bal"}
	fmt.Println(rw.Read())
	rw.Write("abc")
	fmt.Println(rw1.Read())
	rw.Write("bcd")
}
