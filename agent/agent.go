package agent

type Agent struct {
  Name string
  Version string
  Platform string
}

func Parse(b []byte) *Agent {
  a := new(Agent)
  var word []byte

  for c, i := range b {
  }

  return a
}
