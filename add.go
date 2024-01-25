package main

import (
	"errors"

	"github.com/Diegiwg/cli"
)

// rg add  [-a] [reg<m|t>] [regs: <m|t>*]
func AddCommand(ctx *cli.Context) error {
	if len(ctx.Args) < 1 {
		return errors.New("o registro deve ter pelo menos 2 caracteres")
	}

	regRaw := ctx.Args[0]
	reg, err := StringToReg(regRaw)
	if err != nil {
		return err
	}

	ctx.Args = ctx.Args[1:]
	LoadDatabase(ctx).Add(ctx, reg).Save()

	return nil
}
