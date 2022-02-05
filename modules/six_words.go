package modules

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/jojomi/calm-defusor/communication"
	"github.com/rs/zerolog/log"
)

type SixWordsModule struct {
	bigPositionMap map[string]SixWordsPosition
	dangerWords    map[string]string
	pushMap        map[string][]string
}

func (x SixWordsModule) Name() string {
	return "6 Wörter"
}

func (x SixWordsModule) String() string {
	return x.Name()
}

func NewSixWordsModule() *SixWordsModule {
	bigPositionMap := map[string]SixWordsPosition{
		"JA":       SixWordsPositionMiddleLeft,
		"MOMENT":   SixWordsPositionTopRight,
		"OBEN":     SixWordsPositionBottomRight,
		"OKAY":     SixWordsPositionTopRight,
		"DA STEHT": SixWordsPositionBottomRight,
		"NICHTS":   SixWordsPositionMiddleLeft,
		"   ":      SixWordsPositionBottomLeft,
		"LEER":     SixWordsPositionMiddleRight,
		"NEIN":     SixWordsPositionBottomRight,
		"KUH":      SixWordsPositionMiddleLeft,
		"Q":        SixWordsPositionBottomRight,
		"COUP":     SixWordsPositionMiddleRight,
		"WARTE":    SixWordsPositionMiddleRight,
		"OH GOTT":  SixWordsPositionBottomLeft,
		"FERTIG":   SixWordsPositionBottomLeft,
		"BUMM":     SixWordsPositionBottomRight,
		"SO EIN":   SixWordsPositionMiddleRight,
		"SO'N":     SixWordsPositionBottomRight,
		"SOHN":     SixWordsPositionMiddleRight,
		"ZEHN":     SixWordsPositionMiddleRight,
		"CN":       SixWordsPositionTopLeft,
		"ZEHEN":    SixWordsPositionBottomRight,
		"10":       SixWordsPositionBottomLeft,
		"ZÄH":      SixWordsPositionMiddleRight,
		"ZEH":      SixWordsPositionMiddleLeft,
		"CE":       SixWordsPositionBottomRight,
		"C":        SixWordsPositionTopRight,
		"ZU SPÄT":  SixWordsPositionBottomRight,
	}
	pushMap := map[string][]string{
		"DRÜCK":    {"JA", "Q", "NEIN", "O.K.", "OKAY", "COUP", "OK", "LEER", "DRÜCK", "   ", "NOCHMAL", "WAS", "NICHTS", "KUH"},
		"NOCHMAL":  {"OKAY", "Q", "JA", "O.K.", "   ", "OK", "NICHTS", "WAS", "KUH", "DRÜCK", "LEER", "NEIN", "COUP", "NOCHMAL"},
		"   ":      {"LEER", "WAS", "KUH", "NOCHMAL", "NEIN", "DRÜCK", "OK", "JA", "NICHTS", "OKAY", "COUP", "Q", "   ", "O.K."},
		"LEER":     {"KUH", "OK", "Q", "O.K.", "LEER", "COUP", "DRÜCK", "NICHTS", "   ", "NEIN", "OKAY", "WAS", "JA", "NOCHMAL"},
		"NICHTS":   {"WAS", "OK", "Q", "O.K.", "JA", "LEER", "   ", "COUP", "OKAY", "NEIN", "KUH", "NOCHMAL", "NICHTS", "DRÜCK"},
		"JA":       {"Q", "OK", "WAS", "O.K.", "NOCHMAL", "NEIN", "COUP", "DRÜCK", "NICHTS", "JA", "OKAY", "LEER", "   ", "KUH"},
		"NEIN":     {"WAS", "NEIN", "OKAY", "NICHTS", "DRÜCK", "LEER", "O.K.", "   ", "Q", "NOCHMAL", "KUH", "JA", "COUP", "OK"},
		"WAS":      {"DRÜCK", "NICHTS", "OKAY", "NEIN", "Q", "JA", "OK", "   ", "COUP", "LEER", "WAS", "O.K.", "KUH", "NOCHMAL"},
		"OKAY":     {"OK", "OKAY", "NOCHMAL", "   ", "O.K.", "JA", "LEER", "NEIN", "WAS", "KUH", "COUP", "DRÜCK", "Q", "NICHTS"},
		"OK":       {"JA", "NICHTS", "DRÜCK", "COUP", "   ", "KUH", "NEIN", "OK", "O.K.", "OKAY", "WAS", "LEER", "Q", "NOCHMAL"},
		"O.K.":     {"LEER", "DRÜCK", "Q", "NEIN", "NICHTS", "COUP", "   ", "KUH", "OKAY", "O.K.", "OK", "NOCHMAL", "WAS", "JA"},
		"Q":        {"O.K.", "   ", "NOCHMAL", "JA", "WAS", "NICHTS", "KUH", "Q", "OKAY", "DRÜCK", "LEER", "COUP", "NEIN", "OK"},
		"KUH":      {"WAS", "   ", "LEER", "Q", "JA", "OKAY", "NOCHMAL", "COUP", "NEIN", "KUH", "NICHTS", "DRÜCK", "OK", "O.K."},
		"COUP":     {"OK", "O.K.", "JA", "DRÜCK", "COUP", "Q", "NICHTS", "WAS", "LEER", "OKAY", "NOCHMAL", "NEIN", "   ", "KUH"},
		"SOHN":     {"MOMENT", "SO EIN", "SO'N", "OH GOTT", "WAS?", "ZEH", "ZEHN", "WARTE", "C", "SOHN", "10", "DA STEHT", "ZEHEN", "CN"},
		"SO EIN":   {"SO'N", "WAS?", "DA STEHT", "ZEH", "C", "ZEHEN", "10", "WARTE", "SOHN", "CN", "OH GOTT", "MOMENT", "ZEHN", "SO EIN"},
		"SO'N":     {"10", "SO EIN", "ZEH", "SO'N", "WAS?", "ZEHN", "MOMENT", "CN", "OH GOTT", "SOHN", "C", "WARTE", "DA STEHT", "ZEHEN"},
		"OH GOTT":  {"SOHN", "OH GOTT", "ZEHN", "WAS?", "10", "SO EIN", "CN", "SO'N", "C", "ZEH", "MOMENT", "ZEHEN", "DA STEHT", "WARTE"},
		"ZEHN":     {"ZEHEN", "CN", "ZEHN", "ZEH", "C", "MOMENT", "SO'N", "WARTE", "OH GOTT", "DA STEHT", "WAS?", "10", "SO EIN", "SOHN"},
		"CN":       {"ZEH", "MOMENT", "WAS?", "C", "OH GOTT", "ZEHN", "10", "ZEHEN", "CN", "SOHN", "DA STEHT", "WARTE", "SO EIN", "SO'N"},
		"ZEH":      {"ZEH", "SO'N", "SO EIN", "SOHN", "ZEHEN", "WARTE", "10", "WAS?", "MOMENT", "DA STEHT", "OH GOTT", "ZEHN", "CN", "C"},
		"10":       {"ZEHN", "CN", "SO EIN", "OH GOTT", "WAS?", "10", "ZEHEN", "SOHN", "ZEH", "DA STEHT", "SO'N", "MOMENT", "WARTE", "C"},
		"C":        {"SOHN", "WARTE", "OH GOTT", "SO'N", "CN", "ZEHEN", "10", "DA STEHT", "SO EIN", "ZEH", "ZEHN", "WAS?", "C", "MOMENT"},
		"ZEHEN":    {"MOMENT", "ZEH", "WAS?", "C", "SO'N", "ZEHN", "OH GOTT", "WARTE", "DA STEHT", "SOHN", "CN", "SO EIN", "10", "ZEHEN"},
		"WAS?":     {"C", "ZEH", "10", "SO'N", "WARTE", "MOMENT", "WAS?", "DA STEHT", "ZEHEN", "SO EIN", "ZEHN", "OH GOTT", "CN", "SOHN"},
		"WARTE":    {"SO EIN", "CN", "ZEHEN", "10", "SOHN", "ZEHN", "MOMENT", "C", "OH GOTT", "WAS?", "WARTE", "ZEH", "SO'N", "DA STEHT"},
		"MOMENT":   {"SO EIN", "ZEHEN", "DA STEHT", "OH GOTT", "SOHN", "WARTE", "ZEH", "ZEHN", "MOMENT", "CN", "C", "WAS?", "SO'N", "10"},
		"DA STEHT": {"OH GOTT", "WAS?", "CN", "ZEHN", "WARTE", "ZEHEN", "10", "C", "ZEH", "SOHN", "DA STEHT", "MOMENT", "SO EIN", "SO'N"},
	}
	dangerWords := map[string]string{
		"OKAY":   "mit vier Buchstaben",
		"OK":     "mit zwei Buchstaben",
		"O.K.":   "jeweils mit Punkt",
		"NICHTS": "als Wort",
		"   ":    "als Freiraum",
		"LEER":   "wie die Stadt",
		"KUH":    "das Tier",
		"Q":      "der Buchstabe",
		"COUP":   "der französische Trick",
		"SO EIN": "mit Leerstelle",
		"SO'N":   "mit Apostroph",
		"SOHN":   "das Kind",
		"CN":     "Cäsar Nordpol",
		"10":     "zwei Ziffern",
		"ZEHN":   "ausgeschriebene Zahl",
		"ZEH":    "das einzelne Körperteil",
		"ZEHEN":  "das Körperteil im Plural",
		"CE":     "Cäsar Emil",
		"C":      "Cäsar",
		"ZÄH":    "wie Leder",
		"WAS":    "ohne Fragezeichen",
		"WAS?":   "mit Fragezeichen",
	}
	return &SixWordsModule{
		bigPositionMap: bigPositionMap,
		pushMap:        pushMap,
		dangerWords:    dangerWords,
	}
}

func (x *SixWordsModule) Reset() error {
	return nil
}

func (x *SixWordsModule) Solve() error {
	rounds := 5
	for i := 1; i <= rounds; i++ {
		fmt.Println("Runde", i)

		// phase 1
		bigWord, err := communication.ChooseOneWithMapper[string]("Großes Wort?", x.bigWords(), func(word string) string {
			return communication.SafeWord(word, x.dangerWords)
		})
		if err != nil {
			// abort module?
			if err == terminal.InterruptErr {
				return nil
			}
			log.Error().Err(err).Msg("could not get big word")
			i--
			continue
		}
		pos, ok := x.bigPositionMap[bigWord]
		if !ok {
			log.Error().Err(err).Msgf("could not find position for word %s", bigWord)
			i--
			continue
		}

		// phase 2
		lookupWord, err := communication.ChooseOneWithMapper[string](fmt.Sprintf("Das Wort %s lautet?", pos.ToString()), x.smallWords(), func(word string) string {
			return communication.SafeWord(word, x.dangerWords)
		})
		if err != nil {
			// abort module?
			if err == terminal.InterruptErr {
				return nil
			}
			log.Error().Err(err).Msg("could not get lookup word")
			i--
			continue
		}

		matchList, ok := x.pushMap[lookupWord]
		if !ok {
			log.Error().Err(err).Msgf("could not find word list for word %s", lookupWord)
			i--
			continue
		}
		x.tellMatchList(matchList, lookupWord)

		fmt.Println()
	}
	return nil
}

func (x *SixWordsModule) bigWords() []string {
	result := make([]string, 0, len(x.bigPositionMap))
	for key := range x.bigPositionMap {
		result = append(result, key)
	}
	return result
}

func (x *SixWordsModule) smallWords() []string {
	result := make([]string, 0, len(x.pushMap))
	for key := range x.pushMap {
		result = append(result, key)
	}
	return result
}

func (x SixWordsModule) tellMatchList(list []string, stopWord string) {
	list = x.simplifyPushList(list, stopWord)
	for _, word := range list {
		communication.Tell(communication.SafeWord(word, x.dangerWords))
	}
}

func (x SixWordsModule) simplifyPushList(list []string, stopWord string) []string {
	result := make([]string, 0)

	for _, word := range list {
		result = append(result, word)

		if word == stopWord {
			break
		}
	}

	return result
}

//go:generate go-enum -f=$GOFILE --marshal
// SixWordsPosition is an enumeration of positions
/*
ENUM(
topLeft
middleLeft
bottomLeft
topRight
middleRight
bottomRight
)
*/
type SixWordsPosition int

func (x SixWordsPosition) ToString() string {
	switch x {
	case SixWordsPositionTopLeft:
		return "oben links"
	case SixWordsPositionMiddleLeft:
		return "Mitte links"
	case SixWordsPositionBottomLeft:
		return "unten links"
	case SixWordsPositionTopRight:
		return "oben rechts"
	case SixWordsPositionMiddleRight:
		return "Mitte rechts"
	case SixWordsPositionBottomRight:
		return "unten rechts"
	default:
		return ""
	}
}
