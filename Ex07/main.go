/*
Copyright © 2023 AMAN KUMAR <KUMARAMAN7056@GMAIL.COM>
*/

package main

import (
	"CLI_TASK_MANAGER/cmd"
	"CLI_TASK_MANAGER/db"
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
