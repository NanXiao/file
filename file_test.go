package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var (
	testPath    = "."
	testPrefix  = "example_"
	exampleFile = "exampleFile"
	exampleDir  = "exampleDir"
)

func cleanup() {
	os.Remove(exampleFile)
	os.RemoveAll(exampleDir)
}

func TestCopy(t *testing.T) {
	defer cleanup()

	testDir, err := ioutil.TempDir(testPath, testPrefix)
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(testDir)

	testDir2, err := ioutil.TempDir(testPath, testPrefix)
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(testDir2)

	testFile, err := ioutil.TempFile(testDir, testPrefix)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(testFile.Name())

	// test cases:
	// 1. file to file
	err = Copy(testFile.Name(), exampleFile)
	if err != nil {
		t.Errorf("1: %s\n", err)
	}

	if !Exists(exampleFile) {
		t.Error("1: exampleFile does not exist")
	}

	// 2. file into directory
	err = Copy(testFile.Name(), testDir2)
	if err != nil {
		t.Errorf("2: %s\n", err)
	}

	if !Exists(filepath.Join(testDir2, filepath.Base(testFile.Name()))) {
		t.Errorf("2: testFile does not exist in %s\n", testDir2)
	}

	// 3. file to an existing file should give an error
	err = Copy(exampleFile, exampleFile)
	if err == nil {
		t.Errorf("3: %s\n", err)
	}

	// 4. dir to file should give an error
	err = Copy(testDir, exampleFile)
	if err == nil {
		t.Errorf("4: %s\n", err)
	}

	// 5. dir to an existing dir is not allowed due to infinite recursive loop
	err = Copy(testDir, testDir)
	if err == nil {
		t.Errorf("5: %s\n", err)
	}

	// 6. dir to dir
	err = Copy(testDir, exampleDir)
	if err != nil {
		t.Errorf("6: %s\n", err)
	}

	if !Exists(exampleDir) {
		t.Error("6: exampleDir does not exist")
	}

	if !Exists(filepath.Join(exampleDir, filepath.Base(testFile.Name()))) {
		t.Error("6: exampleFile inside exampleDir does not exist")
	}
}
