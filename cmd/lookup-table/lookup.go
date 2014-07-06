package main

import "fmt"

func main() {
	var alphabet []byte

	for c := 'a'; c <= 'z'; c++ {
		alphabet = append(alphabet, byte(c))
	}

	for c := 'A'; c <= 'Z'; c++ {
		alphabet = append(alphabet, byte(c))
	}

	for c := '0'; c <= '9'; c++ {
		alphabet = append(alphabet, byte(c))
	}

	fmt.Printf("const base = %d\n", len(alphabet))
	fmt.Printf("var alphabet = [base]byte{")
	for i, c := range alphabet {
		fmt.Printf("'%c'", c)
		if i != len(alphabet)-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("}\n")

	max := (1 << 8)
	hash := make([]int, max)

	for i := 1; i < max; i++ {
		for j := 1; j < len(alphabet); j++ {
			if i == int(alphabet[j]) {
				hash[i] = j
				break
			}
		}
	}

	fmt.Printf("\nvar alphabetLookup = [1 << 8]int{")

	for i, d := range hash {
		fmt.Printf("%d", d)
		if i != len(hash)-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("}\n")

}
