package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// TODO: Modificar as listas de Previous e Laters para serem Maps
type Reg struct {
	ID       string            // Número de Ordem
	Type     string            // Matricula{m} || Transcrição{t}
	Previous map[string]string // Registros Anteriores
	Laters   map[string]string // Registros Posteriores
}

func (reg *Reg) String() string {
	str := ""

	str += fmt.Sprintf("Número de Ordem: %s\n", reg.ID)
	str += fmt.Sprintf("Tipo: %s\n", reg.Type)

	if len(reg.Previous) > 0 {
		str += "Registros Anteriores:\n"
		for key := range reg.Previous {
			str += fmt.Sprintf("\t%s\n", key)
		}
	} else {
		str += "Nenhum Registro Anterior\n"
	}

	if len(reg.Laters) > 0 {
		str += "Registros Posteriores:\n"
		for key := range reg.Laters {
			str += fmt.Sprintf("\t%s\n", key)
		}
	} else {
		str += "Nenhum Registro Posterior\n"
	}

	return str
}

func (reg *Reg) StrID() string {
	return fmt.Sprintf("%s%s", reg.Type, reg.ID)
}

func StringToReg(str string) (*Reg, error) {
	if len(str) < 2 {
		return nil, errors.New("o registro deve ter pelo menos 2 caracteres")
	}

	str = strings.ToUpper(str)

	if str[0] != 'M' && str[0] != 'T' {
		return nil, errors.New("o primeiro caractere deve ser 'm' ou 't'")
	}

	_type := str[0]

	// Convert string to int
	ID, err := strconv.Atoi(str[1:])
	if err != nil {
		return nil, errors.New("o ID precisa ser um número inteiro")
	}

	return &Reg{
		ID:       fmt.Sprint(ID),
		Type:     string(_type),
		Previous: map[string]string{},
		Laters:   map[string]string{},
	}, nil
}
