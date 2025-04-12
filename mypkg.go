package main

import (
  "os/exec"
  "os"
)

func main() {
  cmd := &exec.Cmd {
    // Path is the path of the command to run.
    Path: "/bin/echo",
    // Args holds command line arguments, including the command itself as Args[0].
    Args: []string{ "echo", "hello world" },
    Stdout: os.Stdout,
    Stderr: os.Stdout,
  }
  cmd.Start();
  cmd.Wait();

  // Args can be also ommited and {Path} will be used by default
  cmd := &exec.Cmd {
    Path: "/bin/echo"
  }

  // Vulnerable
  userInput := "/pwn/exploit" // value supplied by user input
  cmd := &exec.Cmd {
    Path: userInput
  }

  // Vulnerable
  userInput1 := "/bin/bash" // value supplied by user input
  userInput2 := []string{ "bash", "exploit.sh" } // value supplied by user input
  cmd := &exec.Cmd {
    Path: userInput1,
    Args: userInput2
  }

}
