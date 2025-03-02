package cmd

import (
	"fmt"
	"os"
	"todo/helpers"
	"github.com/spf13/cobra"
)

// all the command
var addCmd = &cobra.Command{
    Use: "add",
    Short: "Adding task",
    Long: "Command to ad task",
    Run: addCommand,
}

var listCmd = &cobra.Command{
    Use: "list",
    Short: "Adding task",
    Long: "Command to add task",
    Run: listCommand,
}

var comCmd = &cobra.Command{
    Use: "complete",
    Short: "Complete task",
    Long: "Command to complete task status",
    Run: completeCommand,
}

var deleteCmd = &cobra.Command{
    Use: "delete",
    Short: "Complete task",
    Long: "Command to complete task status",
    Run: deleteCommand,
}

// adding command to root with their flags
func init(){
    rootCmd.AddCommand(addCmd)
    addCmd.Flags().StringP("task", "t", " ", "To add task your task")

    // list command
    rootCmd.AddCommand(listCmd)
    listCmd.Flags().BoolP("top", "F", false, "To get top five")
    listCmd.Flags().BoolP("bottom", "B", false, "To get last five")

    // complete command
    rootCmd.AddCommand(comCmd)
    comCmd.Flags().IntP("id", "i", 0, "to complete the status of taskw with id")

    // delete task
    rootCmd.AddCommand(deleteCmd)
    deleteCmd.Flags().IntP("id", "i", 0, "To complete task status")
}

// command corresponding functions
func addCommand(cmd *cobra.Command, args []string){
    text, _ := cmd.Flags().GetString("task")
    if len(text)>0{
        helpers.AddTask(text)
    }
}

func listCommand(cmd *cobra.Command, args []string){
    top , _ := cmd.Flags().GetBool("top")
    last , _ := cmd.Flags().GetBool("bottom")
    if top {
        fmt.Println("HEAD")
        os.Exit(0)
    }else if last {
        fmt.Println("TAIL")
        os.Exit(0)
    }else {
        helpers.ListTasks()
        os.Exit(0)
    }
}

func completeCommand(cmd *cobra.Command, args[]string){
    id, _ := cmd.Flags().GetInt("id")
    if id > 0 {
        helpers.UpdateTask(id)
    }
}

func deleteCommand(cmd *cobra.Command, args[]string){
    id, _ := cmd.Flags().GetInt("id")
    if id > 0 {
        helpers.DeleteTask(id)
    }
}

