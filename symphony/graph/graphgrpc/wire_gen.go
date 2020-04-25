// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package graphgrpc

import (
	"database/sql"
	"github.com/facebookincubator/symphony/graph/viewer"
	"github.com/facebookincubator/symphony/pkg/actions/action/magmarebootnode"
	"github.com/facebookincubator/symphony/pkg/actions/executor"
	"github.com/facebookincubator/symphony/pkg/actions/trigger/magmaalert"
	"github.com/facebookincubator/symphony/pkg/log"
	"github.com/facebookincubator/symphony/pkg/orc8r"
	"google.golang.org/grpc"
	"net/http"
)

// Injectors from wire.go:

func NewServer(cfg Config) (*grpc.Server, func(), error) {
	tenancy := cfg.Tenancy
	db := cfg.DB
	logger := cfg.Logger
	config := cfg.Orc8r
	client := newOrc8rClient(config)
	registry := newActionsRegistry(client)
	server, cleanup, err := newServer(tenancy, db, logger, registry)
	if err != nil {
		return nil, nil, err
	}
	return server, func() {
		cleanup()
	}, nil
}

// wire.go:

// Config defines the grpc server config.
type Config struct {
	DB      *sql.DB
	Logger  log.Logger
	Orc8r   orc8r.Config
	Tenancy viewer.Tenancy
}

func newOrc8rClient(config orc8r.Config) *http.Client {
	client, _ := orc8r.NewClient(config)
	return client
}

func newActionsRegistry(orc8rClient *http.Client) *executor.Registry {
	registry := executor.NewRegistry()
	registry.MustRegisterTrigger(magmaalert.New())
	registry.MustRegisterAction(magmarebootnode.New(orc8rClient))
	return registry
}
