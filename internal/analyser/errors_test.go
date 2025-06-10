package analyzer

import (
	"errors"
	"testing"
)

func TestFileNotFoundError(t *testing.T) {
	// FileNotFoundError
	path := "/chemin/inexistant.log"
	err := &FileNotFoundError{Path: path}

	// Vérifier le message d'erreur
	expected := "file not found: /chemin/inexistant.log"
	if err.Error() != expected {
		t.Errorf("Message d'erreur incorrect. Obtenu: %s, Attendu: %s", err.Error(), expected)
	}

	// IsFileNotFoundError
	if !IsFileNotFoundError(err) {
		t.Error("IsFileNotFoundError devrait retourner true")
	}

	// GetFileNotFoundError
	fnfErr, ok := GetFileNotFoundError(err)
	if !ok {
		t.Error("GetFileNotFoundError devrait réussir")
	}
	if fnfErr.Path != path {
		t.Errorf("Chemin incorrect. Obtenu: %s, Attendu: %s", fnfErr.Path, path)
	}

	// Tester avec une erreur différente
	otherErr := errors.New("une autre erreur")
	if IsFileNotFoundError(otherErr) {
		t.Error("IsFileNotFoundError devrait retourner false pour une erreur différente")
	}
}

func TestParsingError(t *testing.T) {
	// ParsingError
	logID := "log123"
	message := "format invalide"
	err := &ParsingError{LogID: logID, Message: message}

	// Vérifier le message d'erreur
	expected := "parsing error for log log123: format invalide"
	if err.Error() != expected {
		t.Errorf("Message d'erreur incorrect. Obtenu: %s, Attendu: %s", err.Error(), expected)
	}

	// IsParsingError
	if !IsParsingError(err) {
		t.Error("IsParsingError devrait retourner true")
	}

	// GetParsingError
	parseErr, ok := GetParsingError(err)
	if !ok {
		t.Error("GetParsingError devrait réussir")
	}
	if parseErr.LogID != logID || parseErr.Message != message {
		t.Errorf("Données incorrectes. LogID: %s (attendu %s), Message: %s (attendu %s)",
			parseErr.LogID, logID, parseErr.Message, message)
	}

	// Tester avec une erreur différente
	otherErr := errors.New("une autre erreur")
	if IsParsingError(otherErr) {
		t.Error("IsParsingError devrait retourner false pour une erreur différente")
	}
}

func TestErrorWrapping(t *testing.T) {
	// Tester l'encapsulation d'erreurs avec errors.As
	baseErr := &ParsingError{LogID: "log456", Message: "erreur de syntaxe"}
	wrappedErr := errors.New("erreur d'analyse: " + baseErr.Error())
	wrappedErr = errors.Join(wrappedErr, baseErr)

	// Vérifier qu'on peut toujours extraire l'erreur originale
	if !IsParsingError(wrappedErr) {
		t.Error("IsParsingError devrait fonctionner avec une erreur encapsulée")
	}

	parseErr, ok := GetParsingError(wrappedErr)
	if !ok {
		t.Error("GetParsingError devrait réussir avec une erreur encapsulée")
	}
	if parseErr.LogID != "log456" {
		t.Errorf("LogID incorrect. Obtenu: %s, Attendu: log456", parseErr.LogID)
	}
}
