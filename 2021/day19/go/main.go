package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mreed19/advent-of-code/2021/day19/go/data"
)

func main() {
	fmt.Println("Part 1 Solution:", part1("data/input.txt"))
	fmt.Println("Part 2 Solution:", part2("data/input.txt"))
}

func part1(filename string) int {
	defer duration(track("part1"))

	scanners := data.GetData(filename)

	normalized := []*data.Scanner{scanners[0]}
	unnormalized := scanners[1:]

Main:
	for len(unnormalized) > 0 {
		for uIndex, u := range unnormalized {
			for _, n := range normalized {
				pointMatch := 0

				// compare all points in unnormalized scanner with all points in normalized scanner
				for _, pu := range u.Points {
					for _, pn := range n.Points {
						matchCount := 0

						for _, eu := range pu.Edges {
							for _, en := range pn.Edges {
								if data.SetEqual(eu, en) {
									matchCount++
								}
							}
						}

						if matchCount >= 11 {
							pointMatch++
						}
					} // for each point in normalized scanner
				} // for each pount in unnormalized scanner

				// If 12 or more points match for the magnitude then scanners have overlap
				if pointMatch >= 12 {
					// try each possible rotation to see which rotations match
					for _, rotation := range data.Rotations {
						rotationPointMatch := 0
						var translation [3]int

						for _, pu := range u.Points {
							for _, pn := range n.Points {
								matchCount := 0
								for _, eu := range pu.Edges {
									for _, en := range pn.Edges {
										if data.Equal(data.Rotate(eu, rotation), en) {
											matchCount++
										}
									}
								}

								if matchCount >= 11 {
									rotated := data.Rotate(pu.Coords, rotation)
									translation = [3]int{
										pn.Coords[0] - rotated[0],
										pn.Coords[1] - rotated[1],
										pn.Coords[2] - rotated[2],
									}
									rotationPointMatch++
								}

							} // for each point in normalized scanner
						} // for each point in unnormalized scanner compared with point in normalized scanner

						// if the rotation matches for 12 or more points
						if rotationPointMatch >= 12 {
							var newPoints []*data.Point

							for _, pu := range u.Points {
								var rotatedEdges [][3]int
								for _, edge := range pu.Edges {
									rotatedEdges = append(rotatedEdges, data.Rotate(edge, rotation))
								}
								newPoints = append(newPoints, &data.Point{
									Coords: data.Translate(data.Rotate(pu.Coords, rotation), translation),
									Edges:  rotatedEdges,
								})
							}

							matchCount := 0
							for _, pu := range newPoints {
								for _, pn := range n.Points {
									if data.Equal(pu.Coords, pn.Coords) {
										matchCount++
									}
								}
							}

							if matchCount >= 12 {
								u.Points = newPoints
								normalized = append(normalized, u)
								unnormalized[uIndex] = unnormalized[len(unnormalized)-1]
								unnormalized = unnormalized[:len(unnormalized)-1]
								continue Main
							}
						} // end if rotation matches for 12 or more points
					} // for each rotation
				} // end if scanners match
			} //for each normalized
		} // for each unnormalized
		log.Panicf("Still Unmatched: %d\n", len(unnormalized))
	}

	beaconMap := make(map[[3]int]struct{})
	for _, scanner := range normalized {
		for _, point := range scanner.Points {
			if _, ok := beaconMap[point.Coords]; !ok {
				beaconMap[point.Coords] = struct{}{}
			}
		}
	}

	return len(beaconMap)
}

func part2(filename string) int {
	defer duration(track("part2"))
	scanners := data.GetData(filename)

	normalized := []*data.Scanner{scanners[0]}
	unnormalized := scanners[1:]

Main:
	for len(unnormalized) > 0 {
		for uIndex, u := range unnormalized {
			for _, n := range normalized {
				pointMatch := 0

				// compare all points in unnormalized scanner with all points in normalized scanner
				for _, pu := range u.Points {
					for _, pn := range n.Points {
						matchCount := 0

						for _, eu := range pu.Edges {
							for _, en := range pn.Edges {
								if data.SetEqual(eu, en) {
									matchCount++
								}
							}
						}

						if matchCount >= 11 {
							pointMatch++
						}
					} // for each point in normalized scanner
				} // for each pount in unnormalized scanner

				// If 12 or more points match for the magnitude then scanners have overlap
				if pointMatch >= 12 {
					// try each possible rotation to see which rotations match
					for _, rotation := range data.Rotations {
						rotationPointMatch := 0
						var translation [3]int

						for _, pu := range u.Points {
							for _, pn := range n.Points {
								matchCount := 0
								for _, eu := range pu.Edges {
									for _, en := range pn.Edges {
										if data.Equal(data.Rotate(eu, rotation), en) {
											matchCount++
										}
									}
								}

								if matchCount >= 11 {
									rotated := data.Rotate(pu.Coords, rotation)
									translation = [3]int{
										pn.Coords[0] - rotated[0],
										pn.Coords[1] - rotated[1],
										pn.Coords[2] - rotated[2],
									}
									rotationPointMatch++
								}

							} // for each point in normalized scanner
						} // for each point in unnormalized scanner compared with point in normalized scanner

						// if the rotation matches for 12 or more points
						if rotationPointMatch >= 12 {
							var newPoints []*data.Point

							for _, pu := range u.Points {
								var rotatedEdges [][3]int
								for _, edge := range pu.Edges {
									rotatedEdges = append(rotatedEdges, data.Rotate(edge, rotation))
								}
								newPoints = append(newPoints, &data.Point{
									Coords: data.Translate(data.Rotate(pu.Coords, rotation), translation),
									Edges:  rotatedEdges,
								})
							}

							matchCount := 0
							for _, pu := range newPoints {
								for _, pn := range n.Points {
									if data.Equal(pu.Coords, pn.Coords) {
										matchCount++
									}
								}
							}

							if matchCount >= 12 {
								u.Points = newPoints
								u.Position = translation
								normalized = append(normalized, u)
								unnormalized[uIndex] = unnormalized[len(unnormalized)-1]
								unnormalized = unnormalized[:len(unnormalized)-1]
								continue Main
							}
						} // end if rotation matches for 12 or more points
					} // for each rotation
				} // end if scanners match
			} //for each normalized
		} // for each unnormalized
		log.Panicf("Still Unmatched: %d\n", len(unnormalized))
	}

	maxDistance := 0
	for i := 0; i < len(normalized)-1; i++ {
		for j := i + 1; j < len(normalized); j++ {
			posI := normalized[i].Position
			posJ := normalized[j].Position
			distance := data.Abs(posI[0]-posJ[0]) +
				data.Abs(posI[1]-posJ[1]) +
				data.Abs(posI[2]-posJ[2])
			if distance > maxDistance {
				maxDistance = distance
			}
		}
	}

	return maxDistance
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
