package symbol_test

import (
	"deassembler/symbol"
	"testing"
)

func TestNewSymbolTable(t *testing.T) {
	st := symbol.NewSymbolTable()

	expectedSymbols := map[string]int{
		"R0": 0, "R1": 1, "R2": 2, "R3": 3, "R4": 4, "R5": 5, "R6": 6, "R7": 7,
		"R8": 8, "R9": 9, "R10": 10, "R11": 11, "R12": 12, "R13": 13, "R14": 14, "R15": 15,
		"SP": 0, "LCL": 1, "ARG": 2, "THIS": 3, "THAT": 4,
		"SCREEN": 16384, "KBD": 24576,
	}

	for sym, addr := range expectedSymbols {
		if st.GetAddress(sym) != addr {
			t.Errorf("Expected %s to be %d, got %d", sym, addr, st.GetAddress(sym))
		}
	}
}

func TestAddEntry(t *testing.T) {
	st := symbol.NewSymbolTable()

	st.AddEntry("LOOP", 10)
	if st.GetAddress("LOOP") != 10 {
		t.Errorf("Expected LOOP to be 10, got %d", st.GetAddress("LOOP"))
	}

	st.AddEntry("END", 20)
	if st.GetAddress("END") != 20 {
		t.Errorf("Expected END to be 20, got %d", st.GetAddress("END"))
	}
}

func TestContains(t *testing.T) {
	st := symbol.NewSymbolTable()

	if !st.Contains("SP") {
		t.Errorf("Expected SP to be in symbol table, but it's not")
	}

	st.AddEntry("MY_VAR", 100)
	if !st.Contains("MY_VAR") {
		t.Errorf("Expected MY_VAR to be in symbol table, but it's not")
	}

	if st.Contains("NON_EXISTENT") {
		t.Errorf("Expected NON_EXISTENT not to be in symbol table, but it is")
	}
}

func TestGetNextVarSymbol(t *testing.T) {
	st := symbol.NewSymbolTable()

	expected := []string{"VAR_0", "VAR_1", "VAR_2"}
	for i, exp := range expected {
		actual := st.GetNextVarSymbol()
		if actual != exp {
			t.Errorf("Expected VAR_%d to be %s, got %s", i, exp, actual)
		}
	}
}
