package matching

import (
  "testing"
  "fmt"
)

func TestDerivedGraph(t *testing.T) {
  G := make(Graph)
  LN := []Node{Node{"1"}, Node{"2"}, Node{"3"}, Node{"4"}}
  RN := []Node{Node{"10"}, Node{"11"}, Node{"12"}, Node{"13"}}
  M := make(Matching)
  for i := 0; i < 4; i++ {
    G[LN[i]] = []Node{LN[i],LN[(i+1)%4]}
    M[LN[i]] = RN[i]
  }
  DG := make_derived_graph(G, M)
  fmt.Println(DG)
}
