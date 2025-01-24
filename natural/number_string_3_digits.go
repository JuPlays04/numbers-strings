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

	} else if 9<x && x <20
	{
		if einer == 0 && zehner == 1 && hunderter == 0 {
			return "zehn"
		}
		if einer == 1 && zehner == 1 && hunderter == 0 {
			return "elf"
		}
		if einer == 2 && zehner == 1 && hunderter == 0 {
			return "zwölf"
		}
		if einer == 3 && zehner == 1 && hunderter == 0 {
			return "dreizehn"
		}
		if einer == 4 && zehner == 1 && hunderter == 0 {
			return "vierzehn"
		}
		if einer == 5 && zehner == 1 && hunderter == 0 {
			return "fünfzehn"
		}
		if einer == 6 && zehner == 1 && hunderter == 0 {
			return "sechzehn"
		}
		if einer == 7 && zehner == 1 && hunderter == 0 {
			return "siebzehn"
		}
		if einer == 8 && zehner == 1 && hunderter == 0 {
			return "achtzehn"
		}
		if einer == 9 && zehner == 1 && hunderter == 0 {
			return "neunzehn"
		}
	} else if x < 100 {
		if einer == 0 && zehner == 2 && hunderter == 0 {
			return "zwanzig"
		}
		if einer == 1 && zehner == 2 && hunderter == 0 {
			return "einundzwanzig"
		}
		if einer == 2 && zehner == 2 && hunderter == 0 {
			return "zweiundzwanzig"
		}
		if einer == 3 && zehner == 2 && hunderter == 0 {
			return "dreiundzwanzig"
		}
		if einer == 4 && zehner == 2 && hunderter == 0 {
			return "vierundzwanzig"
		}
		if einer == 5 && zehner == 2 && hunderter == 0 {
			return "fünfundzwanzig"
		} 
