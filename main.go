package main

import (
	"log"
	"os"

	"github.com/calvinfeng/redisgob/cmd"
	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{
		Use:   "redisgob",
		Short: "inter process communication example",
	}

	root.AddCommand(cmd.Serve, cmd.Dial)
	if err := root.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
