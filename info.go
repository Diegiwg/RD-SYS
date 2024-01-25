package main

import (
	"fmt"

	"github.com/Diegiwg/cli"
)

func InfoCommand(ctx *cli.Context) error {
	fmt.Println("Database: '" + DatabaseFilePath() + "'")
	LoadDatabase(ctx).Dump()
	return nil
}
