package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

	reports := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		report := []int{}
		for _, field := range fields {
			value, _ := strconv.Atoi(field)
			report = append(report, value)
		}

		reports = append(reports, report)
	}
	fmt.Println("Total number of reports:", len(reports))
	safeReports := 0
	unsafeIndexes := map[int]int{}
	unsafeReports := [][]int{}
	for row, report := range reports {
		previousNum := 0
		safeReport := true
		isIncrease := true
		for index, currentNum := range report {
			if index == 0 {
				diff := report[1] - currentNum
				if diff == 0 {
					safeReport = false
					break
				}
				isIncrease = (report[1] - currentNum) > 0
			} else if index > 0 {
				diff := int(math.Abs(float64(currentNum - previousNum)))
				if diff == 0 {
					safeReport = false
				} else if isIncrease {
					safeReport = currentNum > previousNum && diff <= 3
				} else if !isIncrease {
					safeReport = currentNum < previousNum && diff <= 3
				}
			}
			previousNum = currentNum
			if !safeReport {
				unsafeIndexes[row] = index
				unsafeReports = append(unsafeReports, report)
				break
			}
		}
		if safeReport {
			fmt.Println("safe:", report)
			safeReports++
		}
	}

	fmt.Println(safeReports)

	for row, report := range unsafeReports {
		previousNum := 0
		safeReport := true
		isIncrease := true
		unSafeIndex := unsafeIndexes[row]
		for index, currentNum := range report {
			if index == unSafeIndex {
				continue
			}
			if index == 0 {
				diff := report[1] - currentNum
				if diff == 0 {
					safeReport = false
					break
				}
				isIncrease = (report[1] - currentNum) > 0
			} else if index > 0 {
				diff := int(math.Abs(float64(currentNum - previousNum)))
				if diff == 0 {
					safeReport = false
				} else if isIncrease {
					safeReport = currentNum > previousNum && diff <= 3
				} else if !isIncrease {
					safeReport = currentNum < previousNum && diff <= 3
				}
			}
			previousNum = currentNum
			if !safeReport {
				break
			}
		}
		if safeReport {
			fmt.Println("safe:", report)
			safeReports++
		}
	}

	fmt.Println(safeReports)
}
