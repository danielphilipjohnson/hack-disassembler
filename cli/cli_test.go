package cli_test

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestDisassemblerCLI(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir := t.TempDir()

	// Copy test.hack to the temporary directory
	inputHackPath := filepath.Join(tmpDir, "test.hack")
	if err := copyFile("testdata/test.hack", inputHackPath); err != nil {
		t.Fatalf("Failed to copy test.hack: %v", err)
	}

	// Define output path
	outputAsmPath := filepath.Join(tmpDir, "test.asm")

	// Build the executable
	hackdisasmPath := buildHackDisasm(t, tmpDir)

	// Run the disassembler
	cmd := exec.Command(hackdisasmPath, inputHackPath, "-o", outputAsmPath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to run hackdisasm: %v\n%s", err, out)
	}

	// Read the expected output
	expectedBytes, err := ioutil.ReadFile("testdata/test.asm.expected")
	if err != nil {
		t.Fatalf("Failed to read expected output: %v", err)
	}
	expected := strings.Split(string(expectedBytes), "\n")

	// Read the actual output
	actualBytes, err := ioutil.ReadFile(outputAsmPath)
	if err != nil {
		t.Fatalf("Failed to read actual output: %v", err)
	}
	actual := strings.Split(string(actualBytes), "\n")
	if len(actual) > 0 && actual[len(actual)-1] == "" {
		actual = actual[:len(actual)-1]
	}

	// Compare outputs, ignoring the first line (timestamp)
	if len(actual) != len(expected) {
		t.Errorf("Output line count mismatch: expected %d, got %d", len(expected), len(actual))
	}

	for i := 1; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Line %d mismatch: expected %q, got %q", i+1, expected[i], actual[i])
		}
	}
}

func TestDisassemblerCLIUsage(t *testing.T) {
	tmpDir := t.TempDir()
	hackdisasmPath := buildHackDisasm(t, tmpDir)

	cmd := exec.Command(hackdisasmPath)
	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatalf("Expected hackdisasm to exit with usage error, but it succeeded: %s", string(out))
	}

	if !strings.Contains(string(out), "Usage: hackdisasm") {
		t.Fatalf("Expected usage message, got: %s", string(out))
	}
}

func buildHackDisasm(t *testing.T, workDir string) string {
	t.Helper()
	cacheDir := filepath.Join(workDir, "gocache")
	if err := os.MkdirAll(cacheDir, 0o755); err != nil {
		t.Fatalf("Failed to create cache dir: %v", err)
	}

	binPath := filepath.Join(workDir, "hackdisasm")
	cmd := exec.Command("go", "build", "-o", binPath, ".")
	cmd.Dir = projectRoot(t)
	cmd.Env = append(os.Environ(), "GOCACHE="+cacheDir)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to build hackdisasm: %v\n%s", err, out)
	}

	return binPath
}

func projectRoot(t *testing.T) string {
	t.Helper()
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatalf("Failed to determine caller path")
	}

	return filepath.Dir(filepath.Dir(filename))
}

func copyFile(src, dst string) error {
	input, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dst, input, 0644)
	if err != nil {
		return err
	}
	return nil
}
