package main

import "fmt"

/* Initially on a notepad only one character 'A' is present. You can perform two operations on this notepad for each step:

    Copy All: You can copy all the characters present on the notepad (partial copy is not allowed).
    Paste: You can paste the characters which are copied last time.



Given a number n. You have to get exactly n 'A' on the notepad by performing the minimum number of steps permitted. Output the minimum number of steps to get n 'A'.

Example 1:

Input: 3
Output: 3
Explanation:
Intitally, we have one character 'A'.
In step 1, we use Copy All operation.
In step 2, we use Paste operation to get 'AA'.
In step 3, we use Paste operation to get 'AAA'. */

func minSteps(n int) int {
	if n == 1 {
		return 0 // edge case
	}

	// init d := 2
	d := 2 // divisor
	i := 0 // paste count
	j := 0 // copy count

	// divide by smallest divisor > 1 w/o remainder to get highest quotient
	// divide quotient by quotient's smallest divisor until quotient == 1
	for {
		// if n % d != 0 (remainder)
		if n%d != 0 {
			d++ // increment divisor if remainder
			continue
		} else {
			if n/d == 1 { // quotient == 1 (cannot divide further)
				i += n // paste count += dividend (1 * d)
				break
			}
			i += (d - 1) // increment paste count by d - 1
			n = n / d    // next dividend = highest quotient
			j++          // increment copy count
		}
	}
	return i + j // total steps == paste count + copy count
}

// test
func main() {
	n := 30
	i := minSteps(n)
	fmt.Println(i)
}

// Proofs
// 30 / 2 = 15
//   i += (2-1)
//   j += 1
// 15 / 3 = 5
//   i += (3-1)
//   j += 1
// 5 / 5 = 1
//   i += 5
//   j += 0 (no further copies required - sum == n)
// i == 1(15) + 2(5) + 5(1) = 8 pastes
// j == 2 copies (copy 5, copy 15)
// (1+1+1+1+1) + (5+5) + 15 = 30
// 5 + 10 + 15 = 30

// 12 / 2 = 6
//   i += 1
//   j++
// 6 / 2 = 3
//   i += 1
//   j++
// 3 / 3 = 1
//   i += 3
// i == 1(6) + 1(3) + 3(1) == 5 pastes
// j == 2 copies (copy 3, copy 6)

// 13 / 13 = 1
//   i += 13
//   j += 0 (prime number)
// i == 13(1)
// j = 0 copies (no copies needed for prime number)
