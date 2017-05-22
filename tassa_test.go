package matching

import (
  "testing"
  "fmt"
  "github.com/looplab/tarjan"
)

func TestDerivedGraph(t *testing.T) {
  G := make(Graph)
  LN := []Node{Node{"1"}, Node{"2"}, Node{"3"}, Node{"4"}}
  RN := []Node{Node{"10"}, Node{"11"}, Node{"12"}, Node{"13"}}
  M := make(Matching)
  right_indices := make(map[Node]int)
  for i := 0; i < 4; i++ {
    G[LN[i]] = []Node{RN[i],RN[(i+1)%4]}
    M[LN[i]] = RN[i]
    right_indices[RN[i]] = i
  }
  DG := make_derived_graph(LN, right_indices, G, M)
  fmt.Println("Derived graph: ", DG)
  TG := make(map[interface{}][]interface{})
  for k, v := range DG {
    TG[k] = make([]interface{}, len(v))
    for i, vv := range v {
      TG[k][i] = vv
    }
  }
  T := tarjan.Connections(TG)
  fmt.Println(T)
}

func TestLeftRight(t *testing.T) {
  G := make(Graph)
  LN := []Node{Node{"1"}, Node{"2"}, Node{"3"}, Node{"4"}}
  RN := []Node{Node{"10"}, Node{"11"}, Node{"12"}, Node{"13"}}
  M := make(Matching)
  for i := 0; i < 4; i++ {
    G[LN[i]] = []Node{RN[i],RN[(i+1)%4]}
    M[LN[i]] = RN[i]
  }
  left_nodes, right_nodes, right_indices := left_right_nodes(G, M)
  fmt.Println(left_nodes, right_nodes, right_indices)

}
