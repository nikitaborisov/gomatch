package matching

import (
    "github.com/looplab/tarjan"
)

type derived_node struct {
   index int
}



// to match with tarjan's graph structure
type derived_graph map[interface{}]([]interface{})

func make_derived_graph(left_nodes []Node, right_indices map[Node]int,
  G Graph, M Matching) (H derived_graph) {
  H = make(derived_graph)
  for i, _ := range left_nodes {
    H[derived_node{i}] = make([]interface{}, 0)
  }

  for i, ln := range left_nodes {
    cur_node := derived_node{i}
    for _, rn := range G[ln] {
      if rn == M[ln] {
        // skip self-links
        continue
      }
      ri := right_indices[rn]
      dest_node := derived_node{ri}
      H[cur_node] = append(H[cur_node], dest_node)
    }
  }
  return
}

func left_right_nodes(G Graph, M Matching) (left_nodes []Node,
  right_nodes []Node, right_indices map[Node]int) {
  left_nodes = make([]Node, len(G))
  right_nodes = make([]Node, len(M))
  right_indices = make(map[Node]int, len(M))
  li := 0

  for left, right := range M {
    left_nodes[li] = left
    right_nodes[li] = right
    right_indices[right] = li
    li += 1
  }

  ri := li
  for left, rights := range G {
    if _, ok := M[left]; !ok {
      left_nodes[li] = left
      li += 1
    }
    for _, right := range rights {
      if _, ok := right_indices[right]; !ok {
        right_nodes = append(right_nodes, right)
        right_indices[right] = ri
        ri += 1
      }
    }
  }
  return
}

type edge struct {
  to Node
  from Node
}

func list_matchable(G Graph, M Matching) (matchable []edge) {
  left_nodes, right_nodes, right_indices := left_right_nodes(G, M)
  DG := make_derived_graph(left_nodes, right_indices, G, M)
  components := tarjan.Connections(DG)
  comp_index := make(map[interface{}]int)
  for i, comp := range components {
    for _, node := range comp {
      comp_index[node] = i
    }
  }
  matchable = make([]edge, len(M))
  i := 0
  for l, r := range M {
    matchable[i] = edge{l, r}
    i += 1
  }
  for n, edges := range DG {
    left_index := n.(derived_node).index
    for _, e := range edges {
      if comp_index[n] == comp_index[e] {
        right_index := e.(derived_node).index
        matchable = append(matchable, edge{left_nodes[left_index], right_nodes[right_index]})
      }
    }
  }
  return
}
