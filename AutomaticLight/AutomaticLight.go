package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type EmployeeTimes struct {
	timeIn, timeOut int
}

func main() {
	times := readInputFile()

	minimum := calculateMinimumUpTime(times)
	fmt.Println(minimum)
}

func calculateMinimumUpTime(times []EmployeeTimes) int {
	var minimumTimes []int
	minimumUpTime := 0
	if len(times) == 1 {
		return getMinimumUpTime(times)
	}
	for _, item := range times {
		employees := find(times, item)
		if len(employees) >= 1 { // não basta ser só um, tem que ser diferente dos últimos que estavam na empresa.
			fmt.Println(employees)
			minimumUpTime = getMinimumUpTime(employees)
		} else {
			minimumTimes = append(minimumTimes, minimumUpTime)
		}
	}
	if len(minimumTimes) == 0 {
		return minimumUpTime
	}
	return max(minimumTimes)
}

func getMinimumUpTime(employees []EmployeeTimes) int {
	maxInTime := getMaxInTime(employees)
	maxOutTime := getMaxOutTime(employees)

	return maxOutTime - maxInTime
}

func find(slice []EmployeeTimes, item EmployeeTimes) []EmployeeTimes {
	var found []EmployeeTimes
	for _, sliceItem := range slice {
		if sliceItem.timeIn <= item.timeIn && sliceItem.timeOut >= item.timeIn {
			found = append(found, sliceItem)
		}
	}
	return found
}

func max(numbers []int) int {
	max := 0
	for _, item := range numbers {
		if item > max {
			max = item
		}
	}

	return max
}

func getMaxInTime(times []EmployeeTimes) int {
	max := 0
	for _, item := range times {
		if item.timeIn > max {
			max = item.timeIn
		}
	}
	return max
}

func getMaxOutTime(times []EmployeeTimes) int {
	max := 0
	for _, item := range times {
		if item.timeOut > max {
			max = item.timeOut
		}
	}
	return max
}

func readInputFile() []EmployeeTimes {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var times []EmployeeTimes
	scanner := bufio.NewScanner(file)
	firstTime := true
	for scanner.Scan() {
		employeeTimes := scanner.Text()

		words := strings.Fields(employeeTimes)
		if firstTime {
			firstTime = false
			continue
		}
		timeIn, _ := strconv.Atoi(words[0])
		timeOut, _ := strconv.Atoi(words[1])

		currentEmployeeTime := EmployeeTimes{timeIn, timeOut}
		times = append(times, currentEmployeeTime)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return times
}
