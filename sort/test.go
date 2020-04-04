package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main(){
	testSuite := getTestSuite()
	testSuiteA := getTestSuite()
	testSuiteB := getTestSuite()
	for i:=0; i< len(testSuiteA); i++ {
		sort.Ints(testSuiteA[i])
		// quick sort
		// qsort(testSuiteB[i])
		// bubble sort
		// bsort(testSuiteB[i])
		// selection sort
		// ssort(testSuiteB[i])
		// insertion sort
		// isort(testSuiteB[i])
		// shell sort
		// shsort(testSuiteB[i])
		// heap sort
		hsort(testSuiteB[i])
		if !equal(testSuiteA[i], testSuiteB[i]) {
			fmt.Printf("There is an error！\nSlice is %v\n", testSuite[i])
			fmt.Printf("The answer should be %v\nbut your answer is %v", testSuiteA[i], testSuiteB[i])
			return
		}
	}
	fmt.Println("Your algorithm is right!")
}

func getTestSuite() (testSuite [][]int){
	fi, err := os.Open("sort/testcase")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		li, _, c := br.ReadLine()
		if c == io.EOF {
			fmt.Errorf("IO is error")
			return
		}
		if len(li) == 0{
			testSuite = append(testSuite, []int{})
			continue
		}
		ts := strings.Split(string(li), ",")
		testcase := []int{}
		for _, v := range ts {
			num, _ := strconv.ParseInt(strings.Trim(v, " "), 10, 0)
			testcase = append(testcase, int(num))
		}
		testSuite = append(testSuite, testcase)
	}
	return
}

func equal(a []int, b []int) bool{
	if len(a) != len(b) {
		return false
	}
	for i:=0; i<len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

