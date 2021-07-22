package psqlutil

import (
	"github.com/pkg/errors"
	"github.com/toowoxx/go-lib-userspace-common/cmds"
)

func DumpDatabase(host, username, password, dbName, file string) error {
	if err := cmds.RunCommandWithEnv("pg_dump",
		map[string]string{
			"PGPASSWORD": password,
		},
		"-Fc", "-v",
		"--host="+host,
		"--username="+username,
		"--dbname="+dbName,
		"-f", file,
	); err != nil {
		return errors.Wrap(err, "could not run pg_dump command")
	}

	return nil
}
