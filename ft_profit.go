package tp

// Ft_profit : fonction pour calculer le plus grand bénéfice possible
func Ft_profit(prices []int) int {
	// Si le tableau de prix est vide ou ne contient qu'un élément, le bénéfice est 0
	if len(prices) == 0 {
		return 0
	}

	// Initialiser les variables
	minPrice := prices[0] // Le prix minimum rencontré jusqu'ici
	maxProfit := 0        // Le bénéfice maximum

	// Parcourir les prix à partir du deuxième jour
	for i := 1; i < len(prices); i++ {
		// Si un prix plus bas est trouvé, mettre à jour minPrice
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			// Calculer le bénéfice si on vend au jour i après avoir acheté au jour où le prix était minPrice
			profit := prices[i] - minPrice
			// Mettre à jour le bénéfice maximum si ce bénéfice est plus élevé
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	// Retourner le bénéfice maximum trouvé
	return maxProfit
}
