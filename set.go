package matching

// a set of nodes
type set map[Node]struct{}


func (s set) in(n Node) bool {
  _, ok := s[n]
  return ok
}

func (s set) add(n Node) {
  s[n] = struct{}{}
}

func (s set) remove(n Node) {
  delete(s, n)
}
