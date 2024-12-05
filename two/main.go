package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
	1. Parse the report into a list of levels.

	2. If the report has fewer than two levels, return "safe."

	3. Determine direction from the first two levels:
	- If `level[1] > level[0]`, it's increasing.
	- If `level[1] < level[0]`, it's decreasing.
	- If `level[1] == level[0]`, decide based on the first non-equal pair.

	4. Iterate through the levels:
	- For each pair of adjacent levels (`i`, `i+1`):
		a. Check if the difference is within the range `[1-3]`. If not, return "unsafe."
		b. Check if the direction is consistent. If not, return "unsafe."

	5. If all checks pass, return "safe."
*/

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		numbers := strings.Split(scanner.Text(), " ")
		report := make([]int, len(numbers))
		for i, num := range numbers {
			val, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(err)
				return
			}
			report[i] = val
		}
		isIncreasing := report[1] > report[0]
		safe := true

		for i := 0; i < len(report)-1; i++ {
			diff := report[i+1] - report[i]

			if abs(diff) < 1 || abs(diff) > 3 {
				fmt.Printf("The level %v is unsafe due to invalid difference between two adjacent levels.\n", report)
				safe = false
				break
			}

			if isIncreasing && diff < 0 || !isIncreasing && diff > 0 {
				fmt.Printf("The level %v is unsafe due to direction inconsistency.\n", report)
				safe = false
				break
			}
		}

		if safe {
			fmt.Printf("The level %v is safe.\n", report)
			count += 1
		}
	}
	fmt.Println("Total count of safe level is:", count)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
