package generator

import (
	"io"
	"os"
	"log"
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
	Inputs []LogPipe
}

type LogPipe struct {
	Region string
}

func (config *CircleciConfigGenerator) AddInput(regions string) CircleciConfig {
	config.Inputs = append(config.Inputs, LogPipe{
		Region: regions,
	})

	return config
}

func (config *CircleciConfigGenerator) WriteCircleciConfig(wr io.Writer) error {
	f, err := os.Create("./config.yaml")
	if err != nil {
		log.Println("create file: ", err)
	}
	
	tmpl, err := template.New("config").Parse(circleciConfigTemplate)
	if err != nil {
		return err
	}
	return tmpl.Execute(f, config)
}
