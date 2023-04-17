package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countDaysTogether("10-01", "10-02", "11-01", "12-31"))
}

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	daysOfMonth := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	aliceBegin := dayOfYear(arriveAlice, daysOfMonth)
	aliceEnd := dayOfYear(leaveAlice, daysOfMonth)
	bobBegin := dayOfYear(arriveBob, daysOfMonth)
	bobEnd := dayOfYear(leaveBob, daysOfMonth)
	beginMax := max(aliceBegin, bobBegin)
	endMin := min(aliceEnd, bobEnd)
	fmt.Println(aliceBegin, aliceEnd, bobBegin, bobEnd, beginMax, endMin)
	return max(endMin-beginMax+1, 0)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func dayOfYear(monthDay string, daysOfMonth []int) int {
	month, _ := strconv.Atoi(monthDay[0:2])
	day, _ := strconv.Atoi(monthDay[3:])

	monthDays := 0
	for i := 0; i < month-1; i++ {
		monthDays += daysOfMonth[i]
	}
	return monthDays + day
}
