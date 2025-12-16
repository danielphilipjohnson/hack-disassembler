package symbol

import "fmt"

type SymbolTable struct {
	symbols        map[string]int
	ReverseSymbols map[int]string
	varCounter     int
}

func NewSymbolTable() *SymbolTable {
	st := &SymbolTable{
		symbols:        make(map[string]int),
		ReverseSymbols: make(map[int]string),
		varCounter:     0,
	}
	for i := 0; i <= 15; i++ {
		st.AddEntry(fmt.Sprintf("R%d", i), i)
	}
	st.AddEntry("SP", 0)
	st.AddEntry("LCL", 1)
	st.AddEntry("ARG", 2)
	st.AddEntry("THIS", 3)
	st.AddEntry("THAT", 4)
	st.AddEntry("SCREEN", 16384)
	st.AddEntry("KBD", 24576)
	return st
}

func (st *SymbolTable) AddEntry(symbol string, address int) {
	st.symbols[symbol] = address
	st.ReverseSymbols[address] = symbol
}

func (st *SymbolTable) Contains(symbol string) bool {
	_, ok := st.symbols[symbol]
	return ok
}

func (st *SymbolTable) GetAddress(symbol string) int {
	return st.symbols[symbol]
}

func (st *SymbolTable) GetNextVarSymbol() string {
	varName := fmt.Sprintf("VAR_%d", st.varCounter)
	st.varCounter++
	return varName
}
