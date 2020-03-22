package main
import (
    "fmt"
    common "github.com/sunshinenum/go_tools/common"
)

func main() {
    // rankByWordCount
    d := make(map[string]int)
    d["a"] = 3
    d["b"] = 5
    d["c"] = 2
    st := common.RankByWordCount(d)
    for i, pair := range(st) {
        fmt.Println(i, pair.Key, pair.Value)
    }

    // join str
    secs := []string{"hello", "world"}
    fmt.Println(common.JoinString(secs, "\t"))
    // join floats arr to str
    secsf := []float32{0.1, 0.2}
    fmt.Println(common.JoinF32(secsf, "\t"))

    // DumpFileFrom2DSliceF32
    arr := [][]float32{{1.0, 2.0}, {3.0, 4.0}}
    common.DumpFileFrom2DSliceF32("arr", arr, "\t")
}
