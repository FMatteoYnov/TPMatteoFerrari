package tp

import (
	"sort"
)

// Ft_non_overlap : fonction pour trouver le plus petit nombre d'intervalles à retirer pour qu'il n'y ait pas de chevauchement
func Ft_non_overlap(intervals [][]int) int {
	// Si la liste est vide ou contient moins de 2 intervalles, il n'y a pas de chevauchement
	if len(intervals) == 0 {
		return 0
	}

	// Trier les intervalles par la fin de l'intervalle (croissant)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	// Initialiser les variables
	nonOverlapCount := 1       // Nombre d'intervalles non chevauchants (on inclut toujours le premier)
	lastEnd := intervals[0][1] // Fin de l'intervalle courant non chevauchant

	// Parcourir les intervalles à partir du second
	for i := 1; i < len(intervals); i++ {
		// Si l'intervalle courant ne chevauche pas l'intervalle précédent
		if intervals[i][0] >= lastEnd {
			nonOverlapCount++
			lastEnd = intervals[i][1] // Mettre à jour la fin de l'intervalle courant non chevauchant
		}
	}

	// Le nombre d'intervalles à supprimer est le nombre total d'intervalles moins le nombre d'intervalles non chevauchants
	return len(intervals) - nonOverlapCount
}
