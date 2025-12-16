package disassembler

import (
	"deassembler/symbol"
	"fmt"
	"strconv"
	"strings"
)

type Disassembler struct {
	symbolTable *symbol.SymbolTable
	compMapA    map[string]string // For a=0
	compMapM    map[string]string // For a=1
	destMap     map[string]string
	jumpMap     map[string]string
}

func NewDisassembler() *Disassembler {
	return &Disassembler{
		symbolTable: symbol.NewSymbolTable(),
		compMapA: map[string]string{
			"101010": "0",
			"111111": "1",
			"111010": "-1",
			"001100": "D",
			"110000": "A",
			"001101": "!D",
			"110001": "!A",
			"001111": "-D",
			"110011": "-A",
			"011111": "D+1",
			"110111": "A+1",
			"001110": "D-1",
			"110010": "A-1",
			"000010": "D+A",
			"010011": "D-A",
			"000111": "A-D",
			"000000": "D&A",
			"010000": "D|A",
		},
		compMapM: map[string]string{
			"110000": "M",
			"110001": "!M",
			"110011": "-M",
			"110111": "M+1",
			"110010": "M-1",
			"000010": "D+M",
			"010011": "D-M",
			"000111": "M-D",
			"000000": "D&M",
			"010000": "D|M",
		},
		destMap: map[string]string{
			"000": "", "001": "M", "010": "D", "011": "MD",
			"100": "A", "101": "AM", "110": "AD", "111": "AMD",
		},
		jumpMap: map[string]string{
			"000": "", "001": "JGT", "010": "JEQ", "011": "JGE",
			"100": "JLT", "101": "JNE", "110": "JLE", "111": "JMP",
		},
	}
}

func (d *Disassembler) DisassembleAInstruction(binary string) (string, error) {
	address, err := strconv.ParseInt(binary[1:], 2, 16)
	if err != nil {
		return "", fmt.Errorf("invalid A-instruction address: %s", binary)
	}

	// 1. Check for R0-R15 (direct mapping)
	if address >= 0 && address <= 15 {
		return fmt.Sprintf("@R%d", address), nil
	}

	// 2. Check if address is already mapped to a symbol (predefined or VAR_X)
	if symbol, ok := d.symbolTable.ReverseSymbols[int(address)]; ok {
		return fmt.Sprintf("@%s", symbol), nil
	}

	// 3. If address >= 16 and not found in reverseSymbols, it's a new variable
	if address >= 16 {
		varName := d.symbolTable.GetNextVarSymbol() // Use GetNextVarSymbol
		d.symbolTable.AddEntry(varName, int(address))
		return fmt.Sprintf("@%s", varName), nil
	}

	// 4. Otherwise, it's a direct numerical address (e.g., @100)
	return fmt.Sprintf("@%d", address), nil
}

func (d *Disassembler) DisassembleCInstruction(binary string) (string, error) {
	if !strings.HasPrefix(binary, "111") {
		return "", fmt.Errorf("invalid C-instruction format: %s", binary)
	}

	aBit := binary[3]
	comp := binary[4:10]
	dest := binary[10:13]
	jump := binary[13:16]

	var compMnemonic string
	var ok bool

	if aBit == '0' {
		compMnemonic, ok = d.compMapA[comp]
	} else {
		compMnemonic, ok = d.compMapM[comp]
	}

	if !ok {
		return "", fmt.Errorf("unknown comp mnemonic for: %s (aBit: %c)", comp, aBit)
	}

	destMnemonic := d.destMap[dest]
	jumpMnemonic := d.jumpMap[jump]

	var parts []string
	if destMnemonic != "" {
		parts = append(parts, destMnemonic+"=")
	}
	parts = append(parts, compMnemonic)
	if jumpMnemonic != "" {
		parts = append(parts, ";"+jumpMnemonic)
	}

	return strings.Join(parts, ""), nil
}
