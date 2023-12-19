package main

import (
	"fmt"
	"time"
)

type Alerte interface {
	Declencher()
}

type SonAlerte struct{}

func (a SonAlerte) Declencher() {
	fmt.Println("*BIP*BIP*BIP* La cuisson est terminée !")
}

type LumiereAlerte struct{}

func (a LumiereAlerte) Declencher() {
	fmt.Println("*lumières clignotantes* La cuisson est terminée !")
}

func attendreCuisson(tempsCuisson int, alerte Alerte) {
	time.Sleep(time.Duration(tempsCuisson) * time.Second)
	alerte.Declencher()
}

type RiceCooker struct {
	Alerte Alerte
}

func (rc *RiceCooker) ChoisirMode() {
	fmt.Println("Modes disponibles :")
	fmt.Println("1. Riz Blanc")
	fmt.Println("2. Riz Complet")
	fmt.Println("3. Cuisson Vapeur")
	fmt.Println("4. Autre aliment")

	var choix string
	fmt.Print("Choisissez un mode de cuisson (1/2/3/4) : ")
	fmt.Scanln(&choix)

	var tempsCuisson int

	switch choix {
	case "1", "2", "3":
		fmt.Printf("Mode sélectionné : %s\n", choix)
		tempsCuisson = 2
	case "4":
		fmt.Print("Entrez le temps de cuisson en secondes pour l'autre aliment : ")
		fmt.Scanln(&tempsCuisson)
		if tempsCuisson <= 0 {
			fmt.Println("Temps invalide. La cuisson ne sera pas effectuée.")
			return
		}
		fmt.Printf("Mode Autre Aliment sélectionné - Cuisson pendant %d secondes.\n", tempsCuisson)
	default:
		fmt.Println("Choix non valide. La cuisson ne sera pas effectuée.")
		return
	}

	if tempsCuisson > 0 {
		fmt.Println("Types d'alertes disponibles :")
		fmt.Println("1. Son")
		fmt.Println("2. Lumières clignotantes")

		var choixAlerte string
		fmt.Print("Choisissez le type d'alerte pour signaler la fin de la cuisson (1/2) : ")
		fmt.Scanln(&choixAlerte)

		fmt.Printf("La cuisson se déroulera pendant %d secondes.\n", tempsCuisson)

		switch choixAlerte {
		case "1", "2":
			go attendreCuisson(tempsCuisson, rc.Alerte)
		default:
			fmt.Println("Choix non valide. Aucune alerte ne sera déclenchée.")
		}

		var choixApresCuisson string
		fmt.Print("Que voulez-vous faire maintenant? (1. Éteindre / 2. Maintenir au chaud) : ")
		fmt.Scanln(&choixApresCuisson)

		switch choixApresCuisson {
		case "1":
			fmt.Println("Le rice cooker a été éteint.")
		case "2":
			fmt.Println("Le riz est maintenu au chaud.")
		default:
			fmt.Println("Choix non valide. Le rice cooker sera éteint par défaut.")
		}
	}
}

func main() {
	riceCooker := RiceCooker{
		Alerte: SonAlerte{},
	}

	riceCooker.ChoisirMode()
}
