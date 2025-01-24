package natural

// NumberStringBelow20 erwartet eine Zahl < 20 und liefert den zugehörigen String.
func NumberStringBelow20(n int) string {
	if n >= 20 {
		return ""
	}

	Einer := n % 10
	Zehner := n / 10

	StEins := ""
	switch Einer {
	case 1:
		StEins = "eins"
	case 2:
		StEins = "zwei"
	case 3:
		StEins = "drei"
	case 4:
		StEins = "vier"
	case 5:
		StEins = "fünf"
	case 6:
		StEins = "sechs"
	case 7:
		StEins = "sieben"
	case 8:
		StEins = "acht"
	case 9:
		StEins = "neun"
	}

	switch n {
	case 11:
		return "elf"
	case 12:
		return "zwölf"
	case 16:
		return "sechzehn"
	case 17:
		return "siebzehn"
	}

	if Zehner == 1 {
		return StEins
	}

	return StEins
}
