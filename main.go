package main

// Import the fmt for formatting strings
// Import os so we can read environment variables from the system
import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, Kubernetes！I'm from Jenkins CI！")
	fmt.Println("分支名:", os.Getenv("branch"))
}
