/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"go.jolman.me/tcs-notes/notes"
)

var (
	date string

	rootCmd = &cobra.Command{
		Use:   "tcs-notes <student-names>",
		Short: "create session note and homework for a student",
		Long: `tcs-notes is for quickly creating session notes and homework
for specified student(s).

It uses editable tamplate files that open open in your text editor
and copies the file contents to your clipboard when finished.`,

		Run: func(cmd *cobra.Command, args []string) {
			for _, arg := range args {
				notes.WriteNoteAndHW(arg, date)
			}
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&date, "date", "d", "01-02", "Edit a note that's not from today")
}
