package utils

import (
	"github.com/newrelic/go-agent/v3/newrelic"
)

// SetupNewRelic setup newrelic for apm
func SetupNewRelic() (*newrelic.Application, error) {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("GoGinApi"),
		newrelic.ConfigLicense(""), // Enable with apikey
		newrelic.ConfigDistributedTracerEnabled(true),
	)
	if err != nil {
		return nil, err
	}
	return app, nil
}
