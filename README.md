# LetsGo
A Github Repo where i Deep Dive GO.

# Go Application Development Guide

## Getting Started

This guide will help you set up, run, and debug your Go applications.

### Prerequisites

- **Go**: Install Go from the [official website](https://golang.org/dl/). Verify installation with:
  ```bash
  go version
  ```
- **IDE/Editor**: We recommend Visual Studio Code with the Go extension, but any text editor will work.

### Setting Up a New Project

1. **Initialize a Module**:
   ```bash
   mkdir my-project
   cd my-project
   go mod init my-project
   ```

2. **Project Structure**:
   ```
   my-project/
   ├── cmd/                    # Command applications
   │   └── main.go            # Main application entry point
   ├── internal/              # Private application code
   ├── pkg/                   # Public library code
   ├── go.mod                # Module dependencies
   └── go.sum                # Dependency checksums
   ```

## Building and Running

### Basic Commands

1. **Build the Project**:
   ```bash
   go build ./...            # Build all packages
   go build -o app ./cmd/main.go  # Build specific binary
   ```

2. **Run the Application**:
   ```bash
   go run ./cmd/main.go      # Build and run
   ```

3. **Install Dependencies**:
   ```bash
   go get github.com/example/package
   go mod tidy              # Clean up dependencies
   ```

### Development Tools

1. **Code Formatting**:
   ```bash
   go fmt ./...             # Format all code
   ```

2. **Static Analysis**:
   ```bash
   go vet ./...            # Check for common errors
   golint ./...            # Style checker
   ```

## Debugging

### Using Delve Debugger

1. **Install Delve**:
   ```bash
   go install github.com/go-delve/delve/cmd/dlv@latest
   ```

2. **Start Debugging**:
   ```bash
   dlv debug ./cmd/main.go
   ```

3. **Common Delve Commands**:
   - `break main.go:20` - Set breakpoint
   - `continue` - Continue execution
   - `next` - Step over
   - `step` - Step into
   - `print variable` - Print variable value

### VS Code Integration

1. Configure `launch.json`:
   ```json
   {
       "version": "0.2.0",
       "configurations": [
           {
               "name": "Launch Package",
               "type": "go",
               "request": "launch",
               "mode": "auto",
               "program": "${workspaceFolder}/cmd/main.go"
           }
       ]
   }
   ```

## Common Issues and Solutions

### Build Errors

1. **Missing Dependencies**:
   ```bash
   go mod tidy
   ```

2. **Incompatible Versions**:
   - Check `go.mod` for version conflicts
   - Update specific dependencies:
     ```bash
     go get -u package@version
     ```

3. **GOPATH Issues**:
   - Ensure `GOPATH` is set correctly:
     ```bash
     go env GOPATH
     ```
   - Add to PATH:
     ```bash
     export PATH=$PATH:$(go env GOPATH)/bin
     ```

### Performance Profiling

1. **CPU Profiling**:
   ```go
   import "runtime/pprof"
   
   f, _ := os.Create("cpu.prof")
   pprof.StartCPUProfile(f)
   defer pprof.StopCPUProfile()
   ```

2. **Memory Profiling**:
   ```go
   import "runtime/pprof"
   
   f, _ := os.Create("mem.prof")
   defer f.Close()
   pprof.WriteHeapProfile(f)
   ```

3. **Analyze Profiles**:
   ```bash
   go tool pprof cpu.prof
   go tool pprof mem.prof
   ```

## Best Practices

1. **Error Handling**:
   ```go
   if err != nil {
       return fmt.Errorf("failed to process: %w", err)
   }
   ```