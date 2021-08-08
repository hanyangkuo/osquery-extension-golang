package main

import (
	"testing"
)

func TestFindScript(t *testing.T) {
	filenames, err := FindScript(`.\script`)
	if err != nil {
		t.Fail()
	}
	if len(filenames) < 1 {
		t.Fail()
	}
}

func TestExecuteScript(t *testing.T) {
	_, err := ExecuteScript(`.\script`, `GetLocalUser.ps1`)
	if err != nil {
		t.Fail()
	}
}
