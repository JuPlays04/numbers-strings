package natural

// DigitString10 erwartet eine Ziffer als int.
// Die Funktion gibt den zugehörigen String zurück, wie er üblicherweise
// an der Zehner-Stelle einer Zahl >= 21 vorkommen würde.
//
// Beispiel: Für 3 soll der String "dreißig" geliefert werden.
//
// Anmerkung:
// Dies ist eine Hilfsfunktion, die genutzt werden soll,
// um den Gesamt-String einer Zahl zusammenzusetzen.
func DigitString10(digit int) string {
	if digit == 0 {
		return ""
	}
	if digit == 1 {
		return ""
	}
	if digit == 2 {
		return "zwanzig"
	}
	if digit == 3 {
		return "dreißig"
	}
	if digit == 4 {
		return "vierzig"
	}
	if digit == 5 {
		return "fünfzig"
	}
	if digit == 6 {
		return "sechzig"
	}
	if digit == 7 {
		return "siebzig"
	}
	if digit == 8 {
		return "achtzig"
	}
	if digit == 9 {
		return "neunzig"
	}
	return ""
}
