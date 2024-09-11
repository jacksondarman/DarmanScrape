package tests

import (
	"reflect"
	"testing"

	"github.com/jacksondarman/lekscrape/internal/parser"
)

// Correctly parses civilizations from well-formed text input
func TestParseCivilizationsWellFormedInput(t *testing.T) {
	text := "Rome - Julius Caesar\nAbility: Roman Legions - Can build roads\nImprovement: Aqueduct - Provides fresh water\nBuilding: Colosseum - Increases happiness\nUnit: Legion - Stronger than swordsman\nGreat Person: Augustus - Increases culture\nBias: Coastal"

	expected := []parser.Civilization{
		{
			Name:   "Rome",
			Leader: "Julius Caesar",
			Bias:   "Coastal",
			UAbility: []parser.Ability{{
				Name:   "Roman Legions",
				Effect: "Can build roads",
			}},
			UImprovements: []parser.Improvement{{Name: "Aqueduct", Effect: "Provides fresh water"}},
			UBuildings:    []parser.Building{{Name: "Colosseum", Effect: "Increases happiness"}},
			UUnits:        []parser.Unit{{Name: "Legion", Effect: "Stronger than swordsman"}},
			UGreatPerson:  []parser.GreatPerson{{Name: "Augustus", Effect: "Increases culture"}},
		},
	}
	result := parser.ParseCivilizations(text)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

// Handles text input with missing civilization blocks
func TestParseCivilizationsMissingBlocks(t *testing.T) {
	text := "Rome - Julius Caesar\nAbility: Roman Legions - Can build roads\nBias: Coastal"

	expected := []parser.Civilization{
		{
			Name:          "Rome",
			Leader:        "Julius Caesar",
			Bias:          "Coastal",
			UAbility:      []parser.Ability{{Name: "Roman Legions", Effect: "Can build roads"}},
			UImprovements: []parser.Improvement{},
			UBuildings:    []parser.Building{},
			UUnits:        []parser.Unit{},
			UGreatPerson:  []parser.GreatPerson{},
		},
	}

	result := parser.ParseCivilizations(text)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
