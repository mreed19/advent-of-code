package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("Part 1 Solution: %d\n", part1())
	fmt.Printf("Part 2 Solution: %d\n", part2())
}

func part1() int64 {
	d := data()

	byteLen := len(d[0])

	gammaBytes := ""
	epsilonBytes := ""
	for i := 0; i < byteLen; i++ {
		oneCount := 0
		zeroCount := 0
		for j := 0; j < len(d); j++ {
			if d[j][i] == '1' {
				oneCount++
			} else {
				zeroCount++
			}
		}
		if oneCount >= zeroCount {
			gammaBytes += "1"
			epsilonBytes += "0"
		} else {
			gammaBytes += "0"
			epsilonBytes += "1"
		}
	}

	gamma, _ := strconv.ParseInt(gammaBytes, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonBytes, 2, 64)

	return gamma * epsilon
}

func part2() int64 {
	d := data()

	byteLen := len(d[0])

	oxygenBytes := d
	co2Bytes := d
	for i := 0; i < byteLen; i++ {
		if len(oxygenBytes) == 1 && len(co2Bytes) == 1 {
			break
		}

		if len(oxygenBytes) > 1 {
			var oneBytes []string
			var zeroBytes []string

			for j := 0; j < len(oxygenBytes); j++ {
				if oxygenBytes[j][i] == '1' {
					oneBytes = append(oneBytes, oxygenBytes[j])
				} else {
					zeroBytes = append(zeroBytes, oxygenBytes[j])
				}
			}

			if len(oneBytes) >= len(zeroBytes) {
				oxygenBytes = oneBytes
			} else {
				oxygenBytes = zeroBytes
			}
		}

		if len(co2Bytes) > 1 {
			var oneBytes []string
			var zeroBytes []string

			for j := 0; j < len(co2Bytes); j++ {
				if co2Bytes[j][i] == '1' {
					oneBytes = append(oneBytes, co2Bytes[j])
				} else {
					zeroBytes = append(zeroBytes, co2Bytes[j])
				}
			}

			if len(zeroBytes) <= len(oneBytes) {
				co2Bytes = zeroBytes
			} else {
				co2Bytes = oneBytes
			}
		}
	}

	gamma, _ := strconv.ParseInt(oxygenBytes[0], 2, 64)
	epsilon, _ := strconv.ParseInt(co2Bytes[0], 2, 64)

	return gamma * epsilon
}
