package strings

func internalFuction() {
	// In go, a name is exported (visibile outside of the package) if it begins with a capital letter
}

// Must begin with captial letter in order to be exprted

func CountOddEven(s string) (odds, evens int) {
	odds, evens = 0, 0
	for _, c := range s {
		if int(c)%2 == 0 {
			evens++
		} else {
			odds++
		}
	}
	return
}
