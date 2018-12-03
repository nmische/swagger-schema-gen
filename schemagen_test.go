package main

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestEndToEnd(t *testing.T) {

	dir, schemagen := buildSchemagen(t)
	defer os.RemoveAll(dir)

	// Read the testdata directory.
	fd, err := os.Open("testdata")
	if err != nil {
		t.Fatal(err)
	}
	defer fd.Close()

	names, err := fd.Readdirnames(-1)
	if err != nil {
		t.Fatalf("Readdirnames: %s", err)
	}
	// Run schemagen against the test files.
	for _, name := range names {
		runSchemagen(t, dir, schemagen, name)
	}
}

// buildSchemagen creates a temporary directory and installs schemagen there.
func buildSchemagen(t *testing.T) (dir string, schemagen string) {
	t.Helper()
	dir, err := ioutil.TempDir("", "schemagen")
	if err != nil {
		t.Fatal(err)
	}
	schemagen = filepath.Join(dir, "schemagen")
	err = run("go", "build", "-o", schemagen)
	if err != nil {
		t.Fatalf("building schemagen: %s", err)
	}
	return dir, schemagen
}

// runSchemagen runs schemagen.
func runSchemagen(t *testing.T, dir, schemagen, fileName string) {
	t.Helper()

	subDir := strings.TrimSuffix(filepath.Base(fileName), ".go")
	pkgDir := filepath.Join(dir, subDir)

	t.Logf("run: %s %s\n", schemagen, pkgDir)

	err := os.MkdirAll(pkgDir, 0744)
	if err != nil {
		t.Fatalf("creating target dir: %s", err)
	}
	source := filepath.Join(pkgDir, fileName)
	err = copy(source, filepath.Join("testdata", fileName))
	if err != nil {
		t.Fatalf("copying file to temporary directory: %s", err)
	}
	// Run schemagen in temporary directory.
	err = run(schemagen, pkgDir)
	if err != nil {
		t.Fatal(err)
	}

}

// copy copies the from file to the to file.
func copy(to, from string) error {
	toFd, err := os.Create(to)
	if err != nil {
		return err
	}
	defer toFd.Close()
	fromFd, err := os.Open(from)
	if err != nil {
		return err
	}
	defer fromFd.Close()
	_, err = io.Copy(toFd, fromFd)
	return err
}

// run runs a single command and returns an error if it does not succeed.
// os/exec should have this function, to be honest.
func run(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
