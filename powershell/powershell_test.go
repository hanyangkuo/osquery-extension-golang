package powershell

import (
	"context"
	"encoding/json"
	"github.com/osquery/osquery-go/plugin/table"
	"log"
	"testing"
)

func TestPScriptColumns(t *testing.T) {
	root = `..`
	tableColumns := PScriptColumns()
	if len(tableColumns) < 1{
		t.Fail()
	}
	t.Log(tableColumns)
}

func TestPScriptGenerate(t *testing.T) {
	root = `..`
	results, err := PScriptGenerate(context.Background(), table.QueryContext{})
	if err != nil {
		t.Fail()
	}
	jsonResults, err := json.MarshalIndent(results, "", "\t")
	log.Println(string(jsonResults))
}

func TestPScriptGenerateWithCombine(t *testing.T) {
	root = `..`
	results, err := PScriptGenerateWithCombine(context.Background(), table.QueryContext{})
	if err != nil {
		t.Fail()
	}
	jsonResults, err := json.MarshalIndent(results, "", "\t")
	log.Println(string(jsonResults))
}
