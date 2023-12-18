package main

import (
	"bytes"
	"fmt"
	"testing"
)

type MockAlerte struct {
	Declenche bool
}

func (a *MockAlerte) Declencher() {
	a.Declenche = true
}

type CustomWriter interface {
	Print(a ...interface{}) (n int, err error)
}

type CustomReader interface {
	Scanln(a ...interface{}) (n int, err error)
}

type CustomBuffer struct {
	buf *bytes.Buffer
}

func NewCustomBuffer() *CustomBuffer {
	return &CustomBuffer{
		buf: new(bytes.Buffer),
	}
}

func (c *CustomBuffer) Print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(c.buf, a...)
}

func (c *CustomBuffer) Scanln(a ...interface{}) (n int, err error) {
	return fmt.Fscanln(c.buf, a...)
}

type CustomIO struct {
	CustomWriter
	CustomReader
}

func (c *CustomIO) SetOutput(w CustomWriter) {
	c.CustomWriter = w
}

func (c *CustomIO) SetInput(r CustomReader) {
	c.CustomReader = r
}

func TestChoisirMode(t *testing.T) {
	mockIO := NewCustomIO()
	mockIO.SetOutput(NewCustomBuffer())
	mockIO.SetInput(NewCustomBuffer())
	mockAlerte := &MockAlerte{}
	riceCooker := RiceCooker{
		Alerte: mockAlerte,
	}

	fmt.Print = mockIO.Print
	fmt.Scanln = mockIO.Scanln

	tests := []struct {
		input          string
		tempsCuisson   int
		alerteDeclenche bool
		expectedOutput string
	}{
		{"1\n", 2, true, "Mode sélectionné : 1\n"},
		{"4\n10\n1\n2\n", 10, true, "Mode Autre Aliment sélectionné - Cuisson pendant 10 secondes.\nTypes d'alertes disponibles :\n1. Son\n2. Lumières clignotantes\nChoisissez le type d'alerte pour signaler la fin de la cuisson (1/2) : La cuisson se déroulera pendant 10 secondes.\n"},
		{"4\n0\n", 0, false, "Mode Autre Aliment sélectionné - Cuisson pendant 0 secondes.\nTemps invalide. La cuisson ne sera pas effectuée.\n"},
		{"invalid\n", 0, false, "Choix non valide. La cuisson ne sera pas effectuée.\n"},
	}

	for _, test := range tests {
		mockIO.CustomReader.(*CustomBuffer).buf.Reset()
		mockIO.CustomWriter.(*CustomBuffer).buf.Reset()
		mockIO.CustomReader.(*CustomBuffer).buf.WriteString(test.input)
		riceCooker.ChoisirMode()

		if mockAlerte.Declenche != test.alerteDeclenche {
			t.Errorf("Pour la saisie %s, attendu que l'alerte soit déclenchée : %t, mais obtenu : %t", test.input, test.alerteDeclenche, mockAlerte.Declenche)
		}

		actualOutput := mockIO.CustomWriter.(*CustomBuffer).buf.String()
		if actualOutput != test.expectedOutput {
			t.Errorf("Pour la saisie %s, attendu la sortie :\n%s\nmais obtenu :\n%s", test.input, test.expectedOutput, actualOutput)
		}
	}
}
