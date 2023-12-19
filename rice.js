const prompt = require('prompt-sync')();

function choisirMode() {
    console.log('Modes disponibles :');
    console.log('1. Riz Blanc');
    console.log('2. Riz Complet');
    console.log('3. Cuisson Vapeur');
    console.log('4. Autre aliment');

    const choix = prompt('Choisissez un mode de cuisson (1/2/3/4) : ');

    let tempsCuisson = 0;

    switch (choix) {
        case '1':
            console.log('Mode Riz Blanc sélectionné');
            tempsCuisson = 2;
            break;
        case '2':
            console.log('Mode Riz Complet sélectionné');
            tempsCuisson = 2;
        case '3':
            console.log('Mode Cuisson Vapeur sélectionné');
            tempsCuisson = 2;
            break;
        case '4':
            const tempsPersonnalise = parseInt(prompt('Entrez le temps de cuisson en secondes pour l\'autre aliment : '));
            if (!isNaN(tempsPersonnalise) && tempsPersonnalise > 0) {
                console.log(`Mode Autre Aliment sélectionné - Cuisson pendant ${tempsPersonnalise} secondes.`);
                tempsCuisson = tempsPersonnalise;
            } else {
                console.log('Temps invalide.');
            }
            break;
        default:
            console.log('Choix non valide');
            break;
    }

    if (tempsCuisson > 0) {
        console.log('Types d\'alertes disponibles :');
        console.log('1. Son');
        console.log('2. Lumières clignotantes');

        const choixAlerte = prompt('Choisissez le type d\'alerte pour signaler la fin de la cuisson (1/2) : ');

        console.log(`La cuisson se déroulera pendant ${tempsCuisson} secondes.`);

        switch (choixAlerte) {
            case '1':
                setTimeout(() => {
                    console.log('*BIP*BIP*BIP* La cuisson est terminée !');
                }, tempsCuisson * 1000);
                break;
            case '2':
                setTimeout(() => {
                    console.log('*lumières clignotantes* La cuisson est terminée !');
                }, tempsCuisson * 1000);
                break;
            default:
                console.log('Choix non valide. Aucune alerte ne sera déclenchée.');
                break;
        }

        setTimeout(() => {
            const choixApresCuisson = prompt('Que voulez-vous faire maintenant? (1. Éteindre / 2. Maintenir au chaud) : ');

            switch (choixApresCuisson) {
                case '1':
                    console.log('Le rice cooker a été éteint.');
                    break;
                case '2':
                    console.log('Le riz est maintenu au chaud.');
                    break;
                default:
                    console.log('Choix non valide. Le rice cooker sera éteint par défaut.');
                    break;
            }
        }, tempsCuisson * 1000);
    }
}

choisirMode();
module.exports = choisirMode;