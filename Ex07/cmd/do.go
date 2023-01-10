/*
Copyright © 2023 AMAN KUMAR <KUMARAMAN7056@GMAIL.COM>
*/

package cmd

import (
	"fmt"
	"os"
	"strconv"

	"CLI_TASK_MANAGER/db"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as Complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int

		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the arg: ", arg)
			} else {
				ids = append(ids, id)
			}
		}

		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something wrong happened!", err)
			os.Exit(1)
		}

		for _, id := range ids {
			if id <= 0 || id > len(ids) {
				fmt.Println("Invalid task number: ", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)

			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error %s\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as completed.\n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
