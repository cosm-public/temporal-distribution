package main

import (
	"github.com/cosm-eng/temporal/distribution/config"
	"go.temporal.io/server/common/build"
	"go.temporal.io/server/common/headers"
	"go.temporal.io/server/common/log"
	"go.temporal.io/server/common/log/tag"
	"go.temporal.io/server/temporal"
	stdlog "log"

	_ "go.temporal.io/server/common/persistence/sql/sqlplugin/postgresql" // needed to load postgresql plugin
)

func main() {
	cfg, services := config.InitConfig()

	logger := log.NewZapLogger(log.BuildZapLogger(cfg.Log))
	logger.Info("Build info.",
		tag.NewTimeTag("git-time", build.InfoData.GitTime),
		tag.NewStringTag("git-revision", build.InfoData.GitRevision),
		tag.NewBoolTag("git-modified", build.InfoData.GitModified),
		tag.NewStringTag("go-arch", build.InfoData.GoArch),
		tag.NewStringTag("go-os", build.InfoData.GoOs),
		tag.NewStringTag("go-version", build.InfoData.GoVersion),
		tag.NewBoolTag("cgo-enabled", build.InfoData.CgoEnabled),
		tag.NewStringTag("server-version", headers.ServerVersion),
	)

	err := cfg.Validate()
	if err != nil {
		stdlog.Fatal(err)
	}

	s := temporal.NewServer(
		temporal.ForServices(services),
		temporal.WithConfig(cfg),
		temporal.WithLogger(logger),
		temporal.InterruptOn(temporal.InterruptCh()),
	)
	err = s.Start()
	if err != nil {
		stdlog.Fatal(err)
	}
}
