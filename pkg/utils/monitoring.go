package utils

import (
	"github.com/newrelic/go-agent/v3/newrelic"
)

// SetupNewRelic setup newrelic for apm
func SetupNewRelic() (*newrelic.Application, error) {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("GoGinApi"),
		newrelic.ConfigLicense("eu01xxf5377b1860a6a52f52c9c20c86adadNRAL"),
		newrelic.ConfigDistributedTracerEnabled(true),
	)
	if err != nil {
		return nil, err
	}
	return app, nil
}
