// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"context"
	"fmt"
	"github.com/facebookincubator/symphony/graph/event"
	"github.com/facebookincubator/symphony/graph/graphgrpc"
	"github.com/facebookincubator/symphony/graph/graphhttp"
	"github.com/facebookincubator/symphony/graph/viewer"
	"github.com/facebookincubator/symphony/pkg/log"
	"github.com/facebookincubator/symphony/pkg/mysql"
	"github.com/facebookincubator/symphony/pkg/server"
	"gocloud.dev/server/health"
	"google.golang.org/grpc"
	"net/url"
)

import (
	_ "github.com/facebookincubator/symphony/graph/ent/runtime"
	_ "github.com/go-sql-driver/mysql"
	_ "gocloud.dev/pubsub/mempubsub"
	_ "gocloud.dev/pubsub/natspubsub"
)

// Injectors from wire.go:

func NewApplication(ctx context.Context, flags *cliFlags) (*application, func(), error) {
	config := flags.Log
	logger, cleanup, err := log.Provider(config)
	if err != nil {
		return nil, nil, err
	}
	string2 := flags.MySQL
	mySQLTenancy, err := newMySQLTenancy(string2, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	eventConfig := flags.Event
	topicEmitter, cleanup2, err := event.ProvideEmitter(ctx, eventConfig)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	tenancy, err := newTenancy(mySQLTenancy, logger, topicEmitter)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	url, err := newAuthURL(flags)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	urlSubscriber := event.ProvideSubscriber(eventConfig)
	options := flags.Census
	v := newHealthChecks(mySQLTenancy)
	orc8rConfig := flags.Orc8r
	graphhttpConfig := graphhttp.Config{
		Tenancy:      tenancy,
		AuthURL:      url,
		Subscriber:   urlSubscriber,
		Logger:       logger,
		Census:       options,
		HealthChecks: v,
		Orc8r:        orc8rConfig,
	}
	server, cleanup3, err := graphhttp.NewServer(graphhttpConfig)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	db := mysql.Open(string2)
	graphgrpcConfig := graphgrpc.Config{
		DB:      db,
		Logger:  logger,
		Orc8r:   orc8rConfig,
		Tenancy: tenancy,
	}
	grpcServer, cleanup4, err := graphgrpc.NewServer(graphgrpcConfig)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	mainApplication := newApplication(logger, server, grpcServer, flags)
	return mainApplication, func() {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

func newApplication(logger log.Logger, httpServer *server.Server, grpcServer *grpc.Server, flags *cliFlags) *application {
	var app application
	app.Logger = logger.Background()
	app.http.Server = httpServer
	app.http.addr = flags.HTTPAddress
	app.grpc.Server = grpcServer
	app.grpc.addr = flags.GRPCAddress
	return &app
}

func newTenancy(tenancy *viewer.MySQLTenancy, logger log.Logger, emitter event.Emitter) (viewer.Tenancy, error) {
	eventer := event.Eventer{Logger: logger, Emitter: emitter}
	return viewer.NewCacheTenancy(tenancy, eventer.HookTo), nil
}

func newHealthChecks(tenancy *viewer.MySQLTenancy) []health.Checker {
	return []health.Checker{tenancy}
}

func newMySQLTenancy(dsn string, logger log.Logger) (*viewer.MySQLTenancy, error) {
	tenancy, err := viewer.NewMySQLTenancy(dsn)
	if err != nil {
		return nil, fmt.Errorf("creating mysql tenancy: %w", err)
	}
	tenancy.SetLogger(logger)
	mysql.SetLogger(logger)
	return tenancy, nil
}

func newAuthURL(flags *cliFlags) (*url.URL, error) {
	u, err := url.Parse(flags.AuthURL)
	if err != nil {
		return nil, fmt.Errorf("parsing auth url: %w", err)
	}
	return u, nil
}
