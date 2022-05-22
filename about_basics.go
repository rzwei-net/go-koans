package go_koans

func aboutBasics() {
	assert(true == true)  // what is truth?
	assert(true != false) // in it there is nothing false

	var i int = 1.0
	assert(i == 1.0000000000000000000000000000000000000) // precision is in the eye of the beholder

	k := 1.0 //short assignment can be used, as well
	assert(k == 1.0000000000000000000000000000000000000)

	assert(5%2 == 1) // remainder from 5/2 = 1
	assert(5*2 == 10)

	// caret means xor => different bit = 1, same bit = 0
	//     4 2 1
	// 5   1 0 1
	// 2   0 1 0
	//          xor
	// 7   1 1 1
	assert(5^2 == 7)

	var x int      // initial value = 0
	assert(x == 0) // zero values are valued in Go

	var f float32    // inital value = 0.0
	assert(f == 0.0) // for types of all types

	var s string    // initial value = ""
	assert(s == "") // both typical or atypical types

	var c struct {
		x int
		f float32
		s string
	}
	assert(c.x == 0)   // and types within composite types
	assert(c.f == 0.0) // which match the other types
	assert(c.s == "")  // in a typical way
}
