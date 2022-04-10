package main

import "fmt"

type student struct {
	name string
	age  int
}

func (s student) show() {
	fmt.Println("name: ", s.name)
	fmt.Println("age: ", s.age)
}

/*func main() {
	/*n := student{"hari", 20}
	s1 := student{name: "hani"}
	fmt.Println(n, s1)
	st := student{"raju", 20}
	st.show()
}*/
