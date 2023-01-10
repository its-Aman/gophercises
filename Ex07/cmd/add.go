/*
Copyright © 2023 AMAN KUMAR <KUMARAMAN7056@GMAIL.COM>
*/

package cmd

import (
	"fmt"
	"os"
	"strings"

	"CLI_TASK_MANAGER/db"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your tasklist",

	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)

		if err != nil {
			fmt.Println("Something wrong happened!", err)
			os.Exit(1)
		}

		fmt.Printf("Added \"%s\" task to your list", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
