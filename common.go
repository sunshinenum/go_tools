package main

import (
	"os"
	"fmt"
	"sort"
	"bufio"
	"strings"
)

// rankByWordCount
func rankByWordCount(wordFrequencies map[string]int) PairList{
  pl := make(PairList, len(wordFrequencies))
  i := 0
  for k, v := range wordFrequencies {
    pl[i] = Pair{k, v}
    i++
  }
  sort.Sort(sort.Reverse(pl))
  return pl
}

type Pair struct {
  Key string
  Value int
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }


// read file to slice
func loadFileToSlice(filePath string) []string {
	ip, _ := os.Open(filePath)
	defer ip.Close()
	scanner := bufio.NewScanner(ip)
	line := ""
	lines := make([]string, 0)
	for scanner.Scan() {
		line = scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Err while load %s\n", filePath)
	}
	return lines
}

// dump file from slice
func dumpFileFromSlice(filePath string, data []string) {
	op, _ := os.Create(filePath)
	defer op.Close()
	for _, s := range data {
		op.WriteString(s + "\n")
	}
}

// read file to 2d slice
func loadFileTo2DSlice(filePath string, sep string) [][]string {
	ip, _ := os.Open(filePath)
	defer ip.Close()
	scanner := bufio.NewScanner(ip)
	line := ""
	lines := make([][]string, 0)
	for scanner.Scan() {
		line = scanner.Text()
		secs := strings.Split(line, sep)
		lines = append(lines, secs)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Err while load %s\n", filePath)
	}
	return lines
}

// dump file from 2d slice
func dumpFileFrom2DSlice(filePath string, data [][]string, sep string) {
	op, _ := os.Create(filePath)
	defer op.Close()
	line := ""
	for _, secs := range data {
		line = joinString(secs, sep)
		op.WriteString(line + "\n")
	}
}

func joinString(secs []string, sep string) string {
	s := ""
	for _, sec := range secs {
		s += (sec + sep)
	}
	return s[0:len(s)-1]
}

func main() {
	// rankByWordCount
	d := make(map[string]int)
	d["a"] = 3
	d["b"] = 5
	d["c"] = 2
	l := rankByWordCount(d)
	for i, v := range(l) {
		fmt.Println(i, v.Key, v.Value)
	}
	// loadFileToSlice
	l1 := loadFileToSlice("a")
	fmt.Println(l1)
	l2 := loadFileTo2DSlice("a", "\t")
	fmt.Println(l2)
	dumpFileFromSlice("b", l1)
	fmt.Println(joinString(l2[0], " "))
	dumpFileFrom2DSlice("c", l2, " ")
}
