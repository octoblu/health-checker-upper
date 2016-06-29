package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/codegangsta/cli"
	"github.com/coreos/go-semver/semver"
	"github.com/fatih/color"
	"github.com/octoblu/health-checker-upper/health"
	"github.com/octoblu/health-checker-upper/vulcand"
	De "github.com/tj/go-debug"
)

var debug = De.Debug("health-checker-upper:main")

func main() {
	app := cli.NewApp()
	app.Name = "health-checker-upper"
	app.Version = version()
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "vulcan-uri, u",
			EnvVar: "HEALTH_CHECKER_UPPER_VULCAN_URI",
			Usage:  "URI to vulcand's API endpoint (ex: http://127.0.0.1:8182)",
		},
	}
	app.Run(os.Args)
}

func fatalIfError(msg string, err error) {
	if err == nil {
		return
	}

	log.Fatalln(msg, err.Error())
}

func run(context *cli.Context) {
	vulcanURI := getOpts(context)

	stopSignal := make(chan os.Signal)
	signal.Notify(stopSignal, syscall.SIGTERM)
	signal.Notify(stopSignal, syscall.SIGINT)

	stopSignalReceived := false

	go func() {
		<-stopSignal
		fmt.Println("Stop Signal received, waiting to exit")
		stopSignalReceived = true
	}()

	for {
		debug("")
		if stopSignalReceived {
			fmt.Println("I'll be back.")
			os.Exit(0)
		}

		manager, err := vulcand.NewManager(vulcanURI)
		fatalIfError("error on vulcand.NewManager", err)
		servers, err := manager.ShuffledServers()
		fatalIfError("error on manager.ShuffledServers", err)

		for _, server := range servers {
			ok := health.Check(server)
			debug("server: %v (ok: %v)", server.ServerID(), ok)
			if !ok {
				fmt.Printf("Bad Server Found: {name: %v, url: %v}\n", server.ServerID(), server.URL())
				err := manager.ServerRm(server)
				fatalIfError("error on manager.ServerRm", err)
			}
		}
	}
}

func getOpts(context *cli.Context) string {
	vulcanURI := context.String("vulcan-uri")

	if vulcanURI == "" {
		cli.ShowAppHelp(context)

		if vulcanURI == "" {
			color.Red("  Missing required flag --vulcan-uri or HEALTH_CHECKER_UPPER_VULCAN_URI")
		}
		os.Exit(1)
	}

	return vulcanURI
}

func version() string {
	version, err := semver.NewVersion(VERSION)
	if err != nil {
		errorMessage := fmt.Sprintf("Error with version number: %v", VERSION)
		log.Panicln(errorMessage, err.Error())
	}
	return version.String()
}
