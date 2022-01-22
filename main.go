package main

import (
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/jojomi/calm-defusor/modules"
	"github.com/jojomi/go-script/v2/interview"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	mods := modules.NewModuleList().AddAllAvailable()
	moduleNames := mods.GetNames()

	var (
		nextModuleName string
		mod            modules.Module
		err            error
	)

	for {
		// select
		nextModuleName, err = interview.ChooseOneString("Nächstes Modul?", moduleNames)
		if err != nil {
			if err == terminal.InterruptErr {
				os.Exit(0)
			}
			log.Error().Err(err).Msg("selection failed")
			continue
		}

		mod = mods.GetByName(nextModuleName)
		if mod == nil {
			log.Error().Err(err).Msgf("module not found: %s", nextModuleName)
			continue
		}

		err = mod.Solve()
		if err != nil {
			panic(err)
		}
	}
}
