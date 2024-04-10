package config

type Jaeger struct {
	JaegerAgentHost string `env:"JAEGER_HOST" envDefault:"localhost"`
	JaegerAgentPort string `env:"JAEGER_PORT" envDefault:"6831"`
}
