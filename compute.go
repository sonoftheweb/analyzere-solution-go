package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "command expects two arguments: compute <threshold> <limit>")
		os.Exit(1)
	}

	threshold, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid threshold value: %v\n", err)
		os.Exit(1)
	}

	limit, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid limit value: %v\n", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	var cumulativeSum float64 = 0
	var outputs []float64

	for scanner.Scan() {
		inputValue, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid input value: %v\n", err)
			continue
		}

		outputValue := inputValue - threshold
		if outputValue < 0 {
			outputValue = 0
		}

		if cumulativeSum+outputValue > limit {
			if limit-cumulativeSum > 0 {
				outputValue = limit - cumulativeSum
				cumulativeSum += outputValue
			} else {
				outputValue = 0
			}
		} else {
			cumulativeSum += outputValue
		}

		outputs = append(outputs, outputValue)
	}

	for _, output := range outputs {
		fmt.Printf("%.1f\n", output)
	}
	fmt.Printf("%.1f\n", cumulativeSum)
}
