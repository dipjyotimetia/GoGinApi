package utils

import (
	"github.com/newrelic/go-agent/v3/newrelic"
	"os"
)

// SetupNewRelic setup newrelic for apm
func SetupNewRelic() (*newrelic.Application, error) {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("GoGinApi"),
		newrelic.ConfigLicense(os.Getenv("NR_APIKEY")),
		newrelic.ConfigDistributedTracerEnabled(true),
	)
	if err != nil {
		return nil, err
	}
	return app, nil
}
