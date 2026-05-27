// Command check_coverage validates total Go test coverage.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const minCoverage = 18.0

func main() {
	file, err := os.Open("coverage.out")
	if err != nil {
		fmt.Println("failed to open coverage.out:", err)
		os.Exit(1)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("failed to close coverage file: %v", err)
		}
	}()

	var total float64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) < 3 || !strings.HasSuffix(fields[2], "%") {
			continue
		}

		value := strings.TrimSuffix(fields[2], "%")
		parsed, err := strconv.ParseFloat(value, 64)
		if err == nil {
			total = parsed
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("failed to scan coverage.out:", err)
		os.Exit(1)
	}

	if total < minCoverage {
		fmt.Printf("coverage %.2f%% is below required %.2f%%\n", total, minCoverage)
		os.Exit(1)
	}

	fmt.Printf("coverage %.2f%% meets required %.2f%%\n", total, minCoverage)
}
