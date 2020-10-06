package txcrafter

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault(MetricsURLViperKey, metricsURLDefault)
	_ = viper.BindEnv(MetricsURLViperKey, metricsURLEnv)
}

const (
	MetricsURLViperKey = "tx-crafter.metrics.url"
	metricsURLDefault  = "localhost:8082"
	metricsURLEnv      = "TX_CRAFTER_METRICS_URL"
)