// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

const manifestToolBin = "/manifest-tool"

// execCmd is a helper function to
// run the provided command.
func execCmd(e *exec.Cmd) error {
	logrus.Tracef("executing cmd %s", strings.Join(e.Args, " "))

	// set command stdout to OS stdout
	e.Stdout = os.Stdout
	// set command stderr to OS stderr
	e.Stderr = os.Stderr

	// output "trace" string for command
	fmt.Println("$", strings.Join(e.Args, " "))

	return e.Run()
}

// versionCmd is a helper function to output
// the client version information.
func versionCmd() *exec.Cmd {
	logrus.Trace("creating manifest-tool version command")

	// variable to store flags for command
	var flags []string

	// add flag to print version of manifest-tool command
	flags = append(flags, "--version")

	return exec.Command(manifestToolBin, flags...)
}
