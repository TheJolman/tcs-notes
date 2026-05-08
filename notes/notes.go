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

func WriteNote(studentName string, date string) error {
	studentDir := path.Join(notesDir, studentName)
	if err := os.MkdirAll(studentDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create student folder: %v", err)
	}
	filepath := path.Join(studentDir, date+".md")
	err := writeFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to create note: %v", err)
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
