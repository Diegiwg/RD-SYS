package main

import (
	"errors"

	"github.com/Diegiwg/cli"
)

// rd del  [-d] [reg<m|t>]
func DelCommand(ctx *cli.Context) error {
	if len(ctx.Args) < 1 {
		return errors.New("não foi fornecido o registro a ser deletado")
	}

	regRaw := ctx.Args[0]
	reg, err := StringToReg(regRaw)
	if err != nil {
		return err
	}

	ctx.Args = ctx.Args[1:]
	LoadDatabase(ctx).Del(ctx, reg).Save()
	return nil
}
