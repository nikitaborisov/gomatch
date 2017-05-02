package matching

import (
  "testing"
)

func TestSet(t *testing.T) {
  a := Node{"a"}
  b := Node{"b"}

  s := make(set)

  if (s.in(a) || s.in(b)) {
    t.Error("set not empty")
  }
  s.add(a)
  if (!s.in(a) || s.in(b)) {
    t.Error("set has wrong elements")
  }
  s.add(b)
  if (!s.in(a) || !s.in(b)) {
    t.Error("set has wrong elements")
  }
  if len(s) != 2 {
    t.Error("set has wrong size")
  }
  s.remove(a)
  if (s.in(a) || !s.in(b)) {
    t.Error("set has wrong elements")
  }
  s.remove(b)
  if (s.in(a) || s.in(b) || len(s) != 0) {
    t.Error("set has wrong elements")
  }
}
