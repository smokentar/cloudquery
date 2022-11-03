// Code generated by codegen; DO NOT EDIT.
package client

import (
    "context"
    "github.com/aws/aws-sdk-go-v2/aws"
{{- range $service := . }}
    "{{ $service.Import }}"
{{- end }}
)

func initServices(region string, c aws.Config) Services {
	awsCfg := c.Copy()
	awsCfg.Region = region
	return Services{
	{{- range $service := . }}
		{{$service.Name}}: {{$service.PackageName}}.NewFromConfig(awsCfg),
    {{- end }}
	}
}

type Services struct {
	{{- range $service := . }}
		{{$service.Name}} {{$service.ClientName}}
    {{- end }}
}

{{- range $service := . }}
//go:generate mockgen -package=mocks -destination=./mocks/{{$service.PackageName}}.go . {{$service.ClientName}}
type {{$service.ClientName}} interface {
    {{- range $sig := $service.Signatures }}
    {{ $sig }}
    {{- end }}
}
{{- end }}