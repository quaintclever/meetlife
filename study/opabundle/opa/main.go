package main

import (
	"context"
	"flag"
	"fmt"
	"opa/internal/api"
	"opa/internal/opa"
	"opa/internal/version"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/open-policy-agent/opa/rego"
)

func main() {
	// 单机 opa test
	//simpleTest()

	// bundle opa test
	bundlesTest()
}

var configFile = flag.String("config", "", "set the OPA config file to load")
var verbose = flag.Bool("verbose", false, "enable verbose logging")
var versionFlag = flag.Bool("version", false, "print version and exit")

func bundlesTest() {
	flag.Parse()

	if *versionFlag {
		fmt.Println("Version:", version.Version)
		fmt.Println("Vcs:", version.Vcs)
		os.Exit(0)
	}
	setupLogging()

	engine, err := opa.New(opa.Config(*configFile))
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Fatal("Failed to initialize OPA.")
	}

	ctx := context.Background()

	if err := engine.Start(ctx); err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Fatal("Failed to start OPA.")
	}

	if err := api.New(engine).Run(ctx); err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("Shutting down.")

}

func setupLogging() {
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	logLevel := logrus.InfoLevel
	if *verbose {
		logLevel = logrus.DebugLevel
	}
	logrus.SetLevel(logLevel)
}

func simpleTest() {

	module := `
package example.authz

default allow = false

allow {
    some id
    input.method = "GET"
    input.path = ["salary", id]
    input.subject.user = id
}

allow {
    is_admin
}

is_admin {
    input.subject.groups[_] = "admin"
}
`
	ctx := context.TODO()
	query, err := rego.New(
		rego.Query("x = data.example.authz.allow"),
		rego.Module("example.rego", module),
	).PrepareForEval(ctx)

	if err != nil {
		// Handle error.
	}

	input := map[string]interface{}{
		"method": "GET",
		"path":   []interface{}{"salary", "bob"},
		"subject": map[string]interface{}{
			"user":   "bob",
			"groups": []interface{}{"sales", "marketing"},
		},
	}

	results, err := query.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		// handle error
	}
	if !results.Allowed() {
		// handle result
		fmt.Println("allowed!")
	}
	fmt.Println()
}
