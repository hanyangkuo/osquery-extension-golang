package schema

import (
	"context"
	"encoding/json"
	"github.com/osquery/osquery-go/plugin/table"
	"log"
	"testing"
)

func TestRegistryColumns(t *testing.T) {
	tableColumns := RegistryColumns()
	if len(tableColumns) < 1{
		t.Fail()
	}
	t.Log(tableColumns)
}
func TestRegistryGenerate(t *testing.T) {
	results, err := RegistryGenerate(context.Background(), table.QueryContext{})
	if err != nil {
		t.Fail()
	}
	jsonResults, err := json.MarshalIndent(results, "", "\t")
	log.Println(string(jsonResults))
}