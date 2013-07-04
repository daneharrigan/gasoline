package agent

import (
	"bufio"
	"strings"
	"fmt"
)

type Agent struct {
	Name string
	Version string
	Platform string
}

func Parse(s string) (a *Agent) {
	r := bufio.NewReader(strings.NewReader(s))

	for {
		b, err := r.ReadBytes(' ')
		if err != nil {
			fmt.Printf("error: %s\n", err)
			return
		}

		
	}

	return
}
