package main

import (
	"errors"
	"fmt"

	"github.com/Diegiwg/cli"
)

// rg find [-f] [reg<m|t>]
func FindCommand(ctx *cli.Context) error {
	if len(ctx.Args) < 1 {
		return errors.New("o registro deve ter pelo menos 2 caracteres")
	}

	regRaw := ctx.Args[0]
	reg, err := StringToReg(regRaw)
	if err != nil {
		return err
	}

	reg, err = LoadDatabase(ctx).Find(reg.StrID())
	if err != nil {
		return err
	}

	fmt.Println(reg.String())
	return nil
}
