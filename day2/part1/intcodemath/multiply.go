package intcodemath

func Multiply(intCode []int, currPosition int) {
	multiplicand := intCode[intCode[currPosition + 1]]
	multiplier := intCode[intCode[currPosition + 2]]
	product := multiplicand * multiplier
	savePosition := intCode[currPosition + 3]
	intCode[savePosition] = product

}
