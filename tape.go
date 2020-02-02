package main

type tape struct {
  values []uint8
}

func newTape () tape {
  var val []uint8
  val = append(val, 0)
  return tape{
    values: val,
  }
}

func (t *tape) addToEnd (val uint8) {
  t.values = append(t.values, val)
}

func (t *tape) addToBeginning (val uint8) {
  t.values = append(t.values, 0)
  for i := len(t.values) - 1; i > 0; i-- {
    t.values[i] = t.values[i - 1]
  }
  t.values[0] = val
}
