package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	numbers := make([]int, 0)

	for scanner.Scan() {
		word := scanner.Text()
		num, _ := strconv.Atoi(strings.Split(word, "\n")[0])
		numbers = append(numbers, num)

		if len(numbers) == 1 {
			fmt.Printf("%d %d\n", numbers[0], numbers[0])
			continue
		}

		from, to := guess(numbers)

		fmt.Printf("%d %d\n", from, to)
	}
}

func guess(numbers []int) (int, int) {
	mean := round(average(numbers))
	stdDev := round(math.Sqrt((variance(numbers))))

	from := mean - stdDev
	to := mean + stdDev

	return from, to
}

func average(numbers []int) float64 {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return float64(sum) / float64(len(numbers))
}

func median(numbers []int) float64 {
	sort.Ints(numbers)
	if len(numbers)%2 != 0 {
		return float64(numbers[len(numbers)+1/2])
	}
	l1 := (len(numbers) / 2) - 1
	l2 := (len(numbers) / 2)
	return float64((numbers[l1] + numbers[l2])) / 2
}

func variance(numbers []int) float64 {
	mean := average(numbers)
	var result float64
	for _, v := range numbers {
		dif := (float64(v) - mean)
		result += float64(dif) * float64(dif)
	}
	return result / float64(len(numbers))
}

func round(x float64) int {
	return int(math.Round(x))
}
