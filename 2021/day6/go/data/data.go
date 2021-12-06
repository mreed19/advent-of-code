package data

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const LanternFishBirth int = 8
const LanternFishReset int = 6

func GetLanternFish(filename string) ([]int, error) {
	lanternFish := make([]int, LanternFishBirth+1)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fishDays := strings.Split(scanner.Text(), ",")
		for _, fish := range fishDays {
			days, err := strconv.Atoi(fish)
			if err != nil {
				return nil, err
			}
			lanternFish[days]++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lanternFish, nil
}
