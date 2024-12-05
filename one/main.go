package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var a []float64
	var b []float64

	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), "   ")
		if len(nums) == 2 {
			first, firsterr := strconv.ParseFloat(nums[0], 64)
			second, seconderr := strconv.ParseFloat(nums[1], 64)

			if firsterr == nil && seconderr == nil {
				a = append(a, first)
				b = append(b, second)
			}
		} else {
			log.Printf("Skipping invalid line: %s", scanner.Text())
		}
	}

	sort.Float64s(a)
	sort.Float64s(b)

	var distance float64

	for i := 0; i < len(a); i++ {
		distance += math.Abs(a[i] - b[i])
	}

	fmt.Printf("Total distance is: %f\n", distance)
}
