package disassembler_test

import (
	"deassembler/disassembler"
	"testing"
)

func TestDisassembleAInstruction(t *testing.T) {
	d := disassembler.NewDisassembler()

	tests := []struct {
		name     string
		binary   string
		expected string
	}{
		{name: "R0", binary: "0000000000000000", expected: "@R0"},
		{name: "R15", binary: "0000000000001111", expected: "@R15"},
		{name: "SCREEN", binary: "0100000000000000", expected: "@SCREEN"},
		{name: "KBD", binary: "0110000000000000", expected: "@KBD"},
		{name: "VAR_0", binary: "0000000000010000", expected: "@VAR_0"},
		{name: "VAR_1", binary: "0000000000010001", expected: "@VAR_1"},
		{name: "Address 100", binary: "0000000001100100", expected: "@VAR_2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := d.DisassembleAInstruction(tt.binary)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if actual != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, actual)
			}
		})
	}
}

func TestDisassembleCInstruction(t *testing.T) {
	d := disassembler.NewDisassembler()

	tests := []struct {
		name     string
		binary   string
		expected string
	}{
		{name: "D=M", binary: "1111110000010000", expected: "D=M"},
		{name: "D=D+A", binary: "1110000010010000", expected: "D=D+A"},
		{name: "0;JMP", binary: "1110101010000111", expected: "0;JMP"},
		{name: "MD=D+1;JLE", binary: "1110011111011110", expected: "MD=D+1;JLE"},
		{name: "A=M-D", binary: "1111000111100000", expected: "A=M-D"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := d.DisassembleCInstruction(tt.binary)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if actual != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, actual)
			}
		})
	}
}
