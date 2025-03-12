package exercises

import "fmt"

func Exercise3() {
	// INFO: adding 1 to each variable causes an overflow, not an error/panic

	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	b += 1
	smallI += 1
	bigI += 1

	fmt.Println(b, smallI, bigI)
}
