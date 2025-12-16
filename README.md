# Hack Disassembler

This command-line tool disassembles Hack machine code (`.hack` files) into Hack assembly code (`.asm` files). It accompanies the Hack computer from the Nand2Tetris course and is implemented in Go.

## Features

- Disassembles A-instructions and C-instructions.
- Handles predefined symbols (R0-R15, SP, LCL, ARG, THIS, THAT, SCREEN, KBD).
- Assigns generic variable names for addresses >= 16 (e.g., `VAR_0`, `VAR_1`).
- Validates input `.hack` file format.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/danielphilipjohnson/hack-disassembler.git
   cd hack-disassembler
   ```

2. Build the executable (outputs `hackdisasm` in the repo root):

   ```bash
   go build -o hackdisasm ./...
   ```

## Usage

```bash
./hackdisasm <input.hack> [-o <output.asm>]
```

- `<input.hack>` Path to the Hack machine code file.
- `-o <output.asm>` Optional output path. Must be provided on the same command (no newline). Defaults to the input basename + `.asm` when omitted.

### Examples

Disassemble `samples/simple.hack` and save the output to `output.asm`:

```bash
./hackdisasm samples/simple.hack -o output.asm
```

## Testing

Run all unit and CLI tests from the repo root:

```bash
go test ./...
```

The CLI test suite builds the binary in a temporary directory, so no extra setup is required.

## Sample `.hack` Programs

Use the bundled files in `samples/` to try the CLI without sourcing your own machine code:

```bash
./hackdisasm samples/simple.hack -o /tmp/simple.asm
./hackdisasm samples/branch.hack -o /tmp/branch.asm
./hackdisasm samples/simple.hack   # uses samples/simple.asm automatically
```
