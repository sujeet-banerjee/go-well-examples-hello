The Structure Explained: Modules vs. Packages vs. Folders
Think of the hierarchy like this: Folders contain Packages, and a Module contains Folders/Packages.

1. The Package (package main)
A package is simply a directory containing one or more .go source files compiled together.

The Rule: Every .go file must declare a package at the top.

'package main' is special: It tells the Go compiler, "This is an executable program, not a reusable library." It mandates that you must have a func main() inside it. For any other package (e.g., package math), compiling it just creates a library archive, not something you can run.

2. The Module (go.mod)
A module is a collection of Go packages (directories) that are versioned together. The go.mod file marks the root of a module.

When you run go mod init examples/hello, it creates the go.mod file.

The string examples/hello is your module path.

The module path acts as the base prefix for importing any package within this module. If you created a subfolder inside hello called utils, you would import it in your hello.go file as import "examples/hello/utils".

3. The Folder Structure
For a simple "Hello World" executable, the physical folder name (C:\Sujeet\GoWorkspace\hello) does not strictly matter to the compiler. However, by convention, the folder name usually matches either the module name or the package name to keep things organized.

CLI:

cd ../GoWorkspace/hello

go run .
>> <executes the main func>


go build examples/hello
>> hello.exe

==================================


4. What is GOROOT? (The Engine)
GOROOT is the directory where the Go SDK itself is installed on your computer. It contains the Go compiler, the core Go tools, and the standard library (packages like fmt, math, os, etc.).

Typical Location: On Windows, it is usually C:\Program Files\Go. On macOS/Linux, it is often /usr/local/go.

Why GoLand is asking for it: Your IDE needs to know where the Go installation lives so it can compile your code, provide syntax highlighting, and allow you to navigate standard library code. Without GOROOT, GoLand doesn't know where to find the code for import "fmt".

5. What is GOPATH? (The Old Workspace & New Cache)
GOPATH used to be the most important environment variable in Go, but today, its role is mostly in the background.

The Old Way (Pre-Modules): Before modules existed, Go dictated that all of your source code, regardless of the project, had to live inside the GOPATH/src directory. You could not just create C:\Sujeet\GoWorkspace\hello. You would have had to put it in something like C:\Users\Sujeet\go\src\github.com\sujeet\hello.

The New Way (With Modules): Because of Go Modules, you can now put your code anywhere (like you did in C:\Sujeet\GoWorkspace). Today, GOPATH acts primarily as a cache. When you download third-party libraries (e.g., go get github.com/gin-gonic/gin), Go stores those downloaded files inside GOPATH/pkg/mod. It also uses GOPATH/bin to store executable files you install via go install.

Typical Location: It usually defaults to a go folder in your user directory (e.g., C:\Users\Sujeet\go).

How They Connect to Your Code Structure
Let's tie GOROOT, GOPATH, and your examples/hello module together using your hello.go file:
---------------
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
-----------------
Here is exactly how the Go compiler uses these paths when it reads your code:

Step 1: Standard Library (GOROOT)
When the compiler sees import "fmt", it immediately checks your GOROOT. It knows that fmt is a built-in standard library package, so it pulls the necessary code right out of your Go installation directory.

Step 2: Your Code (The Module)
When compiling hello.go, it looks at your go.mod file to understand the current context. It recognizes that the code in C:\Sujeet\GoWorkspace\hello belongs to the module examples/hello.

Step 3: Third-Party Dependencies (GOPATH)
If you were to add a line like import "github.com/google/uuid", the compiler would first look in GOROOT. Since it's not a standard package, it would then look in your module cache located in your GOPATH (specifically GOPATH/pkg/mod/github.com/google/uuid...). If it isn't there, Go will download it and save it to that GOPATH cache for future use.

=======================================

To locate your exact GOPATH and GOROOT using PowerShell on Windows, you have two primary methods. While their typical locations are often C:\Program Files\Go for GOROOT and C:\Users\<YourName>\go for GOPATH, it is always best to ask the system directly.

Here is exactly how to find them.

Method 1: Ask Go Directly (Highly Recommended)
This is the most reliable method. Even if you haven't explicitly set these variables in your Windows environment settings, the Go toolchain has built-in defaults that it falls back on.

Open PowerShell and run the following command to see all Go environment variables (PowerShell):
go env

To pinpoint the exact paths without scrolling through the whole list, query them specifically. Find GOROOT (PowerShell):
go env GOROOT