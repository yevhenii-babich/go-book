package main

func main() {
	x := 7 // declare and assign x to 7
	x = 5  // x is now 5
	println("x = 5:", x)
	x = 7  // x is now 7
	x += 5 // x is now 12 (equivalent to x = x + 5)
	println("x += 5:", x)
	x = 7
	x -= 5 // x is now 2 (equivalent to x = x - 5)
	println("x -= 5:", x)
	x = 7
	x *= 5 // x is now 35 (equivalent to x = x * 5)
	println("x *= 5:", x)
	x = 7
	x /= 5 // x is now 1 (equivalent to x = x / 5)
	println("x /= 5:", x)
	x = 7
	x %= 5 // x is now 2 (equivalent to x = x % 5)
	println("x %= 5:", x)
	x = 7
	x <<= 5 // x is now 224 (equivalent to x = x << 5)
	println("x <<= 5:", x)
	x = 7
	x >>= 5 // x is now 0 (equivalent to x = x >> 5)
	println("x >>= 5:", x)
	x = 7
	x &= 5 // x is now 5 (equivalent to x = x & 5)
	println("x &= 5:", x)
	x = 7
	x |= 5 // x is now 7 (equivalent to x = x | 5)
	println("x |= 5:", x)
	x = 7
	x ^= 5 // x is now 2 (equivalent to x = x ^ 5)
	println("x ^= 5:", x)
	x = 7
	x &^= 5 // x is now 2 (equivalent to x = x &^ 5)
	println("x &^= 5:", x)
	x = 7
	x++ // x is now 8 (equivalent to x = x + 1)
	println("x++:", x)
	x = 7
	x-- // x is now 6 (equivalent to x = x - 1)
	println("x--:", x)
}
