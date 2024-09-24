package tp

import (
	"math"
)

// Ft_min_window : fonction pour trouver la plus petite sous-chaîne dans 's' contenant tous les caractères de 't'
func Ft_min_window(s string, t string) string {
	if len(t) == 0 || len(s) == 0 {
		return ""
	}

	// Fréquences des caractères dans t
	tFreq := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tFreq[t[i]]++
	}

	// Variables pour la fenêtre glissante
	start, minStart, minLen := 0, 0, math.MaxInt32
	count := 0
	windowFreq := make(map[byte]int)

	// Parcourir la chaîne s avec une fenêtre glissante
	for end := 0; end < len(s); end++ {
		char := s[end]

		// Si le caractère est dans t, on le compte
		if tFreq[char] > 0 {
			windowFreq[char]++
			if windowFreq[char] <= tFreq[char] {
				count++
			}
		}

		// Dès que la fenêtre contient tous les caractères de t
		for count == len(t) {
			// Mettre à jour la longueur minimale si une fenêtre plus petite est trouvée
			if end-start+1 < minLen {
				minLen = end - start + 1
				minStart = start
			}

			// Réduire la fenêtre depuis le début
			startChar := s[start]
			if tFreq[startChar] > 0 {
				windowFreq[startChar]--
				if windowFreq[startChar] < tFreq[startChar] {
					count--
				}
			}
			start++
		}
	}

	// Si une sous-chaîne valide a été trouvée, retourner le résultat
	if minLen != math.MaxInt32 {
		return s[minStart : minStart+minLen]
	}

	// Si aucune sous-chaîne n'est trouvée
	return ""
}
