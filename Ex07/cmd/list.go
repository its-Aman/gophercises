/*
Copyright © 2023 AMAN KUMAR <KUMARAMAN7056@GMAIL.COM>
*/

package cmd

import (
	"fmt"
	"os"

	"CLI_TASK_MANAGER/db"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()

		if err != nil {
			fmt.Println("Something wrong happened!", err)
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("No pending tasks! It's off today.")
			return
		}

		fmt.Println("You have the following tasks:")

		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
