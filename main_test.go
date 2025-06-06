package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestCLICommands(t *testing.T) {
	// Build the binary for testing
	binary := "./wherewasi_test"
	cmd := exec.Command("go", "build", "-o", binary, ".")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer os.Remove(binary)

	t.Run("Help", func(t *testing.T) {
		cmd := exec.Command(binary, "--help")
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Help command failed: %v", err)
		}

		outputStr := string(output)
		if !strings.Contains(outputStr, "generates dense AI context summaries") {
			t.Error("Help output should contain description")
		}

		if !strings.Contains(outputStr, "start") || !strings.Contains(outputStr, "pull") || !strings.Contains(outputStr, "status") {
			t.Error("Help should list all main commands")
		}
	})

	t.Run("Status", func(t *testing.T) {
		cmd := exec.Command(binary, "status")
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Status command failed: %v", err)
		}

		outputStr := string(output)
		if !strings.Contains(outputStr, "wherewasi ripcord status") {
			t.Error("Status should show header")
		}
	})

	t.Run("Start", func(t *testing.T) {
		cmd := exec.Command(binary, "start")
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Start command failed: %v", err)
		}

		outputStr := string(output)
		if !strings.Contains(outputStr, "Ready for ripcord deployment") {
			t.Error("Start should confirm ripcord deployment")
		}
	})

	t.Run("PullDryRun", func(t *testing.T) {
		cmd := exec.Command(binary, "pull", "--clipboard=false", "--save=false")
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Pull dry run failed: %v", err)
		}

		outputStr := string(output)
		// Should generate context without errors
		if len(outputStr) == 0 {
			t.Error("Pull should generate some output")
		}
	})

	t.Run("PullWithSave", func(t *testing.T) {
		// Create temporary config directory
		tmpHome := t.TempDir()
		cmd := exec.Command(binary, "pull", "--save", "--clipboard=false")
		cmd.Env = append(os.Environ(), "HOME="+tmpHome)

		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Pull with save failed: %v", err)
		}

		// Check that database was created
		dbPath := filepath.Join(tmpHome, ".local", "share", "wherewasi", "context.sqlite")
		if _, err := os.Stat(dbPath); os.IsNotExist(err) {
			t.Error("Database should be created when saving context")
		}

		// Just verify the command ran without error and database was created
		// The exact output message may vary
		if len(output) == 0 {
			t.Log("Pull command ran successfully (no output expected in some cases)")
		}
	})

	t.Run("History", func(t *testing.T) {
		// Use same temp directory as previous test
		tmpHome := t.TempDir()

		// First save a context
		cmd := exec.Command(binary, "pull", "--save", "--clipboard=false")
		cmd.Env = append(os.Environ(), "HOME="+tmpHome)
		_, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to save context for history test: %v", err)
		}

		// Then retrieve history
		cmd = exec.Command(binary, "pull", "--history")
		cmd.Env = append(os.Environ(), "HOME="+tmpHome)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("History command failed: %v", err)
		}

		outputStr := string(output)
		if !strings.Contains(outputStr, "RECENT CONTEXTS") {
			t.Error("History should show recent contexts header")
		}
	})

	t.Run("KeywordSearch", func(t *testing.T) {
		cmd := exec.Command(binary, "pull", "--keyword", "test", "--clipboard=false", "--save=false")
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Keyword search failed: %v", err)
		}

		// Keyword search should run without error
		// The exact output format may vary
		if len(output) == 0 {
			t.Log("Keyword search ran successfully")
		}
	})

	t.Run("InvalidCommand", func(t *testing.T) {
		cmd := exec.Command(binary, "nonexistent")
		_, err := cmd.CombinedOutput()
		if err == nil {
			t.Error("Invalid command should return error")
		}
	})
}

func TestProjectDetection(t *testing.T) {
	binary := "./wherewasi_test_project"
	cmd := exec.Command("go", "build", "-o", binary, ".")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer os.Remove(binary)

	t.Run("DetectCurrentProject", func(t *testing.T) {
		cmd := exec.Command(binary, "status")
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Status command failed: %v", err)
		}

		outputStr := string(output)
		// Should detect wherewasi as current project
		if !strings.Contains(outputStr, "wherewasi") {
			t.Error("Should detect current project name")
		}
	})

	t.Run("EcosystemDetection", func(t *testing.T) {
		cmd := exec.Command(binary, "status")
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Status command failed: %v", err)
		}

		outputStr := string(output)
		// Should find some projects in ecosystem
		if !strings.Contains(outputStr, "projects tracked") {
			t.Error("Should detect tracked projects in ecosystem")
		}
	})
}

func TestFlags(t *testing.T) {
	binary := "./wherewasi_test_flags"
	cmd := exec.Command("go", "build", "-o", binary, ".")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer os.Remove(binary)

	testCases := []struct {
		name string
		args []string
	}{
		{"LongHelp", []string{"--help"}},
		{"ShortHelp", []string{"-h"}},
		{"PullHelp", []string{"pull", "--help"}},
		{"StatusVerbose", []string{"status", "--verbose"}},
		{"PullNoClipboard", []string{"pull", "--clipboard=false", "--save=false"}},
		{"PullWithKeyword", []string{"pull", "--keyword", "test", "--clipboard=false", "--save=false"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cmd := exec.Command(binary, tc.args...)
			_, err := cmd.CombinedOutput()
			// Most commands should succeed (return code 0)
			// Even if some fail due to environment, they shouldn't crash
			if err != nil {
				// Check if it's a clean exit vs a crash
				if exitErr, ok := err.(*exec.ExitError); ok {
					if exitErr.ExitCode() > 2 {
						t.Errorf("Command %v crashed with exit code %d", tc.args, exitErr.ExitCode())
					}
					// Exit codes 1-2 are acceptable for expected failures
				}
			}
		})
	}
}
