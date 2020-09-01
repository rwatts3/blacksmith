package blacksmith

import (
	"github.com/nunchistudio/blacksmith/adapter/pubsub"
	"github.com/nunchistudio/blacksmith/adapter/store"
	"github.com/nunchistudio/blacksmith/adapter/supervisor"
	"github.com/nunchistudio/blacksmith/adapter/wanderer"
	"github.com/nunchistudio/blacksmith/flow/destination"
	"github.com/nunchistudio/blacksmith/flow/source"
	"github.com/nunchistudio/blacksmith/service"

	"github.com/sirupsen/logrus"
)

/*
Options is the options a user can pass to create a new application.
*/
type Options struct {

	// Logger allows you to use a logrus Logger across all Blacksmith adapters and
	// the application built on top of it.
	Logger *logrus.Logger

	// Supervisor is the options passed to use the supervisor adapter.
	// The supervisor is optional.
	Supervisor *supervisor.Options

	// Wanderer is the options passed to use the wanderer adapter.
	// The wanderer is optional.
	Wanderer *wanderer.Options

	// Store is the options passed to use the store adapter.
	Store *store.Options

	// PubSub is the options passed to use the pubsub adapter.
	// The pusub is optional.
	PubSub *pubsub.Options

	// Gateway is the options passed to use the gateway service.
	Gateway *service.Options

	// Scheduler is the options passed to use the scheduler service.
	Scheduler *service.Options

	// Sources is a slice of options passed to create sources.
	Sources []*source.Options

	// Destinations is a slice of options passed to create destinations.
	Destinations []*destination.Options
}
