// pkg/dns/dns.go
package dns

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

// ClearCache flushes the DNS cache.
func ClearCache() error {
	log.Println("Flushing DNS cache...")

	var cmd *exec.Cmd

	// Check the OS type to determine the appropriate command
	if runtime.GOOS == "darwin" {
		// macOS
		cmd = exec.Command("dscacheutil", "-flushcache")
	} else if runtime.GOOS == "linux" {
		// Debian-based Linux (e.g., Ubuntu)
		cmd = exec.Command("systemctl", "restart", "systemd-resolved")
	} else {
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	// Execute the command
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
