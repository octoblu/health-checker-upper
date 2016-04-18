package registry

import (
	"github.com/vulcand/vulcand/plugin"

	"github.com/vulcand/vulcand/plugin/connlimit"

	"github.com/vulcand/vulcand/plugin/ratelimit"

	"github.com/vulcand/vulcand/plugin/rewrite"

	"github.com/vulcand/vulcand/plugin/cbreaker"

	"github.com/vulcand/vulcand/plugin/trace"

	"github.com/octoblu/vulcand-job-logger/joblogger"
)

func GetRegistry() (*plugin.Registry, error) {
	r := plugin.NewRegistry()

	specs := []*plugin.MiddlewareSpec{

		connlimit.GetSpec(),

		ratelimit.GetSpec(),

		rewrite.GetSpec(),

		cbreaker.GetSpec(),

		trace.GetSpec(),

		joblogger.GetSpec(),
	}

	for _, spec := range specs {
		if err := r.AddSpec(spec); err != nil {
			return nil, err
		}
	}
	return r, nil
}
