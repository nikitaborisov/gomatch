package gomatch

import (
  "math/rand"
  "fmt"
)

type Node struct {
  name string
}

type Graph map[Node]([]Node)

func Gen_graph(left []Node, right []Node, rdegree int) (G Graph) {
  num_nodes := len(right)
  G = map[Node][]Node{}
  for i, rn := range right {
    var j int
    for {
      j = rand.Intn(num_nodes)
      if j != i {
        break
      }
    }
    G[left[i]] = []Node{rn,right[j]}
  }
  return
}

func Gen_graph_lr(nleft int, nright int, rdegree int) (G Graph) {
  left := make([]Node, nleft)
  for i := 0; i < nleft; i++ {
    left[i].name = fmt.Sprintf("L%d", i)
  }
  right := make([]Node, nright)
  for i := 0; i < nright; i++ {
    right[i].name = fmt.Sprintf("R%d", i)
  }
  return Gen_graph(left, right, rdegree)
}
