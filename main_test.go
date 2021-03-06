package main

import (
	"janmarten.name/nv/debug"
	"os/exec"
	"regexp"
	"testing"
)

func TestMain_version(t *testing.T) {
	cmd := exec.Command(
		"go",
		"run",
		".",
		"--version",
	)

	got, err := cmd.Output()

	if err != nil {
		t.Fatalf("Could not run main program: %v", err)
	}

	if exit := cmd.ProcessState.ExitCode(); exit != 0 {
		t.Errorf("Program exited with non-zero code %q", exit)
	}

	want := regexp.MustCompile("(?s)^nv version dev-.+$")

	if !want.Match(got) {
		t.Errorf("Expected match for %q, got %q", want.String(), got)
	}
}

func TestMain_illegal(t *testing.T) {
	cmd := exec.Command(
		"go",
		"run",
		".",
		"--illegal",
	)

	_ = cmd.Run()

	if cmd.ProcessState.ExitCode() == 0 {
		t.Error("Program exited zero code, but expected non-zero code")
	}
}

func TestMain_init(t *testing.T) {
	got, ok := debug.Scope("nv.init.0").GetMessages()["Version"].(string)

	if !ok || got != version {
		t.Errorf("Unexpected version. Want %q, got %q", version, got)
	}
}
