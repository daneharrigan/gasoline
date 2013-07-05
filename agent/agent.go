package agent

import (
	"bufio"
	"strings"
)

type Agent struct {
	Name     string
	Version  string
	Platform string
}

func Parse(str string) (a *Agent) {
	a = new(Agent)
	s := bufio.NewScanner(strings.NewReader(str))
	s.Split(bufio.ScanWords)

	for s.Scan() {
		val := s.Text()
		if strings.Contains(val, "/") {
			kv := strings.Split(val, "/")
			switch kv[0] {
			case "Version":
				a.Version = kv[1]
			case "Safari":
				a.Name = kv[0]
			case "Firefox":
				fallthrough
			case "Chrome":
				a.Name = kv[0]
				a.Version = kv[1]
				return
			}
		} else if val == "MSIE" {
			s.Scan()
			a.Name = val
			a.Version = strings.Replace(s.Text(), ";", "", 1)
			return
		}
	}

	return
}
