package go_koans

func aboutPointers() {
	{
		a := 3
		b := a // 'b' is a copy of 'a' (all assignments are copy-operations)

		b++

		assert(a == 3) // variables are independent of one another
	}

	{
		a := 3

		// 'b' is the address of 'a'
		// '&' returns the memory address of 'a'
		// in other words, this is _not_ a copy operation, and we are able to mutate
		// 'a' and 'b' are references to the same memory address
		b := &a

		// de-referencing 'b' means acting like a mutable copy of 'a'
		// since we assigned 'b' to the memory address of 'a'
		// if we want to change the value
		// we must de-reference 'b' to return an int
		*b = *b + 2

		/*
			println(&a) ... memory address -> 0xc00005fed8
			println(b)  ... also memory address -> 0xc00005fed8
			println(a)  ... un-references value -> 5
			println(*b) ... de-references value -> 5
			println(&b) ... still a memory address -> 0xc00005fed8
		*/

		assert(a == 5) // pointers seem complicated at first but are actually simple
	}

	{
		increment := func(i int) {
			i++
		}

		a := 3
		increment(a)
		assert(a == 3) // variables are always passed by value, and so a copy is made
	}

	{
		realIncrement := func(i *int) {
			*i++
		}

		b := 3
		realIncrement(&b)
		assert(b == 4) // but passing a pointer allows others to mutate the value pointed to
	}
}
