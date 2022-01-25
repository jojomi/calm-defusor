package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/jojomi/calm-defusor/communication"
	"github.com/jojomi/calm-defusor/modules"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func main() {
	var (
		mod modules.Module
		err error
	)

	log.Logger = log.With().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})

	mods := modules.NewModuleList().AddAllAvailable()

	isFirstModule := true
	for {
		if !isFirstModule {
			fmt.Println()
			fmt.Println()
		}

		// select
		mod, err = communication.ChooseOneStringable[modules.Module]("Nächstes Modul?", mods.All())
		if err != nil {
			if err == terminal.InterruptErr {
				os.Exit(0)
			}
			log.Error().Err(err).Msg("selection failed")
			continue
		}

		isFirstModule = false
		mod.Reset()
		err = mod.Solve()
		if err != nil {
			if err == terminal.InterruptErr {
				log.Error().Err(err).Msgf(`Module "%s" aborted`, mod.Name())
				continue
			}
			log.Error().Err(err).Msgf(`Module "%s"failed`, mod.Name())
		}
	}
}
