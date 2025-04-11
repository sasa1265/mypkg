package main

import (
    "os"
    "os/exec"
    "log"
)

func main() {
    cmd := exec.Command("whoami")  // مثال لأمر غير مرغوب فيه، يمكنك استبداله بما يناسب
    out, err := cmd.CombinedOutput()
    if err != nil {
        log.Fatal(err)
    }
    log.Println(string(out))
}
