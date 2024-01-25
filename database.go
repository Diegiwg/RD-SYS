package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Diegiwg/cli"
)

type Database struct {
	Regs map[string]*Reg
}

func DatabaseFilePath() string {
	userDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	fPath := filepath.Join(userDir, "filepath.txt")
	file, err := os.Open(fPath)
	if err != nil {
		os.Create(fPath)
		file, _ = os.Open(fPath)
	}

	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fp := filepath.FromSlash(filepath.Join(string(content), "database.json"))

	fPtr, err := os.Open(fp)
	if err != nil {
		return "database.json"
	}
	defer fPtr.Close()

	return fp
}

func (db *Database) Dump() {
	fmt.Println("Registros (" + strconv.Itoa(len(db.Regs)) + "):")
	for k := range db.Regs {
		fmt.Println("\t", k)
	}
}

func (db *Database) Save() {
	content, err := json.MarshalIndent(db, "", " ")
	if err != nil {
		panic(err)
	}

	file, err := os.Create(DatabaseFilePath())
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		panic(err)
	}
}

func (db *Database) Find(strID string) (*Reg, error) {
	reg, exists := db.Regs[strID]
	if !exists {
		return nil, errors.New("o registro informado não existe")
	}

	return reg, nil
}

func (db *Database) Add(ctx *cli.Context, reg *Reg) *Database {
	ptr, exists := db.Regs[reg.StrID()]
	if !exists {
		db.Regs[reg.StrID()] = reg
		ptr = db.Regs[reg.StrID()]
	}

	ParsePrevious(ctx, db, ptr)
	return db
}

// mod

// pReg = Registro Anterior
//
// lReg = Registro Posterior
func (db *Database) Set(ctx *cli.Context, pReg, lReg *Reg) *Database {
	pPtr, exists := db.Regs[pReg.StrID()]
	if !exists {
		pPtr, _ = db.Add(ctx, pReg).Find(pReg.StrID())
	}

	lPtr, exists := db.Regs[lReg.StrID()]
	if !exists {
		lPtr, _ = db.Add(ctx, lReg).Find(lReg.StrID())
	}

	pPtr.Previous[lReg.StrID()] = ""
	lPtr.Laters[pReg.StrID()] = ""

	return db
}

func (db *Database) Del(ctx *cli.Context, reg *Reg) *Database {
	ptr, exists := db.Regs[reg.StrID()]
	if !exists {
		return db
	}

	for k := range ptr.Previous {
		delete(db.Regs[k].Laters, reg.StrID())
	}

	for k := range ptr.Laters {
		delete(db.Regs[k].Previous, reg.StrID())
	}

	delete(db.Regs, reg.StrID())
	return db
}

func ParsePrevious(ctx *cli.Context, db *Database, reg *Reg) {
	for i := 0; i < len(ctx.Args); i++ {
		str := ctx.Args[i]

		pReg, err := StringToReg(str)
		if err != nil {
			continue
		}

		// Check if the pReg is already in the list
		_, exists := reg.Previous[pReg.StrID()]
		if exists {
			continue
		} else {
			reg.Previous[pReg.StrID()] = ""
		}

		// Check of the pReg exist in the database and update
		ptr, exists := db.Regs[pReg.StrID()]
		if exists {
			// Check ig the current reg is already in the list
			_, exists := ptr.Laters[reg.StrID()]
			if exists {
				continue
			} else {
				ptr.Laters[reg.StrID()] = ""
			}
		} else {
			pReg.Laters[reg.StrID()] = ""
			db.Regs[pReg.StrID()] = pReg
		}
	}
}

func LoadDatabase(ctx *cli.Context) *Database {
	db := &Database{
		Regs: make(map[string]*Reg),
	}

	file, err := os.Open(DatabaseFilePath())
	if err != nil {
		fmt.Println("INFO: The database doesn't exist! Creating a new one...")
		return db
	}

	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(content, db)
	if err != nil {
		fmt.Println("ERROR: The database is corrupted! Creating a new one, and saving the old...")
		file.Close()
		os.Rename(DatabaseFilePath(), DatabaseFilePath()+".old")
		return db
	}

	return db
}
