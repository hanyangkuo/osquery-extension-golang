package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/osquery/osquery-go"
	"github.com/osquery/osquery-go/plugin/table"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

var (
	socket   = flag.String("socket", "", "Path to the extensions UNIX domain socket")
	timeout  = flag.Int("timeout", 3, "Seconds to wait for autoloaded extensions")
	interval = flag.Int("interval", 3, "Seconds delay between connectivity checks")
	verbose = flag.Bool("verbose", false, "Seconds delay between connectivity checks")
	root = `C:\Program Files\osquery`
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

	server, err := osquery.NewExtensionManagerServer(
		"script_example",
		*socket,
		serverTimeout,
		serverPingInterval,
	)

	if err != nil {
		log.Fatalf("Error creating extension: %s\n", err)
	}
	server.RegisterPlugin(table.NewPlugin("script_example", PScriptColumns(), PScriptGenerate))
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
// PScriptColumns returns the columns that our table will return.
func PScriptColumns() []table.ColumnDefinition {
	return []table.ColumnDefinition{
		table.TextColumn("script"),
		table.TextColumn("result"),
	}
}


func PScriptGenerate(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	filenames, err := FindScript(filepath.Join(root, "script"))
	if err != nil {
		return nil, fmt.Errorf("Script not found %v\n", err)
	}
	var results []map[string]string
	for _, filename := range filenames {
		stdOut, err := ExecuteScript(filepath.Join(root, "script"), filename)
		if err != nil{
			results = append(results, map[string]string{
				"script": filename,
				"result": fmt.Sprintf("%v", err),
			})
			continue
		}
		results = append(results, map[string]string{
			"script": filename,
			"result": stdOut,
		})
	}
	return results, nil
}


func FindScript(dir string) ([]string,error){
	var filenames []string
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files{
		if !file.IsDir(){
			filenames = append(filenames, file.Name())
		}
	}
	return filenames, nil
}

func ExecuteScript(dir string, filename string) (string, error){
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()cmd.exe
	cmd := exec.Command(`C:\Windows\System32\cmd.exe`, "/c", fmt.Sprintf(`chcp 65001&powershell .\%s &exit`, filename))
	cmd.Dir = dir
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	stdoutStderr, err := cmd.CombinedOutput()
	cmd = nil
	if err != nil {
		return "", err
	}
	return strings.TrimPrefix(string(stdoutStderr), "Active code page: 65001\r\n"), nil
}