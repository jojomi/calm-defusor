package modules

import (
	"fmt"
	"github.com/jojomi/calm-defusor/communication"
	"strings"
	"unicode/utf8"
)

// How to improve?
// entropy analysis to skip a character or even optimize the order by entropy completely

type PasswordModule struct {
	words   []string
	letters []*LetterState
}

func (p *PasswordModule) Name() string {
	return "Passwort"
}

func (p *PasswordModule) String() string {
	return p.Name()
}

func NewPasswordModule() *PasswordModule {
	words := []string{
		"angst",
		"atmen",
		"beten",
		"bombe",
		"danke",
		"draht",
		"druck",
		"drück",
		"farbe",
		"fehlt",
		"ferse",
		"kabel",
		"knall",
		"knapp",
		"knopf",
		"leere",
		"legal",
		"lehre",
		"mathe",
		"matte",
		"panik",
		"pieps",
		"rauch",
		"ruhig",
		"saite",
		"sehne",
		"seite",
		"sende",
		"strom",
		"super",
		"timer",
		"übrig",
		"verse",
		"warte",
		"zange",
	}

	module := &PasswordModule{
		words: words,
	}
	module.Reset()

	return module
}

func (p *PasswordModule) Reset() error {
	length := len(p.words[0])
	letters := make([]*LetterState, length)
	for i := 0; i < length; i++ {
		letters[i] = NewLetterState()
	}
	p.letters = letters
	return nil
}

func (p *PasswordModule) Solve() error {
	wordLength := p.getWordLength()
	for letterIndex := 0; letterIndex < wordLength; letterIndex++ {
		communication.AskPrintf("Lies nacheinander alle möglichen Buchstaben für die %d. Stelle vor!", letterIndex+1)
		fmt.Printf("\n> ")

		var (
			s           string
			startLetter rune
			finished    = false
		)
		letterState := p.letters[letterIndex]
		for {
			_, err := fmt.Scanf("%s", &s)
			if err != nil {
				continue
			}

			s = strings.TrimSpace(s)
			for _, r := range s {
				if startLetter == r || r == '-' {
					finished = true
				}
				if startLetter == 0 {
					startLetter = r
				}
				letterState.AddPossibleLetter(string(r))
			}
			fmt.Printf("%d. > %s\n", letterIndex+1, strings.Join(letterState.letters, " "))

			// multichar mode or single char mode?
			if finished || utf8.RuneCountInString(s) > 1 {
				letterState.SetAllKnown()
				break
			}
		}
		if solution := p.getSolution(); solution != nil {
			communication.Tellf("Die Lösung ist \"%s\".\n", communication.SafeWord(*solution))
			break
		} else {
			possibleWords := p.getPossibleWords()
			fmt.Printf("Noch %d Möglichkeiten: %s\n\n", len(possibleWords), strings.Join(possibleWords, " "))
		}
	}
	return nil
}

func (p PasswordModule) getWordLength() int {
	if len(p.words) < 1 {
		return 0
	}
	return len(p.words[0])
}

func (p PasswordModule) getSolution() *string {
	words := p.getPossibleWords()
	if len(words) == 1 {
		return &words[0]
	}
	return nil
}

func (p PasswordModule) getPossibleWords() []string {
	var result []string
	for _, word := range p.words {
		if !p.isWordPossible(word) {
			continue
		}
		result = append(result, word)
	}
	return result
}

func (p PasswordModule) isWordPossible(word string) bool {
	var ls *LetterState
	letterIndex := 0
	for _, character := range word {
		letterIndex++
		ls = p.letters[letterIndex-1]
		if !ls.IsAllKnown() {
			continue
		}
		if !ls.HasPossibleLetter(string(character)) {
			return false
		}
	}
	return true
}

type LetterState struct {
	letters  []string
	allKnown bool
}

func NewLetterState() *LetterState {
	return &LetterState{}
}

func (l LetterState) HasPossibleLetter(letter string) bool {
	for _, le := range l.letters {
		if le == letter {
			return true
		}
	}
	return false
}

func (l *LetterState) AddPossibleLetter(letter string) {
	l.letters = append(l.letters, letter)
}

func (l *LetterState) SetAllKnown() {
	l.allKnown = true
}

func (l LetterState) IsAllKnown() bool {
	return l.allKnown
}
