# temporal-distribution
![goreleaser](https://github.com/cosm-eng/temporal-distribution/actions/workflows/goreleaser.yaml/badge.svg)

*NOTE: This should be considered an experimental project.*

This is [Cosm](https://cosm.com)'s modification of [Temporal](https://temporal.io), a workflow orchestrator.

### Differences
This distribution uses `go.temporal.io/server/temporal` as a [Go module](https://docs.temporal.io/references/server-options),
and injects configuration through environment variables (via [Viper](https://github.com/spf13/viper)), instead of through 
configuration file.

The target platform is Kubernetes for Temporal services and Postgres for default and visibility datastores 
(advanced visibility not yet supported), this repository hosts basic manifests in the [kubernetes](kubernetes/) directory.

Environment variables use the pattern `TEMPORAL_$SECTION_$ITEM`, like `TEMPORAL_PUBLICCLIENT_HOST`. 
See [config](config/main.go) for more options.

### Usage
Ensure `kubectl` and `kustomize` are available on your machine.

You will also need the `tctl` and `temporal-sql-tool` binaries on your path:
* Clone [https://github.com/temporalio/temporal](https://github.com/temporalio/temporal)
  * In this directory:
    * `make update-tctl`
    * `make temporal-sql-tool` and `ln -s ./temporal-sql-tool $GOBIN/temporal-sql-tool` or move the compiled binary to your $PATH.
* Create the necessary Postgres databases and users (see [postgres.yaml](kubernetes/postgres.yaml) for an example script) and set
the variables required in [dev-migrations.sh](dev-migrations.sh).
* Run migrations
* Deploy to Kubernetes
  * `kubectl apply -f kubernetes/namespace.yaml`
  * `kustomize build kubernetes/ | kubectl apply -f -`
* Create Temporal namespaces (`tctl namespace register` for `default`)

### Packaging
This distribution is built via `goreleaser` and GitHub Actions when a new tag is created. 
This tag adheres to the following conventions:
* `$TEMPORAL_VERSION-$DISTRIBUTION_BUILD_NUMBER`, for example: `v1.18.0-0`
* Only code changes should create a new tag. Changes to manifests do not require a rebuild of the distribution.

To upgrade, update the `go.temporal.io/server` version in `go.mod` and run `go mod tidy` to refresh the lockfile. 
Then push to main, and create a new tag like above, and push that.
On tag, GitHub Actions will build and publish artifacts described in [.goreleaser.yaml](.goreleaser.yaml):
* Server binaries, see [Releases](https://github.com/cosm-eng/temporal-distribution/releases)
* Multi-arch container, see [Packages](https://github.com/orgs/cosm-eng/packages?repo_name=temporal-distribution)

### Developing
In general, local development uses the [Makefile](Makefile) and standard `go` tooling.
You will need a Postgres instance with the databases, users and migrations described above.