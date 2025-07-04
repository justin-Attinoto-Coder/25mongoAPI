# Getting Started with Go in Visual Studio Code

1. Install Go

Download Go from [https://go.dev/dl/](https://go.dev/dl/).
Follow installer instructions for your OS (Windows, macOS, Linux).
Verify installation: go version in terminal.

1. Set Up Visual Studio Code

Install VS Code: [https://code.visualstudio.com/](https://code.visualstudio.com/).
Open VS Code, go to Extensions (Ctrl+Shift+X or Cmd+Shift+X).
Install "Go" extension by Go Team at Google.

1. Configure Go in VS Code

Open terminal in VS Code (Ctrl+``or Cmd+``).
Install Go tools: go install golang.org/x/tools/gopls@latest.
Set GOPATH: go env -w GOPATH=$HOME/go (or custom path).
Ensure $GOPATH/bin is in your PATH.

1. Create and Run Go Project

Create a folder, open it in VS Code.
Initialize module: go mod init example.com/myapp.
Create main.go:package main
import "fmt"
func main() {
    fmt.Println("Hello, Go!")
}

Run: go run main.go.

1. Debugging

Install Delve: go install github.com/go-delve/delve/cmd/dlv@latest.
Set breakpoints in VS Code, use "Run and Debug" (F5).

1. Additional Setup

Format on Save: Enable in VS Code settings (editor.formatOnSave: true) for auto-formatting Go code.
Go Modules: Always use go mod for dependency management (go get for packages).
Testing: Create *_test.go files, run tests with go test.

1. Recommended Tools

Install gofmt for code formatting: go install golang.org/x/tools/cmd/gofmt@latest.
Install goimports for import management: go install golang.org/x/tools/cmd/goimports@latest.

1. Resources

Official Go Docs: [https://go.dev/doc/](https://go.dev/doc/).
Go Tour: [https://tour.golang.org/](https://tour.golang.org/).
VS Code Go Wiki: [https://github.com/golang/vscode-go/wiki](https://github.com/golang/vscode-go/wiki).

Done.
