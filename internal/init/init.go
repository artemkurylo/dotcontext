// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package initcmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/artemkurylo/dotcontext/internal/templates"
)

var subdirs = []string{
	"plans",
	"stories",
	"rfcs",
	"adrs",
	"runbooks",
	"guides",
}

// Run creates the .context/ directory structure with a starter CONTEXT.md.
func Run(dir string, force bool) error {
	contextDir := filepath.Join(dir, ".context")

	if info, err := os.Stat(contextDir); err == nil && info.IsDir() {
		if !force {
			return fmt.Errorf(".context/ already exists (use --force to reinitialize)")
		}
	}

	// Create .context/ and all subdirectories
	for _, sub := range subdirs {
		path := filepath.Join(contextDir, sub)
		if err := os.MkdirAll(path, 0o755); err != nil {
			return fmt.Errorf("failed to create %s: %w", path, err)
		}
	}

	// Write CONTEXT.md from embedded template
	tmpl, err := templates.Files.ReadFile("context.md.tmpl")
	if err != nil {
		return fmt.Errorf("failed to read CONTEXT.md template: %w", err)
	}

	contextMD := filepath.Join(contextDir, "CONTEXT.md")
	if err := os.WriteFile(contextMD, tmpl, 0o644); err != nil {
		return fmt.Errorf("failed to write CONTEXT.md: %w", err)
	}

	fmt.Println("Initialized .context/ directory:")
	fmt.Println("  .context/CONTEXT.md")
	for _, sub := range subdirs {
		fmt.Printf("  .context/%s/\n", sub)
	}
	fmt.Println()
	fmt.Println("Next step: edit .context/CONTEXT.md to describe your project.")

	return nil
}
