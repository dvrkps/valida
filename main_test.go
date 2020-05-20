package valida

import "fmt"

func ExampleJMBG() {
	ok := JMBG("0308964384007")
	fmt.Println(ok)
	// Output:
	// true
}

func ExampleMBS() {
	okShort := MBS("01130234")
	fmt.Println(okShort)

	okLong := MBS("011302340123")
	fmt.Println(okLong)

	// Output:
	// true
	// true
}

func ExampleMID() {
	ok := MID("1333")
	fmt.Println(ok)
	// Output:
	// true
}

func ExampleOIB() {
	ok := OIB("69435151530")
	fmt.Println(ok)
	// Output:
	// true
}
