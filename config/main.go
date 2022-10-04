package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go.temporal.io/server/common/cluster"
	"go.temporal.io/server/common/config"
	"go.temporal.io/server/common/log"
	"go.temporal.io/server/common/metrics"
	"go.temporal.io/server/temporal"
	"strings"
	"time"
)

func InitConfig() (*config.Config, []string) {
	viper.SetEnvPrefix("temporal")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	SetDefaults()
	viper.AutomaticEnv()

	services := strings.Split(viper.GetString("services"), ",")

	cfg := &config.Config{
		Global: config.Global{
			Membership: config.Membership{
				MaxJoinDuration:  viper.GetDuration("global.membership.maxjoinduration"),
				BroadcastAddress: viper.GetString("global.membership.broadcastaddress"),
			},
			Metrics: MetricsConfig(services),
		},
		Log: log.Config{
			Stdout: viper.GetBool("log.stdout"),
			Level:  viper.GetString("log.level"),
		},
		Persistence: config.Persistence{
			DefaultStore:     "default",
			VisibilityStore:  "visibility",
			NumHistoryShards: viper.GetInt32("persistence.numhistoryshards"),
			DataStores:       DataStoreConfig(),
		},
		PublicClient: config.PublicClient{
			HostPort: viper.GetString("publicclient.host") + ":" + viper.GetString("publicclient.port"),
		},
		Services: ServiceConfig(),
		ClusterMetadata: &cluster.Config{
			CurrentClusterName:       "active",
			MasterClusterName:        "active",
			FailoverVersionIncrement: 10,
			ClusterInformation: map[string]cluster.ClusterInformation{
				"active": {
					Enabled:                true,
					InitialFailoverVersion: 1,
					RPCAddress:             viper.GetString("publicclient.host") + ":" + viper.GetString("publicclient.port"),
				},
			},
		},
	}

	return cfg, services
}

func MetricsConfig(services []string) *metrics.Config {
	tags := make(map[string]string)
	if len(services) > 1 {
		tags["service"] = strings.Join(services, ",")
	} else {
		tags["service"] = services[0]
	}

	return &metrics.Config{
		ClientConfig: metrics.ClientConfig{
			Tags: tags,
		},
		Prometheus: &metrics.PrometheusConfig{
			Framework:     viper.GetString("metrics.prometheus.framework"),
			ListenAddress: viper.GetString("metrics.prometheus.listenaddress") + ":" + viper.GetString("metrics.prometheus.listenport"),
			HandlerPath:   viper.GetString("metrics.prometheus.handlerpath"),
		},
	}
}

func DataStoreConfig() map[string]config.DataStore {
	cfg := make(map[string]config.DataStore)

	cfg["default"] = config.DataStore{
		SQL: &config.SQL{
			User:            viper.GetString("sql.user"),
			Password:        viper.GetString("sql.password"),
			PluginName:      viper.GetString("sql.plugin"),
			DatabaseName:    viper.GetString("sql.database"),
			ConnectAddr:     viper.GetString("sql.host") + ":" + viper.GetString("sql.port"),
			ConnectProtocol: "tcp",
			MaxConns:        viper.GetInt("sql.maxconns"),
			MaxIdleConns:    viper.GetInt("sql.maxidleconns"),
			MaxConnLifetime: viper.GetDuration("sql.maxconnlifetime"),
		},
	}

	cfg["visibility"] = config.DataStore{
		SQL: &config.SQL{
			User:            viper.GetString("visibility.sql.user"),
			Password:        viper.GetString("visibility.sql.password"),
			PluginName:      viper.GetString("visibility.sql.plugin"),
			DatabaseName:    viper.GetString("visibility.database"),
			ConnectAddr:     viper.GetString("visibility.sql.host") + ":" + viper.GetString("visibility.sql.port"),
			ConnectProtocol: "tcp",
			MaxConns:        viper.GetInt("visibility.sql.maxconns"),
			MaxIdleConns:    viper.GetInt("visibility.sql.maxidleconns"),
			MaxConnLifetime: viper.GetDuration("visibility.sql.maxconnlifetime"),
		},
	}

	return cfg
}

func ServiceConfig() map[string]config.Service {
	m := make(map[string]config.Service)
	for _, service := range temporal.Services {
		grpcKey := fmt.Sprintf("service.%s.grpcport", service)
		membershipKey := fmt.Sprintf("service.%s.membershipport", service)
		ipKey := fmt.Sprintf("service.%s.bindonip", service)
		m[service] = config.Service{
			RPC: config.RPC{
				//GRPCPort:       viper.GetInt(fmt.Sprintf("service.%s.grpcport", service)),
				GRPCPort:       viper.GetInt(grpcKey),
				MembershipPort: viper.GetInt(membershipKey),
				BindOnIP:       viper.GetString(ipKey),
			},
		}
	}

	return m
}

func SetDefaults() {
	viper.SetDefault("services", "frontend,history,matching,worker")

	viper.SetDefault("global.membership.maxjoinduration", 30*time.Second)
	viper.SetDefault("global.membership.broadcastaddress", "0.0.0.0")

	viper.SetDefault("metrics.prometheus.framework", "opentelemetry")
	viper.SetDefault("metrics.prometheus.listenaddress", "0.0.0.0")
	viper.SetDefault("metrics.prometheus.listenport", "9090")
	viper.SetDefault("metrics.prometheus.handlerpath", "/metrics")

	viper.SetDefault("log.stdout", true)
	viper.SetDefault("log.level", "info")

	viper.SetDefault("persistence.numhistoryshards", 512)

	viper.SetDefault("sql.user", "temporal")
	viper.SetDefault("sql.password", "password")
	viper.SetDefault("sql.plugin", "postgres")
	viper.SetDefault("sql.database", "temporal")
	viper.SetDefault("sql.host", "localhost")
	viper.SetDefault("sql.port", "5432")
	viper.SetDefault("sql.maxconns", 20)
	viper.SetDefault("sql.maxidleconns", 20)
	viper.SetDefault("sql.maxconnlifetime", 1*time.Hour)

	viper.SetDefault("visibility.sql.user", "temporal_visibility")
	viper.SetDefault("visibility.sql.password", "password")
	viper.SetDefault("visibility.sql.plugin", "postgres")
	viper.SetDefault("visibility.sql.database", "temporal_visibility")
	viper.SetDefault("visibility.sql.host", "localhost")
	viper.SetDefault("visibility.sql.port", "5432")
	viper.SetDefault("visibility.sql.maxconns", 20)
	viper.SetDefault("visibility.sql.maxidleconns", 20)
	viper.SetDefault("visibility.sql.maxconnlifetime", 1*time.Hour)

	viper.SetDefault("publicclient.host", "0.0.0.0")
	viper.SetDefault("publicclient.port", "7233")

	viper.SetDefault("service.frontend.grpcport", 7233)
	viper.SetDefault("service.frontend.membershipport", 6933)
	viper.SetDefault("service.frontend.bindonip", "0.0.0.0")

	viper.SetDefault("service.history.grpcport", 7234)
	viper.SetDefault("service.history.membershipport", 6934)
	viper.SetDefault("service.history.bindonip", "0.0.0.0")

	viper.SetDefault("service.matching.grpcport", 7235)
	viper.SetDefault("service.matching.membershipport", 6935)
	viper.SetDefault("service.matching.bindonip", "0.0.0.0")

	viper.SetDefault("service.worker.grpcport", 7239)
	viper.SetDefault("service.worker.membershipport", 6939)
	viper.SetDefault("service.worker.bindonip", "0.0.0.0")
}
