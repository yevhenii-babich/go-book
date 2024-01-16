package main

func callPanic() {
	// Uncomment this code to see the difference between panic and recovered panic.
	//defer func() {
	//	if err := recover(); err != nil {
	//		panicMesssage, ok := err.(string)
	//		if !ok {
	//			panicMesssage = "unknown panic"
	//		}
	//		println("recovered from callPanic():", panicMesssage)
	//	}
	//}()
	println("inside callPanic()")
	defer func() {
		println("useless code in the defer of callPanic()")
	}()
	panic("panic called in callPanic()")
	println("useless code") //this code will not be executed (never)
}
func recoverForDangerousCode() {
	defer func() {
		if err := recover(); err != nil {
			panicMessage, ok := err.(string)
			if !ok {
				panicMessage = "unknown panic"
			}
			println("recovered from recoverForDangerousCode():", panicMessage)
		}
	}()
	callPanic()
}
func recoveredPanic() {
	//defer func() {
	//	if err := recover(); err != nil {
	//		panicMessage, ok := err.(string)
	//		if !ok {
	//			panicMessage = "unknown panic"
	//		}
	//		println("recovered:", panicMessage)
	//	}
	//}()
	println("before callPanic()")
	recoverForDangerousCode() //replace with recoverForDangerousCode() to see the difference
	println("recoveredPanic code after callPanic()")
}

func main() {
	recoveredPanic()
	//callPanic()
	println("continued after panic in the main function")
}
