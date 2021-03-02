package main

import (
	"bytes"
	"fmt"
	"soda-drill/src/generator"
)

func main() {
	var regions = []string{"accolade", "picpay"}

	config := generator.New()
	for _, s := range regions {
		config.AddInput(s)
	}
	fmt.Println()
	circleciConfig := new(bytes.Buffer)
	config.WriteCircleciConfig(circleciConfig)
	fmt.Println(circleciConfig.String())
}
