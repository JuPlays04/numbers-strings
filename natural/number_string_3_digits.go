package natural

// NumberString3Digits erwartet eine Zahl 0 <= n <= 999 und liefert den zugehörigen String.
func NumberString3Digits(n int) string {
	x := n
	einer := n % 10
	n = n / 10
	zehner := n % 10
	n = n / 10
	hunderter := n % 10
	if x < 10 {
		if einer == 0 && zehner == 0 && hunderter == 0 {
			return "null"
		}
		if einer == 1 && zehner == 0 && hunderter == 0 {
			return "eins"
		}
		if einer == 2 && zehner == 0 && hunderter == 0 {
			return "zwei"
		}
		if einer == 3 && zehner == 0 && hunderter == 0 {
			return "drei"
		}
		if einer == 4 && zehner == 0 && hunderter == 0 {
			return "vier"
		}
		if einer == 5 {
			return "fünf"
		}
		if einer == 6 {
			return "sechs"
		}
		if einer == 7 {
			return "sieben"
		}
		if einer == 8 {
			return "acht"
		}
		if einer == 9 {
			return "neun"
		}

	} if 9<x && x <20
}
