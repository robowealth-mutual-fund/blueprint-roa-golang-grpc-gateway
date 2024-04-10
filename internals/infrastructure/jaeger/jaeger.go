package jaeger

import (
	"github.com/opentracing/opentracing-go"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/config"
	log "github.com/sirupsen/logrus"
	jaegerConf "github.com/uber/jaeger-client-go/config"
	jaegerLog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-client-go/rpcmetrics"
	"github.com/uber/jaeger-lib/metrics"
	"io"
)

func NewJaeger(appConfig config.Configuration) io.Closer {
	cfg, err := jaegerConf.FromEnv()
	panicIfErr(err)
	cfg.ServiceName = appConfig.AppName + "-" + appConfig.Env
	cfg.Sampler.Type = "const"
	cfg.Sampler.Param = 1
	cfg.Reporter = &jaegerConf.ReporterConfig{
		LogSpans:           true,
		LocalAgentHostPort: appConfig.Jaeger.JaegerAgentHost + ":" + appConfig.Jaeger.JaegerAgentPort,
	}

	jLogger := jaegerLog.StdLogger
	jMetricsFactory := metrics.NullFactory
	jMetricsFactory = jMetricsFactory.Namespace(metrics.NSOptions{Name: appConfig.AppName + "-" + appConfig.Env, Tags: nil})

	tracer, closer, err := cfg.NewTracer(
		jaegerConf.Logger(jLogger),
		jaegerConf.Metrics(jMetricsFactory),
		jaegerConf.Observer(rpcmetrics.NewObserver(jMetricsFactory, rpcmetrics.DefaultNameNormalizer)),
	)
	panicIfErr(err)
	opentracing.SetGlobalTracer(tracer)

	return closer
}

func panicIfErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
