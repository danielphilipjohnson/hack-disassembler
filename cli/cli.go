package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"deassembler/disassembler"
)

// Disassembles the input Hack file and writes the output to an ASM file.
func Run() {
	inputPath := ""
	outPath := ""

	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		if args[i] == "-o" && i+1 < len(args) {
			outPath = args[i+1]
			i++
		} else {
			inputPath = args[i]
		}
	}

	if inputPath == "" {
		fmt.Println("Usage: hackdisasm <input.hack> [-o <output.asm>]")
		os.Exit(1)
	}

	if outPath == "" {
		outPath = strings.TrimSuffix(inputPath, ".hack") + ".asm"
	}

	inputFile, err := os.Open(inputPath)
	if err != nil {
		fmt.Printf("Error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outPath)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	_, err = writer.WriteString(fmt.Sprintf("// Disassembled on %s by HackDisasm v1.0\n", time.Now().Format("2006-01-02 15:04:05")))
	if err != nil {
		fmt.Printf("Error writing header: %v\n", err)
		os.Exit(1)
	}

	disassembler := disassembler.NewDisassembler()
	scanner := bufio.NewScanner(inputFile)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if len(line) != 16 {
			fmt.Printf("Error: Line %d has invalid length (%d). Expected 16 bits: %s\n", lineNumber, len(line), line)
			os.Exit(1)
		}

		for _, char := range line {
			if char != '0' && char != '1' {
				fmt.Printf("Error: Line %d contains invalid character '%c'. Expected '0' or '1': %s\n", lineNumber, char, line)
				os.Exit(1)
			}
		}

		var asmLine string
		var disassembleErr error
		if line[0] == '0' {
			asmLine, disassembleErr = disassembler.DisassembleAInstruction(line)
		} else if line[0] == '1' {
			asmLine, disassembleErr = disassembler.DisassembleCInstruction(line)
		}

		if disassembleErr != nil {
			fmt.Printf("Error disassembling line %d: %v\n", lineNumber, disassembleErr)
			os.Exit(1)
		}

		_, err = writer.WriteString(asmLine + "\n")
		if err != nil {
			fmt.Printf("Error writing to output file: %v\n", err)
			os.Exit(1)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}

	writer.Flush()
	fmt.Printf("Disassembly complete. Output written to %s\n", outPath)
}
