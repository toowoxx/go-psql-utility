package psqlutil

import (
	"os"
	"path"
	"testing"
)

func TestRestoreDatabase(t *testing.T) {
	dir, err := os.MkdirTemp("", "db_dump_test")
	if err != nil {
		t.Fatal(err)
	}
	err = DumpDatabase(
		"postgres",
		"test",
		"testpw",
		"testdb",
		path.Join(dir, "test.dump"),
	)
	if err != nil {
		t.Fatal(err)
	}

	err = RestoreDatabase(
		"postgres",
		"test",
		"testpw",
		"testdb",
		path.Join(dir, "test.dump"),
	)
	if err != nil {
		t.Fatal(err)
	}

}
