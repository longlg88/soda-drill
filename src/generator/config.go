package generator

var circleciConfigTemplate = `version: 2
jobs:
  {{- range .Inputs}}
  bake-image-{{ .Region }}
    <<: *default-bake-job
	{{ end -}}`
