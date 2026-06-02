CLI application for after-session notes at theCoderSchool.

# Installation
```
go install go.jolman.me/tcs-notes@latest
```

# Usage
First create your template files:
```
tcs-notes --edit-note
tcs-notes --edit-hw
```

Then:
```
tcs-notes <student-name>
```

This will create a note and hw text file at `~/.local/share/tcs-notes/students/<student-name>/<date>/`, copy the template files into them, and then open them into your text editor (uses the `$EDITOR` environment variable).  
The contents of the file will be automatically copied to your clipboard when you close your text editor.

> Above I used a Unix file path but I'm pretty sure this will work on Windows as well
