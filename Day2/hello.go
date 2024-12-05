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
	unsafeReports := [][]int{}
	for _, report := range reports {
		if isSafeReport(report, false) {
			fmt.Println("safe:", report)
			safeReports++
		} else {
			unsafeReports = append(unsafeReports, report)
		}
	}

	fmt.Println("completely safe:", safeReports)

	for _, report := range unsafeReports {
		if isSafeReport(report, true) {
			fmt.Println("unsafe, but now safe:", report)
			safeReports++
		}
	}

	fmt.Println("total safe w/ error correction:", safeReports)
}

func isSafeReport(report []int, allowUnsafe bool) bool {
	previousNum := 0
	safeReport := true
	trends := []int{}
	for index, currentNum := range report {
		if index == 0 {
			previousNum = currentNum
			continue
		}
		trend := currentNum - previousNum
		trends = append(trends, trend)
		magnitude := int(math.Abs(float64(trend)))
		safeReport = allSameSign(trends) && magnitude <= 3 && magnitude > 0
		previousNum = currentNum

		if !safeReport {
			if allowUnsafe {
				allowUnsafe = false
				fmt.Println("apply error correction for report:", report)
				report1 := removeIndex(report, index)
				fmt.Println("apply error correction for report:", report)
				fmt.Println("report1", report1)
				safeReport = isSafeReport(report1, allowUnsafe)
				if safeReport {
					fmt.Println("report1 safe")
					return safeReport
				}
				fmt.Println("report1 not safe")
				report2 := removeIndex(report, index-1)
				fmt.Println("report2", report2)
				safeReport = isSafeReport(report2, allowUnsafe)
				if safeReport {
					fmt.Println("report2 safe")
					return safeReport
				}
				fmt.Println("report2 not safe")
				report3 := removeIndex(report, 0)
				fmt.Println("report3", report3)
				safeReport = isSafeReport(report3, allowUnsafe)
				if safeReport {
					fmt.Println("report3 safe")
					return safeReport
				}
				fmt.Println("report3 not safe")
			}
			break
		}
	}

	return safeReport
}

func removeIndex(slice []int, index int) []int {
	newSlice := make([]int, len(slice)-1)
	copy(newSlice, slice[:index])
	copy(newSlice[index:], slice[index+1:])
	return newSlice
}

func allSameSign(arr []int) bool {
	if len(arr) == 0 {
		return true
	}

	firstSign := arr[0] > 0
	for _, num := range arr {
		if (num > 0) != firstSign && num != 0 {
			return false
		}
	}
	return true
}
