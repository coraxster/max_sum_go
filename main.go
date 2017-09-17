package main

import (
	"fmt"
	"flag"
	"encoding/csv"
	"bufio"
	"os"
	"log"
	"strconv"
)

func main()  {
	var searchSum = flag.Int64("sum", 200, "max sum search for")
	var csvPath = flag.String("products.csv", "products.csv", "path of csv file")
	flag.Parse()

	csvFile, _ := os.Open(*csvPath)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	coasts := make([]int64, len(records))
	for k, record := range records {
		coasts[k], err = strconv.ParseInt(record[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
	}

	coastIndexes := dynamicMaxSumElements(coasts, *searchSum)

	var total int64 = 0
	for _, coastIndex := range coastIndexes {
		fmt.Printf("%s - %v \n", records[coastIndex][0], records[coastIndex][1])
		total = total + coasts[coastIndex]
	}
	fmt.Printf("Total: %v \n", total)


}



func dynamicMaxSumElements(input []int64, searchSum int64) []int {
	sumMaps := make(map[int64][]int, searchSum)
	sumMaps[0] = make([]int, 0)
	for i, v := range input {
		for j := searchSum - v; j >= 0; j-- {
			if sMap, ok := sumMaps[j]; ok {
				newSMap := append(sMap, i)
				sumMaps[j + v] = newSMap
			}
		}
	}

	var maxSum int64
	for k := range sumMaps {
		if k > maxSum {
			maxSum = k
		}
	}

	return sumMaps[maxSum]
}
