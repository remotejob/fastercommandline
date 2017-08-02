package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func processFile(filePath string, keyIndex, valueIndex int) {
	delim := "\t"

	fileHandle, err := os.Open(filePath)
	defer fileHandle.Close()

	maxFieldIndex := int(math.Max(float64(keyIndex), float64(valueIndex)))
	sumByKey := make(map[string]int)

	if err != nil {
		log.Fatal(err)
	}

	fileReader := bufio.NewScanner(fileHandle)

	for fileReader.Scan() {
		fields := strings.Split(strings.TrimSpace(fileReader.Text()), delim)

		if maxFieldIndex < len(fields) {
			value, err := strconv.Atoi(fields[valueIndex])
			if err != nil {
				log.Fatal(err)
			}
			sumByKey[fields[keyIndex]] += value
		}
	}

	maxValue := 0
	maxKey := ""
	for k, v := range sumByKey {
		if v > maxValue {
			maxValue = v
			maxKey = k
		}
	}
	fmt.Println("max_key:", maxKey, "sum:", sumByKey[maxKey])

}

func main() {
	filePath := flag.String("filePath", "", "Name of the file to parse")
	keyIndex := flag.Int("keyIndex", 0, "Index of key (0 is first position)")
	valueIndex := flag.Int("valueIndex", 0, "Index of value (0 is first position)")

	flag.Parse()
	processFile(*filePath, *keyIndex, *valueIndex)

}
