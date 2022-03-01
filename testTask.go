package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

var alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func testValidity(str string) bool {
	arr := strings.Split(str, "-")

	for i, el := range arr {
		if el == "" {
			return false
		}

		_, err := strconv.Atoi(el)
		if i%2 == 0 {
			if err != nil {
				return false
			}
		} else {
			if err == nil {
				return false
			}
		}
	}

	return true
}

func averageNumber(str string) int {
	arr := strings.Split(str, "-")

	sum, count := 0, 0

	for i := 0; i < len(arr); i += 2 {
		el := arr[i]
		n, err := strconv.Atoi(el)
		if err == nil {
			sum += n
			count++
		}
	}

	return int(math.Round(float64(sum) / float64(count)))
}

func wholeStory(str string) string {
	arr := strings.Split(str, "-")

	strArr := make([]string, 0, len(arr)/2)

	for i := 1; i < len(arr); i += 2 {
		el := arr[i]
		strArr = append(strArr, el)
	}

	return strings.Join(strArr, " ")
}

func storyStats(str string) (string, string, int, []string) {
	arr := strings.Split(str, "-")

	strArr := make([]string, 0, len(arr)/2)

	min, max, total, count := len(str), 0, 0, 0
	shortest, longest := "", ""
	for i := 1; i < len(arr); i += 2 {
		el := arr[i]
		strArr = append(strArr, el)
		l := len(el)

		if l < min {
			min = l
			shortest = el
		}

		if l > max {
			max = l
			longest = el
		}

		total += l
		count++
	}

	avg := int(math.Round(float64(total) / float64(count)))

	for i, el := range strArr {
		if len(el) > avg {
			strArr[i] = el[:avg]
		} else if len(el) < avg {
			for ii := 0; ii < avg-len(el); ii++ {
				el += "a"
			}
			strArr[i] = el
		}
	}

	return shortest, longest, avg, strArr
}

func generate(valid bool) string {
	r := rand.Intn(10) + 1

	arr := make([]string, r)
	if valid {
		for i := 0; i < r; i++ {
			if i%2 == 0 {
				arr[i] = fmt.Sprintf("%v", rand.Intn(100)+1)
			} else {
				arr[i] = generateRandomString(rand.Intn(10) + 1)
			}
		}
	} else {
		for i := 0; i < r; i++ {
			arr[i] = generateRandomString(rand.Intn(10) + 1)
		}
	}
	return strings.Join(arr, "-")
}

func generateRandomString(l int) string {
	str := ""
	for i := 0; i < l; i++ {
		r := rand.Intn(len(alphabet))
		str += string(alphabet[r])
	}
	return str
}
