package modules

import (
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/jojomi/calm-defusor/communication"
	"github.com/rs/zerolog/log"
)

type SymbolsModule struct {
	rows [][]SymbolsSymbol
}

func (x SymbolsModule) Name() string {
	return "Symbole"
}

func (x SymbolsModule) String() string {
	return x.Name()
}

func NewSymbolsModule() *SymbolsModule {
	rows := [][]SymbolsSymbol{
		{SymbolsSymbolOStrich, SymbolsSymbolAT, SymbolsSymbolLambda, SymbolsSymbolKleinesN, SymbolsSymbolRakete, SymbolsSymbolH, SymbolsSymbolFalschesC},
		{SymbolsSymbolEuro, SymbolsSymbolOStrich, SymbolsSymbolFalschesC, SymbolsSymbolKringel, SymbolsSymbolLeererStern, SymbolsSymbolH, SymbolsSymbolFragezeichen},
		{SymbolsSymbolCopyright, SymbolsSymbolW, SymbolsSymbolKringel, SymbolsSymbolDoppelK, SymbolsSymbolR, SymbolsSymbolLambda, SymbolsSymbolLeererStern},
		{SymbolsSymbolDelta, SymbolsSymbolP, SymbolsSymbolBT, SymbolsSymbolRakete, SymbolsSymbolDoppelK, SymbolsSymbolFragezeichen, SymbolsSymbolZunge},
		{SymbolsSymbolKronleuchter, SymbolsSymbolZunge, SymbolsSymbolBT, SymbolsSymbolC, SymbolsSymbolP, SymbolsSymbolDrei, SymbolsSymbolVollerStern},
		{SymbolsSymbolDelta, SymbolsSymbolEuro, SymbolsSymbolDoppelkreuz, SymbolsSymbolAE, SymbolsSymbolKronleuchter, SymbolsSymbolN, SymbolsSymbolOmega},
	}

	return &SymbolsModule{
		rows: rows,
	}
}

func (x *SymbolsModule) Reset() error {
	return nil
}

func (x *SymbolsModule) Solve() error {
	numSymbols := 4

	allSymbols := x.allSymbols()
	symbols := make([]SymbolsSymbol, 0, numSymbols)
	for i := 1; i <= numSymbols; i++ {
		sym, err := communication.ChooseOneWithMapper[SymbolsSymbol](fmt.Sprintf("Symbol %d?", i), allSymbols, func(symbol SymbolsSymbol) string {
			return symbol.Name()
		})
		if err != nil {
			if err == terminal.InterruptErr {
				return nil
			}
			log.Error().Err(err).Msg("could not get symbol")
			i--
			continue
		}
		symbols = append(symbols, sym)
	}

	row, err := x.findRow(symbols)
	if err != nil {
		panic(err)
	}
	x.printOrdered(symbols, row)
	return nil
}

func (x SymbolsModule) findRow(symbols []SymbolsSymbol) ([]SymbolsSymbol, error) {
outer:
	for _, row := range x.rows {
		for _, sym := range symbols {
			symFound := false
			for _, rowSym := range row {
				if rowSym != sym {
					continue
				}
				symFound = true
			}
			if !symFound {
				continue outer
			}
		}
		return row, nil
	}
	return nil, errors.New("cant find row")
}

func (x SymbolsModule) printOrdered(symbols []SymbolsSymbol, row []SymbolsSymbol) {
	for _, rowSym := range row {
		for _, sym := range symbols {
			if rowSym != sym {
				continue
			}
			communication.Tell(sym.Name())
			break
		}
	}
}

func (x SymbolsModule) allSymbols() []SymbolsSymbol {
	all := make([]SymbolsSymbol, 0)
	allSet := make(map[SymbolsSymbol]struct{}, 0)
	for _, row := range x.rows {
		for _, sym := range row {
			allSet[sym] = struct{}{}
		}
	}

	for key := range allSet {
		all = append(all, key)
	}
	return all
}

//go:generate go-enum -f=$GOFILE --marshal
// SymbolsSymbol is an enumeration of symbols
/*
ENUM(
OStrich
AT
Lambda
kleinesN
Rakete
H
C
FalschesC
Euro
Kringel
LeererStern
Fragezeichen
Copyright
W
DoppelK
R
Delta
P
BT
Zunge
Kronleuchter
Drei
VollerStern
Doppelkreuz
AE
N
Omega
)
*/
type SymbolsSymbol int

func (x SymbolsSymbol) Aliases() []string {
	switch x {
	case SymbolsSymbolOStrich:
		return []string{"O mit Strich"}
	case SymbolsSymbolAT:
		return []string{"AT", "A mit Dach", "Zelt", "Kirchendach"}
	case SymbolsSymbolLambda:
		return []string{"Lambda"}
	case SymbolsSymbolKleinesN:
		return []string{"Kleines N"}
	case SymbolsSymbolRakete:
		return []string{"startende Rakete"}
	case SymbolsSymbolH:
		return []string{"Schnörkel-H"}
	case SymbolsSymbolC:
		return []string{"C mit Punkt"}
	case SymbolsSymbolFalschesC:
		return []string{"verkehrtes C mit Punkt"}
	case SymbolsSymbolEuro:
		return []string{"umgedrehter Euro mit Punkten"}
	case SymbolsSymbolKringel:
		return []string{"Kringel"}
	case SymbolsSymbolLeererStern:
		return []string{"leerer Stern"}
	case SymbolsSymbolFragezeichen:
		return []string{"umgedrehtes Fragzeichen"}
	case SymbolsSymbolCopyright:
		return []string{"Copyright"}
	case SymbolsSymbolW:
		return []string{"W mit Verzierung"}
	case SymbolsSymbolDoppelK:
		return []string{"Doppel-K"}
	case SymbolsSymbolR:
		return []string{"Roger Federer"}
	case SymbolsSymbolDelta:
		return []string{"Delta"}
	case SymbolsSymbolP:
		return []string{"Word-P"}
	case SymbolsSymbolBT:
		return []string{"BT"}
	case SymbolsSymbolZunge:
		return []string{"Gesicht mit Zunge"}
	case SymbolsSymbolKronleuchter:
		return []string{"Kronleuchter"}
	case SymbolsSymbolDrei:
		return []string{"Drei mit Hörnern und Schwanz"}
	case SymbolsSymbolVollerStern:
		return []string{"Voller Stern"}
	case SymbolsSymbolDoppelkreuz:
		return []string{"Doppelkreuz"}
	case SymbolsSymbolAE:
		return []string{"AE"}
	case SymbolsSymbolN:
		return []string{"Russisches N"}
	case SymbolsSymbolOmega:
		return []string{"Omega"}
	default:
		return []string{x.String()}
	}
}

func (x SymbolsSymbol) Name() string {
	aliases := x.Aliases()
	if len(aliases) == 0 {
		return ""
	}
	return aliases[0]
}
