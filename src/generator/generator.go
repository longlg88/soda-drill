package generator

import (
	"io"
	"text/template"
)

type CircleciConfig interface {
	AddInput(name string) CircleciConfig
	WriteCircleciConfig(wr io.Writer) error
}

func New() CircleciConfig {
	return &CircleciConfigGenerator{}
}

type CircleciConfigGenerator struct {
	Inputs []string
}

func (config *CircleciConfigGenerator) AddInput(regions string) CircleciConfig {
	config.Inputs = append(config.Inputs, regions)

	return config
}

func (config *CircleciConfigGenerator) WriteCircleciConfig(wr io.Writer) error {
	tmpl, err := template.New("config.yaml").Parse(circleciConfigTemplate)
	if err != nil {
		return err
	}
	return tmpl.Execute(wr, config)
}
