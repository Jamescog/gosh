# gosh

A lightweight shell implementation written in Go for educational purposes.

## Overview

gosh is a simple shell that provides basic file system operations and system commands. It works on Unix-like systems and Windows 10 or higher.

**Note:** This project is built for learning purposes and is not intended for production use.

## Features

- Interactive command-line interface
- Built-in help system with command documentation
- Support for command flags and arguments
- Cross-platform support

## Requirements

- Go 1.24.0 or higher
- Windows 10 or higher (for Windows users)

## Installation

### From Source

1. Clone the repository:
```bash
git clone https://github.com/Jamescog/gosh.git
cd gosh
```

2. Build the binary:
```bash
go build
```

3. Run the shell:
```bash
./gosh
```

### Alternative: Install Globally

```bash
go install github.com/Jamescog/gosh@latest
```

## Usage

Launch the shell by running the compiled binary:

```bash
./gosh
```

You will be greeted with the gosh prompt:
```
-> gosh:
```

Type commands followed by their arguments and press Enter to execute them.

### Getting Help

To see all available commands:
```bash
help
```

To get detailed information about a specific command:
```bash
help <command>
```

Example:
```bash
help ls
```

## Available Commands

### File System Operations

#### ls
List directory contents.

**Usage:** `ls [OPTIONS] [PATH]`

**Options:**
- `-a` - Show all files including hidden files
- `-l` - Use long listing format
- `-h` - Human-readable file sizes (use with `-l`)
- `-lh` - Long format with human-readable sizes

**Examples:**
```bash
ls
ls -a
ls -l
ls -lh /home
```

#### cd
Change the current working directory.

**Usage:** `cd DIRECTORY`

**Examples:**
```bash
cd /home
cd ..
cd ~
```

#### pwd
Print the current working directory.

**Usage:** `pwd`

**Example:**
```bash
pwd
```

#### mkdir
Create directories.

**Usage:** `mkdir [OPTIONS] DIRECTORY...`

**Options:**
- `-p` - Create parent directories as needed
- `-v` - Verbose output
- `-m MODE` - Set file permissions (default: 0755)

**Examples:**
```bash
mkdir newdir
mkdir -p dir1/dir2/dir3
mkdir -v newdir
mkdir -m 0700 private
```

#### touch
Create empty files or update file timestamps.

**Usage:** `touch FILE...`

**Examples:**
```bash
touch file.txt
touch file1.txt file2.txt
```

### System Commands

#### clear
Clear the terminal screen.

**Usage:** `clear`

#### help
Display help information about commands.

**Usage:** `help [COMMAND]`

**Examples:**
```bash
help
help ls
```

#### exit
Exit the shell.

**Usage:** `exit`

## Project Structure

```
gosh/
├── main.go                 # Application entry point
├── go.mod                  # Go module definition
├── commands/               # Command implementations
│   ├── base.go            # Command interface definition
│   ├── registry.go        # Command registry system
│   ├── cd.go              # Change directory command
│   ├── clear.go           # Clear screen command
│   ├── exit.go            # Exit shell command
│   ├── help.go            # Help command
│   ├── ls.go              # List directory command
│   ├── mkdir.go           # Make directory command
│   ├── pwd.go             # Print working directory command
│   └── touch.go           # Touch file command
└── utils/                  # Utility functions
    ├── color_other.go     # Color support for Unix-like systems
    ├── color_windows.go   # Color support for Windows
    ├── file_uitls.go      # File utility functions
    ├── parse_command.go   # Command parsing logic
    └── pretty_print.go    # Output formatting utilities
```

## Building for Different Platforms

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o gosh-linux

# macOS
GOOS=darwin GOARCH=amd64 go build -o gosh-darwin

# Windows
GOOS=windows GOARCH=amd64 go build -o gosh.exe
```

## License

This project is available under the MIT License.

## Author

Jamescog

## Acknowledgments

Built with Go's standard library and minimal external dependencies for maximum portability and performance.
