package main

import (
	"errors"

	"github.com/Diegiwg/cli"
)

// rg set  [-s] [reg<m|t>] [anterior: reg<m|t>]
func SetCommand(ctx *cli.Context) error {
	if len(ctx.Args) < 2 {
		return errors.New("devem ser fornecidos o Registro Anterior e o Registro")
	}

	// Registro Anterior
	pReg, err := StringToReg(ctx.Args[0])
	if err != nil {
		return err
	}

	// Registro Posterior
	lReg, err := StringToReg(ctx.Args[1])
	if err != nil {
		return err
	}

	LoadDatabase(ctx).Set(ctx, pReg, lReg).Save()
	return nil
}
