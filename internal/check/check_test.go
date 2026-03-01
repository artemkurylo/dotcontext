// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package check

import (
	"os"
	"path/filepath"
	"testing"
)

func setupValidContext(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	contextDir := filepath.Join(dir, ".context")

	for _, sub := range []string{"plans", "stories", "rfcs", "adrs", "runbooks", "guides"} {
		os.MkdirAll(filepath.Join(contextDir, sub), 0o755)
	}

	content := `# Test Project

## What is this?
A test project.

## Architecture
Monolith.

## Tech stack
Go.

## Status
Active.
`
	os.WriteFile(filepath.Join(contextDir, "CONTEXT.md"), []byte(content), 0o644)
	return dir
}

func TestRun_ValidContext(t *testing.T) {
	dir := setupValidContext(t)

	results, err := Run(dir)
	if err != nil {
		t.Fatalf("check failed: %v", err)
	}

	for _, r := range results {
		if !r.Pass {
			t.Errorf("expected pass but got fail: %s", r.Message)
		}
	}
}

func TestRun_MissingContextDir(t *testing.T) {
	dir := t.TempDir()

	results, err := Run(dir)
	if err != nil {
		t.Fatalf("check failed: %v", err)
	}

	if len(results) == 0 {
		t.Fatal("expected at least one result")
	}
	if results[0].Pass {
		t.Error("expected fail for missing .context/ directory")
	}
}

func TestRun_MissingContextMD(t *testing.T) {
	dir := t.TempDir()
	os.MkdirAll(filepath.Join(dir, ".context"), 0o755)

	results, err := Run(dir)
	if err != nil {
		t.Fatalf("check failed: %v", err)
	}

	hasFail := false
	for _, r := range results {
		if !r.Pass {
			hasFail = true
			break
		}
	}
	if !hasFail {
		t.Error("expected fail for missing CONTEXT.md")
	}
}

func TestRun_EmptyContextMD(t *testing.T) {
	dir := t.TempDir()
	contextDir := filepath.Join(dir, ".context")
	os.MkdirAll(contextDir, 0o755)
	os.WriteFile(filepath.Join(contextDir, "CONTEXT.md"), []byte(""), 0o644)

	results, err := Run(dir)
	if err != nil {
		t.Fatalf("check failed: %v", err)
	}

	hasFail := false
	for _, r := range results {
		if !r.Pass {
			hasFail = true
			break
		}
	}
	if !hasFail {
		t.Error("expected fail for empty CONTEXT.md")
	}
}

func TestRun_MissingSections(t *testing.T) {
	dir := t.TempDir()
	contextDir := filepath.Join(dir, ".context")
	os.MkdirAll(contextDir, 0o755)
	os.WriteFile(filepath.Join(contextDir, "CONTEXT.md"), []byte("# Project\nSome text."), 0o644)

	results, err := Run(dir)
	if err != nil {
		t.Fatalf("check failed: %v", err)
	}

	hasSectionFail := false
	for _, r := range results {
		if !r.Pass && contains(r.Message, "missing required sections") {
			hasSectionFail = true
			break
		}
	}
	if !hasSectionFail {
		t.Error("expected fail for missing required sections")
	}
}

func TestRun_BadADRNaming(t *testing.T) {
	dir := setupValidContext(t)
	// Create a badly named ADR
	os.WriteFile(filepath.Join(dir, ".context", "adrs", "my-decision.md"), []byte("# Bad"), 0o644)

	results, err := Run(dir)
	if err != nil {
		t.Fatalf("check failed: %v", err)
	}

	hasNamingFail := false
	for _, r := range results {
		if !r.Pass && contains(r.Message, "naming convention") {
			hasNamingFail = true
			break
		}
	}
	if !hasNamingFail {
		t.Error("expected fail for bad ADR naming")
	}
}

func TestRun_ValidADRNaming(t *testing.T) {
	dir := setupValidContext(t)
	os.WriteFile(filepath.Join(dir, ".context", "adrs", "001-use-go.md"), []byte("# ADR"), 0o644)

	results, err := Run(dir)
	if err != nil {
		t.Fatalf("check failed: %v", err)
	}

	for _, r := range results {
		if !r.Pass {
			t.Errorf("expected all pass but got fail: %s", r.Message)
		}
	}
}

func TestRun_BrokenLink(t *testing.T) {
	dir := setupValidContext(t)
	// Create a file with a broken internal link
	os.WriteFile(
		filepath.Join(dir, ".context", "adrs", "001-test.md"),
		[]byte("See [plan](../plans/nonexistent.md)"),
		0o644,
	)

	results, err := Run(dir)
	if err != nil {
		t.Fatalf("check failed: %v", err)
	}

	hasLinkFail := false
	for _, r := range results {
		if !r.Pass && contains(r.Message, "broken link") {
			hasLinkFail = true
			break
		}
	}
	if !hasLinkFail {
		t.Error("expected fail for broken internal link")
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && searchString(s, substr)
}

func searchString(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
