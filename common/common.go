package common

import (
    "os"
    "fmt"
    "sort"
    "bufio"
    "math/rand"
    "strings"
)

// rankByWordCount {string: int}
type Pair struct {
    Key string
    Value int
}

type PairList []Pair
func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }

func RankByWordCount(wordFrequencies map[string]int) PairList{
    pl := make(PairList, len(wordFrequencies))
    i := 0
    for k, v := range wordFrequencies {
        pl[i] = Pair{k, v}
        i++
    }
    sort.Sort(sort.Reverse(pl))
    return pl
}

// read file to slice
func LoadFileToSlice(filePath string) []string {
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
func DumpFileFromSlice(filePath string, data []string) {
    op, _ := os.Create(filePath)
    defer op.Close()
    for _, s := range data {
        op.WriteString(s + "\n")
    }
}

// read file to 2d slice
func LoadFileTo2DSlice(filePath string, sep string) [][]string {
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

// dump file from 2d slice (str)
func DumpFileFrom2DSlice(filePath string, data [][]string, sep string) {
    op, _ := os.Create(filePath)
    defer op.Close()
    line := ""
    for _, secs := range data {
        line = JoinString(secs, sep)
        op.WriteString(line + "\n")
    }
}

// dump file from 2d slice (f32)
func DumpFileFrom2DSliceF32(filePath string, data [][]float32, sep string) {
    op, _ := os.Create(filePath)
    defer op.Close()
    line := ""
    for _, secs := range data {
        line = JoinF32(secs, sep)
        op.WriteString(line + "\n")
    }
}

// sep.join(secs)  str
func JoinString(secs []string, sep string) string {
    s := ""
    for _, sec := range secs {
        s += (sec + sep)
    }
    return s[0:len(s)-len(sep)]
}

// sep.join(secs) F32
func JoinF32(secs []float32, sep string) string {
    s := ""
    flt_str := ""
    for _, sec := range secs {
        flt_str = fmt.Sprintf("%.4f", sec)
        s += (flt_str + sep)
    }
    return s[0:len(s)-len(sep)]
}

// new 2D float slice
func New2DSliceF64(a int, b int) [][]float64 {
    ret := make([][]float64, a)
    for i:=0; i<a; i++ {
        ret[i] = make([]float64, b)
    }
    return ret
}

func New2DSliceF64Rnd(a int, b int) [][]float64 {
    ret := make([][]float64, a)
    for i:=0; i<a; i++ {
        ret[i] = make([]float64, b)
        for j := range ret[i] {
            ret[i][j] = -0.05 + rand.Float64()*0.1
        }
    }
    return ret
}

func New2DSliceF32(a int, b int) [][]float32 {
    ret := make([][]float32, a)
    for i:=0; i<a; i++ {
        ret[i] = make([]float32, b)
    }
    return ret
}

func New2DSliceF32Rnd(a int, b int) [][]float32 {
    ret := make([][]float32, a)
    for i:=0; i<a; i++ {
        ret[i] = make([]float32, b)
        for j := range ret[i] {
            ret[i][j] = -0.05 + rand.Float32()*0.1
        }
    }
    return ret
}

// new 2D int slice
func New2DSliceI64(a int, b int) [][]int64 {
    ret := make([][]int64, a)
    for i:=0; i<a; i++ {
        ret[i] = make([]int64, b)
    }
    return ret
}

func New2DSliceI32(a int, b int) [][]int32 {
    ret := make([][]int32, a)
    for i:=0; i<a; i++ {
        ret[i] = make([]int32, b)
    }
    return ret
}

