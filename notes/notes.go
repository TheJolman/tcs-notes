package notes

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/adrg/xdg"
)

const (
	appName    = "tcs-notes"
	timeLayout = "01-02"
)

var (
	notesDir         = path.Join(xdg.DataHome, appName, "students")
	configDir        = path.Join(xdg.ConfigHome, appName)
	noteTemplatePath = path.Join(configDir, "note_template.md")
	hwTemplatePath   = path.Join(configDir, "hw_template.txt")
)

// Open note, then hw file in editor for user to edit.
// Date must be in mm-dd format.
func WriteNoteAndHW(studentName string, date string) error {
	if date == "" {
		currTime := time.Now()
		date = currTime.Format(timeLayout)
	} else {
		if _, err := time.Parse(timeLayout, date); err != nil {
			return fmt.Errorf("invalid date format")
		}
	}
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
	hwPath := path.Join(studentDateDir, "hw.txt")
	if err := writeFile(hwPath); err != nil {
		return fmt.Errorf("failed to create hemework: %v", err)
	}
	return nil
}

// Opens note template in editor for user to create/edit
func WriteNoteTemplate() error {
	if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create config directory: %v", err)
	}

	if err := writeFile(noteTemplatePath); err != nil {
		return fmt.Errorf("failed to create note template: %v", err)
	}
	return nil
}

// Opens hw template in editor for user to create/edit
func WriteHWTemplate() error {
	if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create config directory: %v", err)
	}

	if err := writeFile(hwTemplatePath); err != nil {
		return fmt.Errorf("failed to create hw template: %v", err)
	}
	return nil
}

// ==== UTILITIES ==================================================================================

func writeFile(filepath string) error {
	editor := os.Getenv("EDITOR")
	// TODO: this could be made more robust, especially for windows users.
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
