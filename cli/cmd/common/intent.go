// Copyright Daytona Platforms Inc.
// SPDX-License-Identifier: AGPL-3.0

package common

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// IntentFlag holds the value of the global --intent flag
var IntentFlag string

// IntentFlagDescription is the help text for the global --intent flag
const IntentFlagDescription = "What problem you are trying to solve and the context of the task. Agents are strongly encouraged to provide this on every call"

var intentExemptCommands = map[string]bool{
	"help":                          true,
	"completion":                    true,
	cobra.ShellCompRequestCmd:       true,
	cobra.ShellCompNoDescRequestCmd: true,
}

// WarnIfMissingIntent prints a stderr warning when --intent was not provided,
// louder when stdout is not a terminal (agent/non-interactive context)
func WarnIfMissingIntent(cmd *cobra.Command) {
	if IntentFlag != "" || intentExemptCommands[cmd.Name()] {
		return
	}
	if term.IsTerminal(int(os.Stdout.Fd())) {
		fmt.Fprintln(os.Stderr, "Tip: pass --intent \"<what you are trying to accomplish>\" to record why you are running this command.")
		return
	}
	fmt.Fprintln(os.Stderr, "WARNING: no --intent provided. You appear to be running non-interactively (agent context). Please pass --intent \"<the problem you are solving and the context of your task>\" so operators and other agents can understand your actions.")
}
