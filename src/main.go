package main

import (
	"bytes"
	"github.com/longlg88/soda-drill"
)

func main() {
	var regions = []string{"accolade", "picpay"}

	config := New()
	for i, s := range regions {
		config.AddInput(s)
		print(i)
	}
	circleciConfig := new(bytes.Buffer)
	config.WriteCircleciConfig(circleciConfig)
}
