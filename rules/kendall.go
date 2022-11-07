package rules

import (
	"gitlab.utc.fr/lagruesy/ia04/utils"
)

func KemenySWF(p Profile) []Alternative {
	a := p[0]
	n := len(a)
	index_list := [][]int{}

	// création et affichage de la première permutation
	perm := utils.FirstPermutation(n)
	index_list = append(index_list, perm)

	// itération et affichage de la permutation suivante
	perm, ok := utils.NextPermutation(perm)
	index_list = append(index_list, perm)
	for ok {
		perm, ok = utils.NextPermutation(perm)
		index_list = append(index_list, perm)
	}

	// on transforme les permutations (sous forme d'index) en Alternatives
	alt_perm := [][]Alternative{}
	for _, p := range index_list {
		alts := []Alternative{}
		for i := range p {
			alts = append(alts, a[p[i]])
		}
		alt_perm = append(alt_perm, alts)
	}

	// on calcule la distance par rapport au profil pour chacun des arrangements
	minDist := profileEditionDistance(alt_perm[0], p)
	bestOrder := alt_perm[0]
	for i := range alt_perm {
		profDist := profileEditionDistance(alt_perm[i], p)
		if profDist == 0 {
			return alt_perm[i]
		}
		if profDist < minDist {
			minDist = profDist
			bestOrder = alt_perm[i]
		}
	}

	return bestOrder
}

func KemenySCF(p Profile) Alternative {
	bestOrder := KemenySWF(p)
	return bestOrder[0]
}

func profileEditionDistance(alts []Alternative, p Profile) (dist int) {
	for _, a := range p {
		dist += editionDistance(alts, a)
	}
	return
}

func editionDistance(alts1 []Alternative, alts2 []Alternative) (dist int) {
	length := len(alts1)
	pairs1 := make([][]Alternative, 0)
	pairs2 := make([][]Alternative, 0)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			toAppend1 := []Alternative{alts1[i], alts1[j]}
			pairs1 = append(pairs1, toAppend1)

			toAppend2 := []Alternative{alts2[i], alts2[j]}
			pairs2 = append(pairs2, toAppend2)
		}
	}
	for i := 0; i < len(pairs1); i++ {
		contains := false
		for j := 0; j < len(pairs2); j++ {
			if pairs1[i][0] == pairs2[j][0] && pairs1[i][1] == pairs2[j][1] {
				contains = true
			}
		}
		if contains == false {
			dist++
		}
	}
	return
}
