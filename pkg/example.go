package example

import (
	logging "github.com/zaphiro-technologies/logging/pkg"
)

func Hello() string {
	return "Hello, World!"
}

func HelloWithLog() string {
	log := logging.SetUp("hello", &logging.LoggerFields{
		Tenant:          "tenant1",
		NetworkId:       "my-network",
		CorrelationTime: "1231236274",
		DeviceId:        "pmu-1",
	})

	log.Info().
		Msg("ok")
	return Hello()
}
