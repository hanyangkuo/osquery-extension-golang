package powershell

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/osquery/osquery-go/plugin/table"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

var (
	root = `C:\Program Files\osquery`
)

// PScriptColumns returns the columns that our table will return.
func PScriptColumns() []table.ColumnDefinition {
	return []table.ColumnDefinition{
		table.TextColumn("script"),
		table.TextColumn("result"),
	}
}

// PScriptGenerate generate data that our table will return.
func PScriptGenerate(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	filenames, err := findScript(filepath.Join(root, "scripts"))
	if err != nil {
		return nil, fmt.Errorf("Script not found %v\n", err)
	}
	var results []map[string]string
	for _, filename := range filenames {
		stdOut, err := executeScript(filepath.Join(root, "scripts"), filename)
		if err != nil{
			results = append(results, map[string]string{
				"script": filename,
				"result": "execute script error",
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

// PScriptGenerateWithCombine generate data that our table will return.
func PScriptGenerateWithCombine(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	filenames, err := findScript(filepath.Join(root, "scripts"))
	if err != nil {
		return nil, fmt.Errorf("Script not found %v\n", err)
	}
	var results []map[string]string
	for _, filename := range filenames {
		stdOut, err := executeScriptCombine(filepath.Join(root, "scripts"), filename)
		if err != nil{
			results = append(results, map[string]string{
				"scripts": filename,
				"result": "execute script error",
			})
			continue
		}
		results = append(results, map[string]string{
			"scripts": filename,
			"result": stdOut,
		})
	}
	return results, nil
}

func findScript(dir string) ([]string,error){
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

// executeScript used to execute ps1 file with powershell
func executeScript(dir string, filename string) (string, error){
	cmd := exec.Command(`C:\Windows\System32\cmd.exe`, "/c", fmt.Sprintf(`chcp 65001 && powershell .\%s && exit`, filename))
	cmd.Dir = dir
	var stdOut, stdErr bytes.Buffer

	// defer is used to ensure that a function call is performed later in a program’s execution
	// details: https://gobyexample.com/defer
	// claims stdOut -> stdErr, then close order stdErr -> stdOut
	defer stdOut.Reset()
	defer stdErr.Reset()
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := cmd.Run()
	cmd = nil
	if err != nil {
		return "", err
	}
	if stdErr.Len() > 0 {
		return "", errors.New(stdErr.String())
	}
	return strings.TrimPrefix(stdOut.String(), "Active code page: 65001\r\n"), nil
}

func executeScriptCombine(dir string, filename string) (string, error){
	cmd := exec.Command(`C:\Windows\System32\cmd.exe`, "/c", fmt.Sprintf(`chcp 65001 & powershell .\%s && exit`, filename))
	cmd.Dir = dir
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	stdoutStderr, err := cmd.CombinedOutput()
	cmd = nil
	if err != nil {
		return "", err
	}
	return strings.TrimPrefix(string(stdoutStderr), "Active code page: 65001\r\n"), nil
}
