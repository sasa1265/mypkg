package main

import (
    "fmt"
    "os"
)

func main() {
    // استخدام مكتبة os لقراءة متغير بيئة
    user := os.Getenv("USER")
    fmt.Println("User:", user)
}
