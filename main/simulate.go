package main

import (
    "github.com/nikitaborisov/gomatch"
    "fmt"
)

func main() {
    n := 100
    degree := 2
    G := gomatch.Gen_graph_lr(n, n, degree)
    //fmt.Println(G)
    nmatch, LR, _ := gomatch.HopcroftKarp(G)
    if nmatch != n {
      fmt.Println("Can't find perfect matching!", nmatch)
      return
    }
    M := make(map[gomatch.Node]gomatch.Node)
    for l, r := range(LR) {
      M[l] = r
    }
    //fmt.Println(M)

    matchable := gomatch.MatchableEdges(G, M)
    //fmt.Println(matchable)

    Gmatch := make(gomatch.Graph)
    for _, e := range matchable {
      cur := Gmatch[e.From]
      if cur == nil {
        cur = make([]gomatch.Node, 0)
      }
      Gmatch[e.From] = append(cur, e.To)
    }
    //fmt.Println(Gmatch)

    histogram := make([]int, degree)

    for _, es := range Gmatch {
      histogram[len(es)-1] += 1
    }
    fmt.Println(histogram)
}
