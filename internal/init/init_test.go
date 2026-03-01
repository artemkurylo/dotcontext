// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package initcmd

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRun_CreatesStructure(t *testing.T) {
	dir := t.TempDir()

	if err := Run(dir, false); err != nil {
		t.Fatalf("init failed: %v", err)
	}

	// Check all expected directories exist
	expectedDirs := []string{"plans", "stories", "rfcs", "adrs", "runbooks", "guides"}
	for _, sub := range expectedDirs {
		path := filepath.Join(dir, ".context", sub)
		info, err := os.Stat(path)
		if err != nil {
			t.Errorf("expected directory %s to exist: %v", sub, err)
			continue
		}
		if !info.IsDir() {
			t.Errorf("expected %s to be a directory", sub)
		}
	}

	// Check CONTEXT.md exists and is non-empty
	contextMD := filepath.Join(dir, ".context", "CONTEXT.md")
	info, err := os.Stat(contextMD)
	if err != nil {
		t.Fatalf("CONTEXT.md not created: %v", err)
	}
	if info.Size() == 0 {
		t.Error("CONTEXT.md is empty")
	}
}

func TestRun_RefusesOverwrite(t *testing.T) {
	dir := t.TempDir()

	// First init should succeed
	if err := Run(dir, false); err != nil {
		t.Fatalf("first init failed: %v", err)
	}

	// Second init without force should fail
	err := Run(dir, false)
	if err == nil {
		t.Fatal("expected error on second init without --force")
	}
}

func TestRun_ForceOverwrites(t *testing.T) {
	dir := t.TempDir()

	if err := Run(dir, false); err != nil {
		t.Fatalf("first init failed: %v", err)
	}

	// Modify CONTEXT.md
	contextMD := filepath.Join(dir, ".context", "CONTEXT.md")
	os.WriteFile(contextMD, []byte("custom content"), 0o644)

	// Force init should succeed and overwrite
	if err := Run(dir, true); err != nil {
		t.Fatalf("force init failed: %v", err)
	}

	content, _ := os.ReadFile(contextMD)
	if string(content) == "custom content" {
		t.Error("force init did not overwrite CONTEXT.md")
	}
}
