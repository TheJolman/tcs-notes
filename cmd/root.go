/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.jolman.me/tcs-notes/notes"
)

var (
	editNote bool
	editHW   bool
	date     string

	rootCmd = &cobra.Command{
		Use:   "tcs-notes <student-names>",
		Short: "create session note and homework for a student",
		Long: `tcs-notes is for quickly creating session notes and homework
for specified student(s).

It uses editable tamplate files that open open in your text editor
and copies the file contents to your clipboard when finished.
If this is your first time running this program, run with --edit-note and
--edit-hw to customize your template files.`,

		Run: func(cmd *cobra.Command, args []string) {
			if editNote {
				cobra.CheckErr(notes.WriteNoteTemplate())
				return
			}
			if editHW {
				cobra.CheckErr(notes.WriteHWTemplate())
				return
			}
			fullName := args[0]
			if len(args) > 1 {
				for i := 1; i < len(args); i++ {
					fullName = fmt.Sprintf("%s_%s", fullName, args[i])
				}
			}
			err := notes.WriteNoteAndHW(fullName, date)
			cobra.CheckErr(err)
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
	rootCmd.Flags().BoolVar(&editNote, "edit-note", false, "Edit note template")
	rootCmd.Flags().BoolVar(&editHW, "edit-hw", false, "Edit homework template")
	rootCmd.Flags().StringVarP(&date, "date", "d", "", "Edit a note that's not from today")
}
