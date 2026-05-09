package notes

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/adrg/xdg"
)

const appName = "tcs-notes"

var notesDir = path.Join(xdg.DataHome, appName, "students")
var configDir = path.Join(xdg.ConfigHome, appName)

func WriteNoteAndHW(studentName string, date string) error {
	// date must be in mm-dd format
	// TODO: use time.Time instead
	studentDateDir := path.Join(notesDir, studentName, date)
	if err := os.MkdirAll(studentDateDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create student folder: %v", err)
	}
	// Write note
	notePath := path.Join(studentDateDir, "note.md")
	if err := writeFile(notePath); err != nil {
		return fmt.Errorf("failed to create note: %v", err)
	}
	// Write homework
	hwPath := path.Join(studentDateDir, "hw.md")
	if err := writeFile(hwPath); err != nil {
		return fmt.Errorf("failed to create hemework: %v", err)
	}
	return nil
}

func MakeNoteTemplate() {
	_ = configDir
}

func writeFile(filepath string) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "nano"
	}

	cmd := exec.Command(editor, filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	return err
}
