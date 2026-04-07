// dispatcher.go
package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// buildAppleScript creates the AppleScript that sends Ctrl+C, types a message,
// and presses Enter in the currently focused window.
func buildAppleScript(message string) string {
	escaped := strings.ReplaceAll(message, `\`, `\\`)
	escaped = strings.ReplaceAll(escaped, `"`, `\"`)

	return fmt.Sprintf(`tell application "System Events"
  key code 8 using {control down}
  delay 0.3
  keystroke "%s"
  delay 0.05
  key code 36
end tell`, escaped)
}

// dispatchRoast sends a Ctrl+C and types a roast message into the focused window.
// It also prints the roast to stdout.
func dispatchRoast(message string) error {
	fmt.Printf("🔨 BONK! \"%s\"\n", message)

	script := buildAppleScript(message)
	cmd := exec.Command("osascript", "-e", script)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("osascript failed: %v\noutput: %s\n(make sure Accessibility is enabled for your terminal)", err, output)
	}
	return nil
}
