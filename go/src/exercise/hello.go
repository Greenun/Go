package main

func main(){
	const (
		test1 = iota
		test2
		test3
	)
	var testVar int = 140
	testVar2 := 160

	literals := `what is it..dododo..\n[]'2abcvdtttttttttttttttt'''`
	nonLiterals := "shit\ngood?"
	println(test1, test2, test3)
	println(testVar, testVar2)
	println(literals)
	println(nonLiterals)

	testVarU := uint(testVar)
	println(testVarU)
	b := []byte(nonLiterals)
	b2 := []byte(literals)
	println(b)
	println(b2)
	testVar++
	testVar--
	xorTest := (testVar ^ testVar2) << 10
	x := &testVar
	println(xorTest)
	println(x)
	println(*x)
	*x++
	if *x == 140{
		println("x == 140!")
	} else if *x == 141 {
		println("x == 141!")
	}
	if tel := testVar2; tel == 160 {
		println(tel)
	}

	switch testVar {
	case 140, 141:
		println("141, 142")
	default:
		println("Nope")
	}

	switch {
	case testVar >= 140:
		println("wow..?")
		fallthrough
	default:
		println("great")
	}

	switch interface{}(literals).(type) {
	case string:
		println("String")
	}

	L1:
		for i := 0; i < 10; i++ {
			println(i)
			if i == 5 {
				break L1
			}
		}
		for i := 0; i < 5; i++ {
			println("next loop", i)
		}
	println("test")
}



