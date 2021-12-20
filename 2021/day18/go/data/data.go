package data

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type SnailNumber struct {
	IsPair bool
	Val    int
	Left   *SnailNumber
	Right  *SnailNumber
	Level  int
}

func NewSnailNumber(line string) *SnailNumber {
	if !strings.Contains(line, ",") {
		val, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		return &SnailNumber{
			Val: val,
		}
	}
	r := regexp.MustCompile(`\[(.*)\]`)
	matches := r.FindStringSubmatch(line)
	numberString := []rune(matches[1])
	braceCount := 0
	for i := 0; i < len(numberString); i++ {
		if numberString[i] == ',' && braceCount == 0 {
			return &SnailNumber{
				IsPair: true,
				Left:   NewSnailNumber(string(numberString[:i])),
				Right:  NewSnailNumber(string(numberString[i+1:])),
			}
		} else if numberString[i] == '[' {
			braceCount++
		} else if numberString[i] == ']' {
			braceCount--
		}
	}

	panic("Invalid SnailNumber")
}

func (s *SnailNumber) Copy() *SnailNumber {
	if !s.IsPair {
		newNum := *s
		return &newNum
	}
	return &SnailNumber{
		IsPair: s.IsPair,
		Level:  s.Level,
		Left:   s.Left.Copy(),
		Right:  s.Right.Copy(),
	}
}

func AddSnailNumbers(s1, s2 *SnailNumber) *SnailNumber {
	s := &SnailNumber{
		IsPair: true,
		Left:   s1,
		Right:  s2,
	}
	s.Reduce()
	return s
}

func (s *SnailNumber) Reduce() {
	var stack []*SnailNumber
	level := 0

	exploding := false
	var explodeRight int
	var dfs func(*SnailNumber)
	dfs = func(number *SnailNumber) {
		if !number.IsPair {
			number.Level = level
			if exploding {
				number.Val += explodeRight
				exploding = false
			}

			stack = append(stack, number)
		} else {
			level++
			if level == 5 {
				if exploding {
					number.Left.Val += explodeRight
				}
				exploding = true
				explodeRight = number.Right.Val

				if len(stack) > 0 {
					stack[len(stack)-1].Val += number.Left.Val
				}

				number.Level = level - 1
				number.Val = 0
				number.IsPair = false
				number.Left = nil
				number.Right = nil
				stack = append(stack, number)
				level--
				return
			}
			dfs(number.Left)
			dfs(number.Right)
			level--
		}
	}

	dfs(s)

	for i := 0; i < len(stack); i++ {
		curr := stack[i]
		if curr.Val >= 10 {
			mid := float64(curr.Val) / 2
			leftVal := int(math.Floor(mid))
			rightVal := int(math.Ceil(mid))
			if curr.Level == 4 {
				if i < len(stack)-1 {
					stack[i+1].Val += rightVal
				}
				if i >= 1 {
					stack[i-1].Val += leftVal
					i -= 2
				}
				curr.Val = 0
			} else {
				curr.Left = &SnailNumber{
					Val:   leftVal,
					Level: curr.Level + 1,
				}
				curr.Right = &SnailNumber{
					Val:   rightVal,
					Level: curr.Level + 1,
				}
				curr.Val = 0
				curr.IsPair = true
				stack = append(stack[:i], append([]*SnailNumber{curr.Left, curr.Right}, stack[i+1:]...)...)
				if leftVal >= 10 {
					i--
				} else if rightVal < 10 {
					i++
				}
			}
		}
	}
}

func (s *SnailNumber) Magnitude() int {
	if !s.IsPair {
		return s.Val
	}
	return 3*s.Left.Magnitude() + 2*s.Right.Magnitude()
}

func (s *SnailNumber) Print() {
	s.printHelper(true)
}

func (s *SnailNumber) printHelper(newLine bool) {
	if s.IsPair {
		fmt.Print("[")
		s.Left.printHelper(false)
		fmt.Print(",")
		s.Right.printHelper(false)
		fmt.Print("]")
	} else {
		fmt.Printf("%d", s.Val)
	}
	if newLine {
		fmt.Print("\n")
	}
}

func GetData(filename string) []*SnailNumber {
	var numbers []*SnailNumber

	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range strings.Split(string(file), "\n") {
		if line != "" {
			numbers = append(numbers, NewSnailNumber(line))
		}
	}

	return numbers
}
