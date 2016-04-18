package joblogger

import (
	"github.com/codegangsta/cli"
	"github.com/vulcand/vulcand/plugin"
)

// Type is a short name for the middleware
const TYPE = "job-logger"

// GetSpec describes the job-logger middleware
func GetSpec() *plugin.MiddlewareSpec {
	return &plugin.MiddlewareSpec{
		Type:      TYPE,       // A short name for the middleware
		FromOther: FromOther,  // Tells vulcand how to create middleware from another one
		FromCli:   FromCli,    // Tells vulcand how to create middleware from CLI
		CliFlags:  CliFlags(), // Vulcand will add this flags CLI command
	}
}

// FromOther Will be called by Vulcand when engine or API will read the middleware from the serialized format.
// It's important that the signature of the function will be exactly the same, otherwise Vulcand will
// fail to register this middleware.
// The first and the only parameter should be the struct itself, no pointers and other variables.
// Function should return middleware interface and error in case if the parameters are wrong.
func FromOther(middleware Middleware) (plugin.Middleware, error) {
	return NewMiddleware(middleware.RedisURI, middleware.RedisQueueName, middleware.BackendID)
}

// FromCli constructs the middleware from the command line
func FromCli(context *cli.Context) (plugin.Middleware, error) {
	return NewMiddleware(context.String("redis-uri"), context.String("redis-queue-name"), context.String("backend"))
}

// CliFlags will be used by Vulcand construct help and CLI command for the vctl command
func CliFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "redis-uri",
			Usage:  "URI where redis can be reached",
			EnvVar: "LOGGER_REDIS_URI",
		},
		cli.StringFlag{
			Name:   "redis-queue-name",
			Usage:  "URI where redis can be reached",
			EnvVar: "LOGGER_REDIS_QUEUE_NAME",
		},
		cli.StringFlag{
			Name:   "backend",
			Usage:  "id of backend the frontend of this middleware is attached to",
			EnvVar: "LOGGER_BACKEND",
		},
	}
}
