package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/mreed19/advent-of-code/2021/day16/go/data"
)

func main() {
	bits := data.GetData("data/input.txt")

	fmt.Println("Part 1 Solution:", part1(bits))
	fmt.Println("Part 2 Solution:", part2(bits))
}

func part1(bits []rune) int {
	defer duration(track("part1"))

	packet, _ := parsePackets(bits)

	return packet.versionSum
}

func part2(bits []rune) int {
	defer duration(track("part2"))

	packet, _ := parsePackets(bits)

	return packet.val
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

type Packet struct {
	version    int
	versionSum int
	typeID     int
	val        int
	subPackets []*Packet
	visited    bool
}

var TypeMap = map[int]func([]*Packet) int{
	0: func(subPackets []*Packet) int {
		sum := 0
		for _, packet := range subPackets {
			sum += packet.val
		}
		return sum
	},
	1: func(subPackets []*Packet) int {
		product := subPackets[0].val
		for i := 1; i < len(subPackets); i++ {
			product *= subPackets[i].val
		}
		return product
	},
	2: func(subPackets []*Packet) int {
		min := subPackets[0].val
		for i := 1; i < len(subPackets); i++ {
			if subPackets[i].val < min {
				min = subPackets[i].val
			}
		}
		return min
	},
	3: func(subPackets []*Packet) int {
		max := subPackets[0].val
		for i := 1; i < len(subPackets); i++ {
			if subPackets[i].val > max {
				max = subPackets[i].val
			}
		}
		return max
	},
	5: func(subPackets []*Packet) int {
		if subPackets[0].val > subPackets[1].val {
			return 1
		}
		return 0
	},
	6: func(subPackets []*Packet) int {
		if subPackets[0].val < subPackets[1].val {
			return 1
		}
		return 0
	},
	7: func(subPackets []*Packet) int {
		if subPackets[0].val == subPackets[1].val {
			return 1
		}
		return 0
	},
}

func parsePackets(bits []rune) (*Packet, []rune) {
	versionBits := bits[0:3]
	version, err := strconv.ParseInt(string(versionBits), 2, 8)
	if err != nil {
		log.Fatal(err)
	}

	typeBits := bits[3:6]
	typeID, err := strconv.ParseInt(string(typeBits), 2, 8)
	if err != nil {
		log.Fatal(err)
	}

	if typeID == 4 {
		i := 6
		breakFlag := false
		var literal []rune
		for {
			if bits[i] == '0' {
				breakFlag = true
			}
			// process literal
			literal = append(literal, bits[i+1:i+5]...)

			i += 5
			if breakFlag {
				break
			}
		}

		val, err := strconv.ParseInt(string(literal), 2, 64)
		if err != nil {
			log.Fatal(err)
		}

		if i < len(bits) {
			bits = bits[i:]
		} else {
			bits = nil
		}

		return &Packet{
			version:    int(version),
			versionSum: int(version),
			typeID:     int(typeID),
			val:        int(val),
		}, bits
	} else {
		var packet *Packet

		if bits[6] == '0' {
			subBitCount64, err := strconv.ParseInt(string(bits[7:22]), 2, 64)
			if err != nil {
				log.Fatal(err)
			}
			subBitCount := int(subBitCount64)
			subBits := bits[22 : 22+subBitCount]
			var subPackets []*Packet
			i := 0
			for len(subBits) > 0 {
				newSubPacket, newSubBits := parsePackets(subBits)
				subBits = newSubBits
				subPackets = append(subPackets, newSubPacket)
				i++
			}

			if 22+subBitCount < len(bits) {
				bits = bits[22+subBitCount:]
			} else {
				bits = nil
			}
			packet = &Packet{
				version:    int(version),
				typeID:     int(typeID),
				subPackets: subPackets,
			}
		} else {
			subPacketCount, err := strconv.ParseUint(string(bits[7:18]), 2, 64)
			if err != nil {
				log.Fatal(err)
			}

			var subPackets []*Packet
			subBits := bits[18:]
			for i := 0; i < int(subPacketCount); i++ {
				subPacket, newBits := parsePackets(subBits)
				subBits = newBits
				subPackets = append(subPackets, subPacket)
			}
			bits = subBits
			packet = &Packet{
				version:    int(version),
				typeID:     int(typeID),
				subPackets: subPackets,
			}
		}

		packet.val = TypeMap[packet.typeID](packet.subPackets)
		packet.versionSum = packet.version
		for _, subPacket := range packet.subPackets {
			packet.versionSum += subPacket.versionSum
		}

		return packet, bits
	}
}
