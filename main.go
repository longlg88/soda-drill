package main

import (
	"bytes"
	"soda-drill/src/generator"
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
