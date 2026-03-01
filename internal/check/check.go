// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package check

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Result represents a single check outcome.
type Result struct {
	Pass    bool
	Message string
}

// Run validates the .context/ directory and returns check results.
func Run(dir string) ([]Result, error) {
	contextDir := filepath.Join(dir, ".context")
	var results []Result

	// Check 1: .context/ directory exists
	if info, err := os.Stat(contextDir); err != nil || !info.IsDir() {
		results = append(results, Result{Pass: false, Message: ".context/ directory exists"})
		return results, nil
	}
	results = append(results, Result{Pass: true, Message: ".context/ directory exists"})

	// Check 2: CONTEXT.md exists and is non-empty
	contextMD := filepath.Join(contextDir, "CONTEXT.md")
	info, err := os.Stat(contextMD)
	if err != nil {
		results = append(results, Result{Pass: false, Message: ".context/CONTEXT.md exists and is non-empty"})
		return results, nil
	}
	if info.Size() == 0 {
		results = append(results, Result{Pass: false, Message: ".context/CONTEXT.md exists and is non-empty"})
	} else {
		results = append(results, Result{Pass: true, Message: ".context/CONTEXT.md exists and is non-empty"})
	}

	// Check 3: CONTEXT.md has required sections
	results = append(results, checkRequiredSections(contextMD)...)

	// Check 4: ADR naming convention
	results = append(results, checkNamingConvention(filepath.Join(contextDir, "adrs"), "adrs")...)

	// Check 5: RFC naming convention
	results = append(results, checkNamingConvention(filepath.Join(contextDir, "rfcs"), "rfcs")...)

	// Check 6: Internal links resolve
	results = append(results, checkInternalLinks(contextDir)...)

	return results, nil
}

var requiredSections = []string{
	"what is this",
	"architecture",
	"tech stack",
	"status",
}

func checkRequiredSections(path string) []Result {
	content, err := os.ReadFile(path)
	if err != nil {
		return []Result{{Pass: false, Message: "CONTEXT.md has required sections (could not read file)"}}
	}

	lower := strings.ToLower(string(content))
	var missing []string
	for _, section := range requiredSections {
		// Match ## heading that starts with the section name
		if !strings.Contains(lower, "## "+section) {
			missing = append(missing, section)
		}
	}

	if len(missing) > 0 {
		return []Result{{
			Pass:    false,
			Message: fmt.Sprintf("CONTEXT.md missing required sections: %s", strings.Join(missing, ", ")),
		}}
	}
	return []Result{{Pass: true, Message: "CONTEXT.md has required sections"}}
}

var namingPattern = regexp.MustCompile(`^\d{3}-.+\.md$`)

func checkNamingConvention(dir string, label string) []Result {
	entries, err := os.ReadDir(dir)
	if err != nil {
		// Directory doesn't exist or is empty — that's fine
		return nil
	}

	var results []Result
	for _, entry := range entries {
		name := entry.Name()
		if name == ".gitkeep" || entry.IsDir() {
			continue
		}
		if !namingPattern.MatchString(name) {
			results = append(results, Result{
				Pass:    false,
				Message: fmt.Sprintf(".context/%s/%s does not follow naming convention (expected NNN-slug.md)", label, name),
			})
		}
	}

	if len(results) == 0 && len(entries) > 0 {
		// Only report pass if there were files to check
		hasFiles := false
		for _, entry := range entries {
			if !entry.IsDir() && entry.Name() != ".gitkeep" {
				hasFiles = true
				break
			}
		}
		if hasFiles {
			results = append(results, Result{Pass: true, Message: fmt.Sprintf(".context/%s/ naming convention", label)})
		}
	}

	return results
}

var linkPattern = regexp.MustCompile(`\[([^\]]*)\]\(([^)]+)\)`)

func checkInternalLinks(contextDir string) []Result {
	var results []Result
	broken := false

	err := filepath.Walk(contextDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".md") {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return nil
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		lineNum := 0
		for scanner.Scan() {
			lineNum++
			matches := linkPattern.FindAllStringSubmatch(scanner.Text(), -1)
			for _, match := range matches {
				linkTarget := match[2]
				// Skip external links and anchors
				if strings.HasPrefix(linkTarget, "http") || strings.HasPrefix(linkTarget, "#") {
					continue
				}
				// Resolve relative to the file's directory
				resolved := filepath.Join(filepath.Dir(path), linkTarget)
				if _, err := os.Stat(resolved); err != nil {
					rel, _ := filepath.Rel(contextDir, path)
					results = append(results, Result{
						Pass:    false,
						Message: fmt.Sprintf(".context/%s:%d broken link to %s", rel, lineNum, linkTarget),
					})
					broken = true
				}
			}
		}
		return nil
	})

	if err != nil {
		results = append(results, Result{Pass: false, Message: "Could not scan for internal links"})
	} else if !broken {
		results = append(results, Result{Pass: true, Message: "All internal links resolve"})
	}

	return results
}

// PrintResults prints check results and returns true if all passed.
func PrintResults(results []Result) bool {
	failures := 0
	fmt.Println("dotcontext check")
	fmt.Println()
	for _, r := range results {
		if r.Pass {
			fmt.Printf("  PASS  %s\n", r.Message)
		} else {
			fmt.Printf("  FAIL  %s\n", r.Message)
			failures++
		}
	}
	fmt.Println()
	if failures > 0 {
		fmt.Printf("%d check(s) failed\n", failures)
		return false
	}
	fmt.Println("All checks passed")
	return true
}
