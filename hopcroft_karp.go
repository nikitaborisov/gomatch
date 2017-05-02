package matching

type internals struct {
  matchLR map[Node]Node
  matchRL map[Node]Node
  lev map[Node]int
  G Graph
}

func bfs(state *internals) bool {
  L := make([]Node, 0)
  state.lev = make(map[Node]int)
  for node, _ := range state.G {
    if _, ok := state.matchLR[node]; !ok {
      L = append(L, node)
    }
  }
  depth := 0
  seen := make(map[Node]bool)
  for len(L) > 0 {
    for _, l := range L {
      state.lev[l] = depth
    }
    L1 := make(map[Node]bool)
    for _, l := range L {
      for _, r := range state.G[l] {
        if _, ok := seen[r]; !ok {
            seen[r] = true
            if _, ok2 := state.matchRL[r]; !ok2 {
              return true
            }
            L1[state.matchRL[r]] = true
        }
      }
      depth += 1
      L = make([]Node, len(L1))
      i := 0
      for node, _ := range L1 {
        L[i] = node
        i++
      }
    }
  }
  return false
}

func dfs(l Node, targl int, state *internals) bool {
  if ll, ok := state.lev[l]; !ok || ll != targl {
    return false
  }
  delete(state.lev, l)
  for _, r := range state.G[l] {
    _, ok := state.matchRL[r]
    if !ok || dfs(state.matchRL[r], targl+1, state) {
      state.matchLR[l] = r
      state.matchRL[r] = l
      return true
    }
  }
  return false
}

func HopcroftKarp(G Graph) (nmatch int, matchLR map[Node]Node,
  matchRL map[Node]Node) {
  nmatch = 0
  matchLR = make(map[Node]Node)
  matchRL = make(map[Node]Node)
  state := internals{matchLR, matchRL, make(map[Node]int), G}
  for bfs(&state) {
    for l := range(G) {
      _, ok := matchLR[l]
      if !ok && dfs(l,0, &state) {
        nmatch += 1
      }
    }
  }
  return
}
