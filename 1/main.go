// Solution for: https://adventofcode.com/2024/day/1
package main

import (
	"bufio"
	"errors"
	"flag"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type list []int64

var (
	inputFile string
)

func init() {
	flag.StringVar(&inputFile, "input-file", "./input", "The file containing the input")
	flag.Parse()
}

func main() {
	leftList, rightList, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	sortLists(leftList, rightList)

	distance, err := sumDiffOfPairs(leftList, rightList)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Sum of distance between left and right list is: %d", distance)
	log.Printf("Similarity score between left and right list is: %d", calcSimilarityScore(leftList, rightList))
}

func readInput() (*list, *list, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	leftList := make(list, 0)
	rightList := make(list, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")
		if len(parts) != 2 {
			return nil, nil, errors.New("wrong length for parts")
		}

		leftNum, leftErr := strconv.ParseInt(parts[0], 10, 64)
		rightNum, rightErr := strconv.ParseInt(parts[1], 10, 64)
		if leftErr != nil || rightErr != nil {
			return nil, nil, errors.New("failed to convert numbers on line: '" + line + "'")
		}

		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	if len(leftList) != len(rightList) {
		return nil, nil, errors.New("unbalanced list lengths")
	}

	return &leftList, &rightList, nil
}

func sortLists(list1, list2 *list) {
	// sort the lists
	sortFunc := func(a, b int64) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	}

	slices.SortFunc(*list1, sortFunc)
	slices.SortFunc(*list2, sortFunc)
}

func sumDiffOfPairs(list1, list2 *list) (int64, error) {
	sum := int64(0)

	for i := 0; i < len(*list1); i++ {
		diff := (*list1)[i] - (*list2)[i]
		sum += int64(math.Abs(float64(diff)))
	}

	return sum, nil
}

func calcSimilarityScore(list1, list2 *list) int64 {
	similarity := int64(0)

	// first count the number of occurrences in each list
	list1Counts := countOccurencesInList(list1)
	list2Counts := countOccurencesInList(list2)

	// then check which numbers are in both list
	for num, list1Count := range list1Counts {
		// continue if the num isn't in list 2
		if _, ok := list2Counts[num]; !ok {
			continue
		}

		similarity += num * list2Counts[num] * list1Count
	}
	return similarity
}

func countOccurencesInList(list *list) map[int64]int64 {
	counts := make(map[int64]int64)
	for _, val := range *list {
		// if first time seeing number init; otherwise increase count by 1
		if _, ok := counts[val]; !ok {
			counts[val] = 1
			continue
		}
		counts[val]++
	}
	return counts
}
