/**
 * @author Nicholas Sun
 * @version 1.0.0
 * @date 2025-10-01
 * @fileoverview This program converts a dog's weight of 34.5 kg to pounds.
 */

package main

import "fmt"

func main() {
	// Calculates using the kilogams to pounds formula.
	fmt.Println("If your friend's dog weighs 34.5 kg, how many pounds is that?")
	fmt.Println("34.5 kg = " + fmt.Sprint(34.5*2.20462) + " pounds")

	fmt.Println("\nDone.")
}
