const chai = require('chai');
const expect = chai.expect;
const sinon = require('sinon');
const prompt = require('prompt-sync')();

const choisirMode = require('../rice');

describe('Choix de mode valide', function () {
  it('devrait afficher le bon message et attribuer le bon temps de cuisson', function () {
    sinon.stub(prompt, 'call').returnsThis();

    const consoleSpy = sinon.spy(console, 'log');

    choisirMode();

    expect(consoleSpy.calledWith('Mode Riz Blanc sélectionné')).to.be.true;
    expect(consoleSpy.calledWith('La cuisson se déroulera pendant 2 secondes.')).to.be.true;

    prompt.restore();
    consoleSpy.restore();
  });
});

describe('Choix de mode invalide', function () {
  it('devrait afficher un message d\'erreur', function () {
    sinon.stub(prompt, 'call').returnsThis();

    const consoleSpy = sinon.spy(console, 'log');

    choisirMode();

    expect(consoleSpy.calledWith('Choix non valide')).to.be.true;

    prompt.restore();
    consoleSpy.restore();
  });
});

describe('Choix d\'alerte valide', function () {
  it('devrait planifier l\'alerte correctement en fonction du choix', function () {
    sinon.stub(prompt, 'call').returnsThis();

    const consoleSpy = sinon.spy(console, 'log');
    const clock = sinon.useFakeTimers();

    choisirMode();

    clock.tick(2000);

    expect(consoleSpy.calledWith('*BIP*BIP*BIP* La cuisson est terminée !')).to.be.true;

    prompt.restore();
    consoleSpy.restore();
    clock.restore();
  });
});

describe('Choix d\'alerte invalide', function () {
  it('devrait afficher un message d\'erreur', function () {
    sinon.stub(prompt, 'call').returnsThis();

    const consoleSpy = sinon.spy(console, 'log');

    choisirMode();

    expect(consoleSpy.calledWith('Choix non valide. Aucune alerte ne sera déclenchée.')).to.be.true;

    prompt.restore();
    consoleSpy.restore();
  });
});
