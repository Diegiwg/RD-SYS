package main

import (
	"fmt"
	"strings"

	"github.com/Diegiwg/cli"
)

func main() {
	app := cli.NewApp()

	app.AddCommand(&cli.Command{
		Name:  "find",
		Desc:  "Procura no sistema, o registro informado, e mostra as informações do registro",
		Help:  "Procura no sistema, o registro informado, e mostra as informações do registro",
		Usage: "[reg<m|t>]",
		Exec:  FindCommand,
	})

	app.AddCommand(&cli.Command{
		Name:  "add",
		Desc:  "Adiciona um novo registro ao sistema",
		Help:  "Adiciona um novo registro ao sistema",
		Usage: "[reg<m|t>] [regs: <m|t>*]",
		Exec:  AddCommand,
	})

	// mod

	app.AddCommand(&cli.Command{
		Name:  "set",
		Desc:  "Marcar um registro como Registro Anterior de outro registro",
		Help:  "Marcar um registro como Registro Anterior de outro registro",
		Usage: "[reg<m|t>] [anterior: reg<m|t>]",
		Exec:  SetCommand,
	})

	app.AddCommand(&cli.Command{
		Name:  "del",
		Desc:  "Deleta um registro do sistema",
		Help:  "Deleta um registro do sistema",
		Usage: "[reg<m|t>]",
		Exec:  DelCommand,
	})

	app.AddCommand(&cli.Command{
		Name:  "info",
		Desc:  "Mostra as informações do sistema",
		Help:  "Mostra as informações do sistema",
		Usage: "",
		Exec:  InfoCommand,
	})

	err := app.Run()
	if err != nil {
		fmt.Println("ERROR:", strings.ToTitle(err.Error()))
	}
}
