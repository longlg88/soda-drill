package generator

var circleciConfigTemplate = `version: 2
jobs:
  {{- range .Inputs}}
  bake-image-{{ .region }}
    <<: *default-bake-job
	{{- end}}`
