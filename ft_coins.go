package tp

import (
	"math"
)

func Ft_coin(coins []int, amount int) int {
	// Créez un tableau dp de taille amount + 1 pour stocker le nombre minimum de pièces pour chaque valeur
	dp := make([]int, amount+1)

	// Initialisez chaque élément de dp à une valeur infinie (ou une valeur très grande)
	for i := range dp {
		dp[i] = math.MaxInt32
	}

	// Le montant 0 peut être atteint avec 0 pièce
	dp[0] = 0

	// Parcourez chaque montant de 1 à amount
	for i := 1; i <= amount; i++ {
		// Pour chaque pièce dans coins, vérifiez si elle peut être utilisée
		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-coin]+1)))
			}
		}
	}

	// Si le montant n'est pas atteignable, retournez -1
	if dp[amount] == math.MaxInt32 {
		return -1
	}

	// Sinon, retournez le nombre minimum de pièces nécessaires
	return dp[amount]
}
