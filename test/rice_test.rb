# frozen_string_literal: true

require_relative '../rice'
require 'rspec'
require 'tty-prompt'

RSpec.describe 'VotreApplication' do
    describe '#choisir_mode' do
      it 'affiche un message de choix non valide' do
        allow_any_instance_of(TTY::Prompt).to receive(:select).and_return('Autre aliment')
        allow_any_instance_of(TTY::Prompt).to receive(:ask).and_return('po')
        allow_any_instance_of(TTY::Prompt).to receive(:select).and_return('Son')
        allow_any_instance_of(TTY::Prompt).to receive(:select).and_return('Éteindre')
        expect { choisir_mode }.to output(/Choix non valide/).to_stdout
      end
    end

  describe '#determiner_temps' do
    it 'retourne le temps de cuisson pour les modes prédéfinis' do
      prompt = TTY::Prompt.new
      expect(determiner_temps('Riz Blanc', prompt)).to eq(2)
      expect(determiner_temps('Riz Complet', prompt)).to eq(2)
      expect(determiner_temps('Cuisson Vapeur', prompt)).to eq(2)
    end

    it 'retourne le temps de cuisson pour "Autre aliment"' do
      allow_any_instance_of(TTY::Prompt).to receive(:ask).and_return('10')
      prompt = TTY::Prompt.new
      expect(determiner_temps('Autre aliment', prompt)).to eq(10)
    end

    it 'retourne 0 pour un temps de cuisson invalide pour "Autre aliment"' do
      allow_any_instance_of(TTY::Prompt).to receive(:ask).and_return('-5')
      prompt = TTY::Prompt.new
      expect(determiner_temps('Autre aliment', prompt)).to eq(0)
    end
  end

  describe '#afficher_alertes' do
    it 'affiche les alertes disponibles' do
      allow_any_instance_of(TTY::Prompt).to receive(:select).and_return('Son')
      prompt = TTY::Prompt.new
      expect { afficher_alertes(prompt, 10) }.to output(/Types d'alertes disponibles/).to_stdout
    end
  end
end
