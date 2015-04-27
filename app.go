/**
 * Copyright 2014 @ S1N1 Team.
 * name : app1.go
 * author : newmin
 * date : 2015-04-27 20:43:
 * description :
 * history :
 */
package gof

import (
	"github.com/atnet/gof/db"
	"github.com/atnet/gof/log"
)

type App interface {
	// Provided db access
	Db() db.Connector

	// Return a Wrapper for golang template.
	Template() *Template

	// Return application configs.
	Config() *Config

	// Return a logger
	Log() log.ILogger

	// Get a reference of AppContext
	Source() interface{}

	// Application is running debug mode
	Debug() bool
}