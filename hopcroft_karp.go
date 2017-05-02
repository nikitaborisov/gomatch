package matching

type Matching map[Node]Node

type internals struct {
  matchLR Matching
  matchRL Matching
  lev map[Node]int
  G Graph
}


func bfs(state *internals) bool {
  L := make(set)
  state.lev = make(map[Node]int)
  for node, _ := range state.G {
    if _, ok := state.matchLR[node]; !ok {
      L.add(node)
    }
  }
  depth := 0
  seen := make(set)
  for len(L) > 0 {
    for l := range L {
      state.lev[l] = depth
    }
    L1 := make(set)
    for l := range L {
      for _, r := range state.G[l] {
        if !seen.in(r) {
            seen.add(r)
            if _, ok2 := state.matchRL[r]; !ok2 {
              return true
            }
            L1.add(state.matchRL[r])
        }
      }
      depth += 1
      L = L1
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

func HopcroftKarp(G Graph) (nmatch int, matchLR Matching,
  matchRL Matching) {
  nmatch = 0
  matchLR = make(Matching)
  matchRL = make(Matching)
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
