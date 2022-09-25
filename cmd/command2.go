package cmd

import (
  "fmt"
  "os/exec"

  "github.com/spf13/cobra"
)

var command2Cmd = &cobra.Command{
  Use:   "2",
  Short: "Print the version number of Hugo",
  Long:  `All software has versions. This is Hugo's`,
  Run: func(cmd *cobra.Command, args []string) {
	out, err := exec.Command("go", "version").Output()
	if err != nil {
	  fmt.Printf("%s", err)
	}
	fmt.Println("Command2 output: ", string(out))
  },
}


