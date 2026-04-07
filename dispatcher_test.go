// dispatcher_test.go
package main

import (
	"strings"
	"testing"
)

func TestBuildAppleScript(t *testing.T) {
	msg := "SPEED UP CLANKER"
	script := buildAppleScript(msg)

	expected := `tell application "System Events"
  key code 8 using {control down}
  delay 0.3
  keystroke "SPEED UP CLANKER"
  delay 0.05
  key code 36
end tell`

	if script != expected {
		t.Errorf("unexpected script:\ngot:  %q\nwant: %q", script, expected)
	}
}

func TestBuildAppleScriptEscapesQuotes(t *testing.T) {
	msg := `HE SAID "YOU'RE SLOW" AND HE'S RIGHT`
	script := buildAppleScript(msg)

	if !strings.Contains(script, `\"YOU'RE SLOW\"`) {
		t.Errorf("expected escaped quotes in script: %s", script)
	}
}
