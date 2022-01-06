package communication

import "fmt"

func SafeWord(input string) string {
	dangerousWords := map[string]string{
		"mathe": "wie Mathematik",
		"matte": "wie Gymnastikmatte",
		"saite": "wie bei der Geige",
		"seite": "wie beim Würfel",
		"leere": "wie das Vakuum",
		"lehre": "wie die Ausbildung",
		"druck": "mit U wie Ullrich",
		"drück": "mit Ü wie Überschall",
		"ferse": "wie am Fuß",
		"verse": "wie im Gedicht",
		"knapp": "wie eng",
		"sehne": "wie Achillessehne",
		"sende": "wie verschicke",
	}

	if explanation, ok := dangerousWords[input]; ok {
		return fmt.Sprintf("%s (%s)", input, explanation)
	}
	return input
}
