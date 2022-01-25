package modules

import (
	"fmt"
	"github.com/jojomi/calm-defusor/communication"
	"github.com/rs/zerolog/log"
	"regexp"
	"strings"
)

type MorseModule struct {
	morseLetters []string
}

func (x MorseModule) Name() string {
	return "Morsecode"
}

func (x MorseModule) String() string {
	return x.Name()
}

func NewMorseModule() *MorseModule {
	return &MorseModule{}
}

func (x *MorseModule) Reset() error {
	x.morseLetters = []string{}
	return nil
}

func (x *MorseModule) Solve() error {
	fmt.Println("Buchstaben mit l (lang) und k (kurz) eingeben.")
	fmt.Println()
	communication.Tell("Lies Buchstaben einzeln vor, die große Pause zum Wortanfang ist egal.")
	fmt.Println()

	var (
		str string
		err error
	)

	for {
		str, err = communication.AskString("Nächster Buchstabe in Morsecode:")
		if err != nil {
			log.Error().Err(err).Msg("")
			continue
		}

		x.morseLetters = append(x.morseLetters, strings.ToLower(str))

		// solved?
		if x.isSolved() {
			word := x.getPossibleTexts()[0]
			fmt.Printf("Lösungswort: %s\n", word)
			communication.Tellf("Die Lösungsfrequenz ist \"%d MHz\".", x.mapTextToFreq(word))
			fmt.Println()
			return nil
		} else {
			x.printState()
			fmt.Println()
		}

		// valid?
		_, err = x.mapMorseToString(str)
		if err != nil {
			communication.Tell("Buchstabe nicht erkannt, weitermachen!")
			fmt.Println()
			fmt.Println()
			continue
		}
	}
	return nil
}

func (x MorseModule) printState() {
	fmt.Print("Aktuelles Teilwort: ")
	fmt.Println(x.mapMorseListToString(x.morseLetters))
	possibleTexts := x.getPossibleTexts()
	fmt.Printf("Mögliche Worte (%d): %s\n\n", len(possibleTexts), strings.Join(possibleTexts, ", "))
}

func (x *MorseModule) mapMorseListToString(morse []string) string {
	var (
		v   string
		err error
		sb  strings.Builder
	)
	for _, l := range morse {
		v, err = x.mapMorseToString(l)
		if err != nil {
			sb.WriteString(".")
		}
		sb.WriteString(v)
	}
	return sb.String()
}

func (x *MorseModule) mapMorseToString(morse string) (string, error) {
	m := map[string]string{
		"kl":     "a",
		"lkkk":   "b",
		"lklk":   "c",
		"lkk":    "d",
		"k":      "e",
		"kklk":   "f",
		"llk":    "g",
		"kkkk":   "h",
		"kk":     "i",
		"klll":   "j",
		"lkl":    "k",
		"klkk":   "l",
		"ll":     "m",
		"lk":     "n",
		"lll":    "o",
		"kllk":   "p",
		"llkl":   "q",
		"klk":    "r",
		"kkk":    "s",
		"l":      "t",
		"kkl":    "u",
		"kkkl":   "v",
		"kll":    "w",
		"lkkl":   "x",
		"lkll":   "y",
		"llkk":   "z",
		"klkl":   "ä",
		"lllk":   "ö",
		"kkll":   "ü",
		"kkkkkk": "ß",
		"llll":   "ch",
	}
	val, found := m[morse]

	if !found {
		return "", fmt.Errorf("no match for input %s", morse)
	}
	return val, nil
}

func (x *MorseModule) isSolved() bool {
	return len(x.getPossibleTexts()) == 1
}

func (x *MorseModule) getPossibleTexts() []string {
	var goodStrings []string

	for _, text := range x.getTexts() {
		if !x.isPossible(text) {
			continue
		}
		goodStrings = append(goodStrings, text)
	}

	return goodStrings
}

func (x *MorseModule) getTextFreqMap() map[string]int {
	return map[string]int{
		"halle":  3505,
		"hallo":  3515,
		"lokal":  3522,
		"steak":  3532,
		"stück":  3535,
		"speck":  3542,
		"bistro": 3545,
		"robust": 3522,
		"säbel":  3555,
		"sülze":  3565,
		"sektor": 3572,
		"vektor": 3575,
		"bravo":  3582,
		"kobra":  3592,
		"bombe":  3595,
		"süden":  3600,
	}
}

func (x *MorseModule) getTexts() []string {
	freqMap := x.getTextFreqMap()
	texts := make([]string, len(freqMap))

	index := 0
	for key := range freqMap {
		texts[index] = key
		index++
	}

	return texts
}

func (x *MorseModule) mapTextToFreq(text string) int {
	return x.getTextFreqMap()[text]
}

func (x MorseModule) isPossible(text string) bool {
	if len(x.morseLetters) == 0 {
		return true
	}
	r := regexp.MustCompile(x.mapMorseListToString(x.morseLetters))
	return r.FindString(text+text) != ""
}
