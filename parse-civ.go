package main

import (
	"fmt"
	"regexp"
	"strings"
)

type Civilization struct {
	Name          string
	Leader        string
	UAbility      Ability
	UBuildings    []Building
	UImprovements []Improvement
	UUnits        []Unit
	UGreatPerson  GreatPerson
	Bias          string
}

type Ability struct {
	Name   string
	Effect string
}

type Building struct {
	Name   string
	Effect string
}

type Unit struct {
	Name       string
	Effect     string
	Promotions []string
}

type Improvement struct {
	Name   string
	Effect string
}

type GreatPerson struct {
	Name   string
	Effect string
}

// ANSI color codes
const (
	Red   = "\033[31m"
	Reset = "\033[0m"
)

func parser(text string) {
	// Regex to match each civilization block
	civBlockRegex := regexp.MustCompile(`(?m)^([^\n]+- [^\n]+)\n((?:.*\n)*?)^Bias: (.*)$`)
	blocks := civBlockRegex.FindAllStringSubmatch(text, -1)

	// Regex patterns for each field
	abilityRegex := regexp.MustCompile(`(?m)^Ability: ([^\n]+) - (.*)$`)
	improvementRegex := regexp.MustCompile(`(?m)^Improvement: ([^\n]+) - (.*)$`)
	buildingRegex := regexp.MustCompile(`(?m)^Building: ([^\n]+) - (.*)$`)
	unitRegex := regexp.MustCompile(`(?m)^Unit: ([^\n]+) - (.*)$`)
	greatPersonRegex := regexp.MustCompile(`(?m)^Great Person: ([^\n]+) - (.*)$`)

	var civilizations []Civilization

	for _, block := range blocks {
		// Extract name and leader
		civHeader := block[1]
		civBody := block[2]
		bias := block[3]

		civHeaderParts := strings.SplitN(civHeader, "-", 2)
		if len(civHeaderParts) < 2 {
			continue // Skip if unable to split name and leader
		}

		civ := Civilization{
			Name:          strings.TrimSpace(civHeaderParts[0]),
			Leader:        strings.TrimSpace(civHeaderParts[1]),
			Bias:          strings.TrimSpace(bias),
			UImprovements: []Improvement{},
			UBuildings:    []Building{},
			UUnits:        []Unit{},
			UGreatPerson:  GreatPerson{},
		}

		// Extract Ability
		if abilityMatches := abilityRegex.FindStringSubmatch(civBody); len(abilityMatches) > 2 {
			civ.UAbility = Ability{
				Name:   strings.TrimSpace(abilityMatches[1]),
				Effect: strings.TrimSpace(abilityMatches[2]),
			}
		}

		// Extract Great Person
		if greatPersonMatches := greatPersonRegex.FindStringSubmatch(civBody); len(greatPersonMatches) > 2 {
			civ.UGreatPerson = GreatPerson{
				Name:   strings.TrimSpace(greatPersonMatches[1]),
				Effect: strings.TrimSpace(greatPersonMatches[2]),
			}
		}

		// Extract Improvement
		for _, match := range improvementRegex.FindAllStringSubmatch(civBody, -1) {
			if len(match) > 2 {
				civ.UImprovements = append(civ.UImprovements, Improvement{
					Name:   strings.TrimSpace(match[1]),
					Effect: strings.TrimSpace(match[2]),
				})
			}
		}

		// Extract Buildings
		for _, match := range buildingRegex.FindAllStringSubmatch(civBody, -1) {
			if len(match) > 2 {
				civ.UBuildings = append(civ.UBuildings, Building{
					Name:   strings.TrimSpace(match[1]),
					Effect: strings.TrimSpace(match[2]),
				})
			}
		}

		// Extract Unique Units
		for _, match := range unitRegex.FindAllStringSubmatch(civBody, -1) {
			if len(match) > 2 {
				civ.UUnits = append(civ.UUnits, Unit{
					Name:   strings.TrimSpace(match[1]),
					Effect: strings.TrimSpace(match[2]),
				})
			}
		}

		// Append to civilizations slice
		civilizations = append(civilizations, civ)
	}

	// Print all extracted civilizations
	for _, civ := range civilizations {
		fmt.Printf("%sCivilization:%s %s\n", Red, Reset, civ.Name)
		fmt.Printf("%sLeader:%s %s\n", Red, Reset, civ.Leader)
		fmt.Printf("%sAbility:%s %s - %s\n", Red, Reset, civ.UAbility.Name, civ.UAbility.Effect)

		if len(civ.UGreatPerson.Name) > 0 {
			fmt.Printf("%sGreat Person:%s %s - %s\n", Red, Reset, civ.UGreatPerson.Name, civ.UGreatPerson.Effect)
		}

		for _, imp := range civ.UImprovements {
			fmt.Printf("%sImprovement:%s %s - %s\n", Red, Reset, imp.Name, imp.Effect)
		}
		for _, bld := range civ.UBuildings {
			fmt.Printf("%sBuilding:%s %s - %s\n", Red, Reset, bld.Name, bld.Effect)
		}
		for _, unit := range civ.UUnits {
			fmt.Printf("%sUnique Unit:%s %s - %s\n", Red, Reset, unit.Name, unit.Effect)
		}
		fmt.Printf("%sBias:%s %s\n\n", Red, Reset, civ.Bias)
	}
}
