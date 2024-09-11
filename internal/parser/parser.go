package parser

import (
	"log"
	"regexp"
	"strings"
)

type Civilization struct {
	Name          string
	Leader        string
	UAbility      []Ability
	UBuildings    []Building
	UImprovements []Improvement
	UUnits        []Unit
	UGreatPerson  []GreatPerson
	Bias          string
}

type CivAttribute struct {
	Name   string
	Effect string
}

type Ability CivAttribute
type Building CivAttribute
type Improvement CivAttribute
type Unit CivAttribute
type GreatPerson CivAttribute

// ANSI color codes
const (
	Red   = "\033[31m"
	Reset = "\033[0m"
)

/*
*

  - parseCivilizations parses the given text to extract information about Civilizations.

  - First, this function uses regular expressions to separate the text into "blocks" (each block corresponds to a Civilization)

  - Then, this function uses regular expressions to define various attributes of Civilizations within the Civilization struct.

  - The function returns a slice of Civilization structs containing the extracted data.
*/
func ParseCivilizations(text string) []Civilization {
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
			UAbility:      []Ability{},
			UImprovements: []Improvement{},
			UBuildings:    []Building{},
			UUnits:        []Unit{},
			UGreatPerson:  []GreatPerson{},
		}

		log.Printf("New Civilization %s found with leader %s\n", Red+civ.Name+Reset, Red+civ.Leader+Reset)

		// Extract Abilities
		for _, match := range abilityRegex.FindAllStringSubmatch(civBody, -1) {
			if len(match) > 2 {
				civ.UAbility = append(civ.UAbility, Ability{
					Name:   strings.TrimSpace(match[1]),
					Effect: strings.TrimSpace(match[2]),
				})
			}
		}

		// Extract Great People
		for _, match := range greatPersonRegex.FindAllStringSubmatch(civBody, -1) {
			if len(match) > 2 {
				civ.UGreatPerson = append(civ.UGreatPerson, GreatPerson{
					Name:   strings.TrimSpace(match[1]),
					Effect: strings.TrimSpace(match[2]),
				})
			}
		}

		// Extract Improvements
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

	return civilizations
}
