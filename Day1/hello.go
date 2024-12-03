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
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	list1 := []int{}
	list2 := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			num1, err1 := strconv.Atoi(fields[0])
			num2, err2 := strconv.Atoi(fields[1])
			if err1 != nil || err2 != nil {
				panic("Should not be here")
			}

			list1 = append(list1, num1)
			list2 = append(list2, num2)
		}
	}

	sort.Ints(list1)
	sort.Ints(list2)

	//fmt.Println(list1)
	//fmt.Println(list2)

	total := 0

	for index, num1 := range list1 {
		num2 := list2[index]
		total += int(math.Abs(float64(num1 - num2)))
	}

	fmt.Println(total)

	countMap := make(map[int]int)

	for _, num2 := range list2 {
		countMap[num2]++
	}

	total2 := 0
	for _, num1 := range list1 {
		total2 += num1 * countMap[num1]
	}

	fmt.Println(total2)
}
