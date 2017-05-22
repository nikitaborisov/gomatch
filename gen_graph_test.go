package gomatch

import (
  "testing"
  "fmt"
)

func TestGenGraphLR(t *testing.T) {
  G := Gen_graph_lr(10, 10, 2)
  fmt.Println(G)
}
