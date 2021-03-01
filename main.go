package main

import (
	"bytes"
	"fmt"
	"github.com/longlg88/soda-drill"
)

func main() {
	var regions = []string{"accolade", "picpay"}

	config := New()
	for i, s := range regions {
		config.AddInput(s)
		fmt.Println(i)
	}
	circleciConfig := new(bytes.Buffer)
	config.WriteCircleciConfig(circleciConfig)
}
