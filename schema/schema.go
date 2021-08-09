package schema

import (
	"context"
	"fmt"
	"github.com/osquery/osquery-go/plugin/table"
	"golang.org/x/sys/windows/registry"
)


// RegistryColumns returns the columns that our table will return.
func RegistryColumns() []table.ColumnDefinition {
	return []table.ColumnDefinition{
		table.TextColumn("CurrentVersion"),
		table.TextColumn("ProductName"),
		table.TextColumn("CurrentMajorVersionNumber"),
		table.TextColumn("CurrentMinorVersionNumber"),
		table.TextColumn("CurrentBuild"),
	}
}

// RegistryGenerate generate data that our table will return.
func RegistryGenerate(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		return nil, err
	}
	defer k.Close()

	currentVersion, _, err := k.GetStringValue("CurrentVersion")
	if err != nil {
		return nil, err
	}

	productName , _, err := k.GetStringValue("ProductName")
	if err != nil {
		return nil, err
	}

	majorVersion, _, err := k.GetIntegerValue("CurrentMajorVersionNumber")
	if err != nil {
		return nil, err
	}

	minorVersion, _, err := k.GetIntegerValue("CurrentMinorVersionNumber")
	if err != nil {
		return nil, err
	}

	currentBuild, _, err := k.GetStringValue("CurrentBuild")
	if err != nil {
		return nil, err
	}
	return []map[string]string{
		{
			"CurrentVersion":             currentVersion,
			"ProductName":                productName,
			"CurrentMajorVersionNumber":  fmt.Sprintf(`%d`, majorVersion),
			"CurrentMinorVersionNumber":  fmt.Sprintf(`%d`,minorVersion),
			"CurrentBuild":               currentBuild,
		},
	}, nil
}
