package main

import (
	"fmt"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
	WhiteColor   = "\033[1;37m%s\033[0m"
)

func InitGrid(grille *[3][3]string) {

	// Permet de remplir la grille de "[]"
	for s := range grille {
		for i := range grille[s] {
			grille[s][i] = "[ ]"
		}
	}

	// Permet de print ligne par ligne
	for s := range grille {
		fmt.Println(grille[s])
	}

}

func PrintGrid(grille [3][3]string) {
	for s := range grille {
		fmt.Println(grille[s])
	}
}

func CaseIsFull(position int, grille [3][3]string) bool {
	switch position {
	case 1:
		if grille[0][0] != "[ ]" {
			return true
		}
	case 2:
		if grille[0][1] != "[ ]" {
			return true
		}
	case 3:
		if grille[0][2] != "[ ]" {
			return true
		}
	case 4:
		if grille[1][0] != "[ ]" {
			return true
		}
	case 5:
		if grille[1][1] != "[ ]" {
			return true
		}
	case 6:
		if grille[1][2] != "[ ]" {
			return true
		}
	case 7:
		if grille[2][0] != "[ ]" {
			return true
		}
	case 8:
		if grille[2][1] != "[ ]" {
			return true
		}
	case 9:
		if grille[2][2] != "[ ]" {
			return true
		}
	}
	return false
}

func CheckWin(grille [3][3]string) bool {
	// Check ligne
	for s := range grille {
		if grille[s][0] == grille[s][1] && grille[s][1] == grille[s][2] && grille[s][0] != "[ ]" {
			return true
		}
	}
	// Check colonne
	for i := range grille {
		if grille[0][i] == grille[1][i] && grille[1][i] == grille[2][i] && grille[0][i] != "[ ]" {
			return true
		}
	}
	// Check diagonale
	if grille[0][0] == grille[1][1] && grille[1][1] == grille[2][2] && grille[0][0] != "[ ]" {
		return true
	}
	if grille[0][2] == grille[1][1] && grille[1][1] == grille[2][0] && grille[0][2] != "[ ]" {
		return true
	}
	return false
}

func AddMorpion(joueur string, position int, choix string, grille *[3][3]string) {
	if AllCaseIsFull(*grille) {
		fmt.Printf(ErrorColor, "Toutes les cases sont prises ! Match nul !\n")
		return
	}
	if CaseIsFull(position, *grille) {
		fmt.Printf(ErrorColor, "La case est déjà prise !\n")
		fmt.Scanln(&position)
	}
	switch position {
	case 1:
		grille[0][0] = "[" + choix + "]"
	case 2:
		grille[0][1] = "[" + choix + "]"
	case 3:
		grille[0][2] = "[" + choix + "]"
	case 4:
		grille[1][0] = "[" + choix + "]"
	case 5:
		grille[1][1] = "[" + choix + "]"
	case 6:
		grille[1][2] = "[" + choix + "]"
	case 7:
		grille[2][0] = "[" + choix + "]"
	case 8:
		grille[2][1] = "[" + choix + "]"
	case 9:
		grille[2][2] = "[" + choix + "]"
	}
	PrintGrid(*grille)
}

func AllCaseIsFull(grille [3][3]string) bool {
	for s := range grille {
		for i := range grille[s] {
			if grille[s][i] == "[ ]" {
				return false
			}
		}
	}
	return true
}

func main() {
	count := 0
	fmt.Printf(InfoColor, "Bienvenue dans le Jeux du")
	fmt.Printf(NoticeColor, " Morpion\n")
	var joueur1 string
	var joueur2 string
	fmt.Printf(InfoColor, "Joueur 1, entrez votre nom : ")
	fmt.Scanln(&joueur1)
	fmt.Printf(InfoColor, "Joueur 2, entrez votre nom : ")
	fmt.Scanln(&joueur2)
	if joueur1 == joueur2 {
		fmt.Printf(ErrorColor, "Vous ne pouvez pas avoir le même nom !\n")
		return
	}
	if len(joueur1) > 20 || len(joueur2) > 20 {
		fmt.Printf(ErrorColor, "Votre nom est trop long !\n")
		return
	}
	if len(joueur1) < 3 || len(joueur2) < 3 {
		fmt.Printf(ErrorColor, "Votre nom est trop court !\n")
		return
	}
	if len(joueur1) == 0 || len(joueur2) == 0 {
		fmt.Printf(ErrorColor, "Vous devez entrer un nom !\n")
		return
	}
	fmt.Printf(InfoColor, "Bienvenue: "+joueur1+" et "+joueur2+"\n")
	var grille [3][3]string
	InitGrid(&grille)

	var choixj1 string
	var choixj2 string
	fmt.Printf(InfoColor, "Joueur 1, choisissez entre X ou O : ")
	fmt.Scanln(&choixj1)
	if choixj1 != "X" && choixj1 != "O" {
		fmt.Printf(ErrorColor, "Vous devez choisir entre X ou O !\n")
		fmt.Scanln(&choixj1)
	}
	fmt.Printf(InfoColor, "Joueur 2, choisissez entre X ou O : ")
	fmt.Scanln(&choixj2)
	if choixj2 != "X" && choixj2 != "O" {
		fmt.Printf(ErrorColor, "Vous devez choisir entre X ou O !\n")
		fmt.Scanln(&choixj2)
		if choixj2 == choixj1 {
			fmt.Printf(ErrorColor, "Vous ne pouvez pas avoir le même choix !\n")
			return
		}

	}
	if choixj1 == choixj2 || choixj1 == "x" && choixj2 == "x" || choixj1 == "o" && choixj2 == "o" || choixj1 == "X" && choixj2 == "X" || choixj1 == "O" && choixj2 == "O" {
		fmt.Printf(ErrorColor, "Vous ne pouvez pas avoir le même choix !\n")
		fmt.Scanln(&choixj2)
	}

	for AllCaseIsFull(grille) == false {
		var positionj1 int
		if count == 9 || count == 18 {
			fmt.Printf(ErrorColor, "Toutes les cases sont prises ! Match nul !\n")
			return
		}
		fmt.Printf(InfoColor, "Joueur 1, choisissez une position entre 1 et 9 : ")
		fmt.Scanln(&positionj1)
		if positionj1 < 1 || positionj1 > 9 {
			fmt.Printf(ErrorColor, "Vous devez choisir une position entre 1 et 9 !\n")
			fmt.Scanln(&positionj1)
		}
		if AllCaseIsFull(grille) {
			fmt.Printf(ErrorColor, "Toutes les cases sont prises ! Match nul !\n")
			return
		}
		if CheckWin(grille) {
			fmt.Printf(NoticeColor, "Le joueur "+joueur2+" a gagné !\n")
			return
		}
		if AllCaseIsFull(grille) {
			fmt.Printf(ErrorColor, "Toutes les cases sont prises ! Match nul !\n")
			return
			break
		}
		AddMorpion(joueur1, positionj1, choixj1, &grille)
		count = count + 1
		var positionj2 int
		fmt.Printf(InfoColor, "Joueur 2, choisissez une position entre 1 et 9 : ")
		fmt.Scanln(&positionj2)
		if positionj2 < 1 || positionj2 > 9 {
			fmt.Printf(ErrorColor, "Vous devez choisir une position entre 1 et 9 !\n")
			fmt.Scanln(&positionj2)
		}
		if AllCaseIsFull(grille) {
			fmt.Printf(ErrorColor, "Toutes les cases sont prises ! Match nul !\n")
			return
			break
		}
		if CheckWin(grille) {
			fmt.Printf(NoticeColor, "Le joueur "+joueur1+" a gagné !\n")
			return
		}
		AddMorpion(joueur2, positionj2, choixj2, &grille)
		count = count + 1
		if AllCaseIsFull(grille) {
			fmt.Printf(ErrorColor, "Toutes les cases sont prises ! Match nul !\n")
			break
		}
	}
}
