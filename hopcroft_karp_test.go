package gomatch

import (
  "testing"
  "fmt"
)

func TestHK100(t *testing.T) {
  runHK(100, true)
}

func runHK(n int, print bool) {
  G := Gen_graph_lr(n, n, 2)
  L0 := Node{"L0"}
  G[L0] = G[L0][1:]

  if print {
    fmt.Println(G)
  }

  nmatch, LR, _ := HopcroftKarp(G)
  if print {
    fmt.Println(nmatch, LR)
  } else {
    fmt.Println(nmatch)
  }
}

func BenchmarkHK1000000(b *testing.B) {
  for n := 0; n < b.N; n++ {
    runHK(1000000, false)
  }
}
