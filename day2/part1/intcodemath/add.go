package intcodemath

func Add(intCode []int, currPosition int) {
	augend := intCode[intCode[currPosition + 1]]
	addend := intCode[intCode[currPosition + 2]]
	sum := augend + addend
	savePosition := intCode[currPosition + 3]
	intCode[savePosition] = sum
}
