package main

import (
	"reflect"
	"testing"
)

// Correctly parses civilizations from well-formed text input
func TestParseCivilizationsWellFormedInput(t *testing.T) {
	text := "Rome - Julius Caesar\nAbility: Roman Legions - Can build roads\nImprovement: Aqueduct - Provides fresh water\nBuilding: Colosseum - Increases happiness\nUnit: Legion - Stronger than swordsman\nGreat Person: Augustus - Increases culture\nBias: Coastal"

	expected := []Civilization{
		{
			Name:   "Rome",
			Leader: "Julius Caesar",
			Bias:   "Coastal",
			UAbility: Ability{
				Name:   "Roman Legions",
				Effect: "Can build roads",
			},
			UImprovements: []Improvement{
				{Name: "Aqueduct", Effect: "Provides fresh water"},
			},
			UBuildings: []Building{
				{Name: "Colosseum", Effect: "Increases happiness"},
			},
			UUnits: []Unit{
				{Name: "Legion", Effect: "Stronger than swordsman"},
			},
			UGreatPerson: GreatPerson{
				Name:   "Augustus",
				Effect: "Increases culture",
			},
		},
	}

	result := parseCivilizations(text)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

// Handles text input with missing civilization blocks
func TestParseCivilizationsMissingBlocks(t *testing.T) {
	text := "Rome - Julius Caesar\nAbility: Roman Legions - Can build roads\nBias: Coastal"

	expected := []Civilization{
		{
			Name:   "Rome",
			Leader: "Julius Caesar",
			Bias:   "Coastal",
			UAbility: Ability{
				Name:   "Roman Legions",
				Effect: "Can build roads",
			},
			UImprovements: []Improvement{},
			UBuildings:    []Building{},
			UUnits:        []Unit{},
			UGreatPerson:  GreatPerson{},
		},
	}

	result := parseCivilizations(text)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
