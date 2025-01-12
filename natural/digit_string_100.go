package natural

// DigitString100 erwartet eine Ziffer als int.
// Die Funktion gibt den zugehörigen String zurück, wie er üblicherweise
// an der Hunderter-Stelle einer Zahl >= 21 vorkommen würde.
//
// Beispiel: Für 3 soll der String "dreihundert" geliefert werden.
//
// Anmerkung:
// Dies ist eine Hilfsfunktion, die genutzt werden soll,
// um den Gesamt-String einer Zahl zusammenzusetzen.
func DigitString100(digit int) string {
	if digit == 0 {
		return ""
	}
	if digit == 1 {
		return "einhundert"
	}
	if digit == 2 {
		return "zweihundert"
	}
	if digit == 3 {
		return "dreihundert"
	}
	if digit == 4 {
		return "vierhundert"
	}
	if digit == 5 {
		return "fünfhundert"
	}
	if digit == 6 {
		return "sechshundert"
	}
	if digit == 7 {
		return "siebenhundert"
	}
	if digit == 8 {
		return "achthundert"
	}
	if digit == 9 {
		return "neunhundert"
	}

	return ""

}
