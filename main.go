package main
import (
    common "github.com/sunshinenum/go_tools/common"
)

func main() {
    // // rankByWordCount
    // d := make(map[string]int)
    // d["a"] = 3
    // d["b"] = 5
    // d["c"] = 2
    // l := RankByWordCount(d)
    // for i, v := range(l) {
    //     fmt.Println(i, v.Key, v.Value)
    // }
    // // loadFileToSlice
    // l1 := LoadFileToSlice("a")
    // fmt.Println(l1)
    // l2 := LoadFileTo2DSlice("a", "\t")
    // fmt.Println(l2)
    // DumpFileFromSlice("b", l1)
    // fmt.Println(JoinString(l2[0], " "))
    // DumpFileFrom2DSlice("c", l2, " ")
    // // join str
    // secs := []string{"hello", "world"}
    // fmt.Println(JoinString(secs, "____"))
    arr := [][]float32{{1.0, 2.0}, {3.0, 4.0}}
    common.DumpFileFrom2DSliceF32("arr", arr, "\t")
}
