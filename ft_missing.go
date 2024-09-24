package tp

// FindMissingNumber : fonction pour trouver le nombre manquant dans la séquence [0, n]
func FindMissingNumber(nums []int) int {
	// Taille de la séquence (longueur de nums + 1)
	n := len(nums)

	// Calcul de la somme attendue des nombres de 0 à n (utilisation de la formule de la somme des entiers)
	expectedSum := n * (n + 1) / 2

	// Calcul de la somme actuelle des éléments dans le tableau nums
	actualSum := 0
	for _, num := range nums {
		actualSum += num
	}

	// Le nombre manquant est la différence entre la somme attendue et la somme actuelle
	return expectedSum - actualSum
}
