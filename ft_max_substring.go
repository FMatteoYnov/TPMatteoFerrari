package tp

// Ft_max_substring : fonction pour trouver la longueur de la plus grande sous-chaîne avec des caractères uniques
func Ft_max_substring(s string) int {
	// Créer une map pour stocker l'index des caractères déjà rencontrés
	charIndex := make(map[byte]int)
	maxLength := 0 // La longueur maximale de la sous-chaîne sans répétition
	start := 0     // Le début de la fenêtre actuelle

	// Parcourir chaque caractère de la chaîne
	for i := 0; i < len(s); i++ {
		char := s[i]

		// Si le caractère a déjà été rencontré et qu'il est dans la fenêtre actuelle
		if index, found := charIndex[char]; found && index >= start {
			// Déplacer le début de la fenêtre juste après la position précédente de ce caractère
			start = index + 1
		}

		// Mettre à jour la position actuelle du caractère dans la map
		charIndex[char] = i

		// Calculer la longueur de la sous-chaîne actuelle et mettre à jour maxLength si nécessaire
		currentLength := i - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	// Retourner la longueur maximale de la sous-chaîne sans répétition
	return maxLength
}
