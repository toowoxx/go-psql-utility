package psqlutil

import (
	"github.com/pkg/errors"
	"github.com/toowoxx/go-lib-userspace-common/cmds"
)

func RestoreDatabase(host, username, password, dbName, file string) error {
	if err := cmds.RunCommandWithEnv("pg_restore",
		map[string]string{
			"PGPASSWORD": password,
		},
		"-v", "--no-owner",
		"--host="+host,
		"--username="+username,
		"--dbname="+dbName,
		file,
	); err != nil {
		return errors.Wrap(err, "could not run pg_restore command")
	}

	return nil
}
