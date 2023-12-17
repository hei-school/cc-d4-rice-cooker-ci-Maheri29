# frozen_string_literal: true

require 'tty-prompt'

def choisir_mode
  prompt = TTY::Prompt.new

  afficher_modes
  choix = prompt.select('Choisissez un mode de cuisson :', modes_disponibles)
  temps_cuisson = determiner_temps(choix, prompt)

  return unless temps_cuisson.positive?

  afficher_alertes(prompt, temps_cuisson)
end

def afficher_modes
  puts 'Modes disponibles :'
  puts '1. Riz Blanc'
  puts '2. Riz Complet'
  puts '3. Cuisson Vapeur'
  puts '4. Autre aliment'
end

def modes_disponibles
  ['Riz Blanc', 'Riz Complet', 'Cuisson Vapeur', 'Autre aliment']
end

def determiner_temps(choix, prompt)
  case choix
  when 'Riz Blanc', 'Riz Complet', 'Cuisson Vapeur'
    afficher_mode_selectionne(choix)
    2
  when 'Autre aliment'
    determiner_temps_autre_aliment(prompt)
  else
    afficher_erreur_choix
    0
  end
rescue StandardError => e
  gerer_erreur_temps(e)
end

def afficher_mode_selectionne(choix)
  puts "Mode #{choix} sélectionné"
end

def determiner_temps_autre_aliment(prompt)
  temps_personnalise = prompt.ask('Entrez le temps de cuisson en secondes pour l\'autre aliment : ').to_i
  temps_personnalise.positive? ? temps_personnalise : afficher_temps_invalide
rescue StandardError => e
  gerer_erreur_temps(e)
end

def afficher_temps_invalide
  puts 'Temps invalide.'
  0
end

def afficher_erreur_choix
  puts 'Choix non valide'
end

def afficher_alertes(prompt, temps_cuisson)
  puts 'Types d\'alertes disponibles :'
  alertes = ['Son', 'Lumières clignotantes']
  choix_alerte = prompt.select('Choisissez le type d\'alerte pour signaler la fin de la cuisson :', alertes)

  return unless alertes.include?(choix_alerte)

  alerte_selectionnee(prompt, choix_alerte, temps_cuisson)
end

def alerte_selectionnee(prompt, choix_alerte, temps_cuisson)
  threads = []
  threads << Thread.new do
    afficher_alerte(choix_alerte, temps_cuisson)
  end

  threads << Thread.new do
    choix_apres_cuisson = prompt.select('Que voulez-vous faire maintenant?', ['Éteindre', 'Maintenir au chaud'])
    traiter_choix_apres_cuisson(choix_apres_cuisson)
  end

  threads.each(&:join)
end

def afficher_alerte(choix_alerte, temps_cuisson)
  sleep temps_cuisson
  message = message_alerte(choix_alerte)
  puts message
end

def message_alerte(choix_alerte)
  if choix_alerte == 'Son'
    '*BIP*BIP*BIP* La cuisson est terminée !'
  else
    '*lumières clignotantes* La cuisson est terminée !'
  end
end

def traiter_choix_apres_cuisson(choix_apres_cuisson)
  case choix_apres_cuisson
  when 'Éteindre'
    puts 'Le rice cooker a été éteint.'
  when 'Maintenir au chaud'
    puts 'Le riz est maintenu au chaud.'
  else
    puts 'Choix non valide. Le rice cooker sera éteint par défaut.'
  end
end

def gerer_erreur_temps(erreur)
  puts "Une erreur s'est produite : #{erreur.message}"
  puts 'Veuillez réessayer ou contacter le support.'
  0
end

choisir_mode
