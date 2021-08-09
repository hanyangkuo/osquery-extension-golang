package main

import (
	"flag"
	"github.com/osquery/osquery-go"
	"github.com/osquery/osquery-go/plugin/table"
	"log"
	"osquery-extension/powershell"
	"osquery-extension/schema"
	"time"
)

var (
	socket   = flag.String("socket", "", "Path to the extensions UNIX domain socket")
	timeout  = flag.Int("timeout", 3, "Seconds to wait for autoloaded extensions")
	interval = flag.Int("interval", 3, "Seconds delay between connectivity checks")
	verbose = flag.Bool("verbose", false, "Seconds delay between connectivity checks")

)

func main() {
	flag.Parse()
	if *socket == "" {
		log.Fatalln("Missing required --socket argument")
	}
	serverTimeout := osquery.ServerTimeout(
		time.Second * time.Duration(*timeout),
	)
	serverPingInterval := osquery.ServerPingInterval(
		time.Second * time.Duration(*interval),
	)

	time.Sleep(5*time.Second)

	// Create A NewExtensionMangerServer connect with osquery extensions UNIX domain socket.
	// serverTimeout: Timeout for autoloaded extensions server.
	// serverPingInterval: How often to ping osquery server, default: 5 seconds.
	server, err := osquery.NewExtensionManagerServer(
		"script_example",
		*socket,
		serverTimeout,
		serverPingInterval,
	)

	if err != nil {
		log.Fatalf("Error creating extension: %s\n", err)
	}

	// Create and register a new table plugin with the server.
	// table.NewPlugin requires the table plugin name, a slice of Columns and a Generate function.
	// <Note> Several plugins are allowed to register in the server.
	server.RegisterPlugin(
		table.NewPlugin("script_example", powershell.PScriptColumns(), powershell.PScriptGenerate),
		table.NewPlugin("registry_example", schema.RegistryColumns(), schema.RegistryGenerate),
	)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
