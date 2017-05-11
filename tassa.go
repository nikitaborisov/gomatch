package matching

type derived_node struct {
  left_node Node
  right_node Node
}

type derived_graph map[derived_node]([]derived_node)

func make_derived_graph(G Graph, M Matching) (H derived_graph) {
  left_nodes := make([]Node, len(G))
  i := 0
  for n := range G {
    left_nodes[i] = n
    i++
  }
  right_indices := make(map[Node]int, len(G))
  H = make(derived_graph)
  for i, n := range left_nodes {
    right_indices[n] = i
    H[derived_node{n,M[n]}] = make([]derived_node, 0)
  }

  for _, ln := range left_nodes {
    cur_node := derived_node{ln,M[ln]}
    for _, rn := range G[ln] {
      if rn == M[ln] {
        // skip self-links
        continue
      }
      index := right_indices[rn]
      dest_ln := left_nodes[index]
      dest_node := derived_node{dest_ln,M[dest_ln]}
      H[cur_node] = append(H[cur_node], dest_node)
    }
  }
  return
}
