package main

import (
	"fmt"
	"regexp"
	"strings"
)

type Civilization struct {
	Name         string
	Leader       string
	UAbility     Ability
	UBuilding    Building
	UImprovement Improvement
	UUnit        Unit
	Bias         string
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

func parser(text string) {
	civRegex := regexp.MustCompile(`(?m)^(.*) - (.*)$`)
	abilityRegex := regexp.MustCompile(`(?m)^Ability: ([^-\n]+) - (.*)$`)
	improvementRegex := regexp.MustCompile(`(?m)^Improvement: ([^-\n]+) - (.*)$`)
	buildingRegex := regexp.MustCompile(`(?m)^Building: ([^-\n]+) - (.*)$`)
	unitRegex := regexp.MustCompile(`(?m)^Unit: ([^-\n]+) - (.*)$`)
	// For Madagascar
	unitPromotionRegex := regexp.MustCompile(`(?m)^(Kelimazala|Ramahavaly|Manjakatsiroa|Rafantaka|Mosasa|Rabehaza|Ambohimanambola|Sehatra|Lambamena|Famadiahona|Razana|Masina) - (.*)$`)
	biasRegex := regexp.MustCompile(`(?m)^Bias: (.*)$`)

	entries := regexp.MustCompile(`(?m)(?s)([A-Za-z ]+ - [A-Za-z]+)(.*?)Bias: [^\n]+`).FindAllString(text, -1)

	var civilizations []Civilization

	for _, entry := range entries {
		entry = strings.TrimSpace(entry)

		civMatches := civRegex.FindStringSubmatch(entry)
		abilityMatches := abilityRegex.FindStringSubmatch(entry)
		improvementMatches := improvementRegex.FindStringSubmatch(entry)
		buildingMatches := buildingRegex.FindStringSubmatch(entry)
		unitMatches := unitRegex.FindStringSubmatch(entry)
		biasMatches := biasRegex.FindStringSubmatch(entry)

		civ := Civilization{}

		if len(civMatches) > 2 {
			civ.Name = strings.TrimSpace(civMatches[1])
			civ.Leader = strings.TrimSpace(civMatches[2])
		}

		if len(abilityMatches) > 2 {
			civ.UAbility = Ability{
				Name:   strings.TrimSpace(abilityMatches[1]),
				Effect: strings.TrimSpace(abilityMatches[2]),
			}
		}

		if len(improvementMatches) > 2 {
			civ.UImprovement = Improvement{
				Name:   strings.TrimSpace(improvementMatches[1]),
				Effect: strings.TrimSpace(improvementMatches[2]),
			}
		}

		if len(buildingMatches) > 2 {
			civ.UBuilding = Building{
				Name:   strings.TrimSpace(buildingMatches[1]),
				Effect: strings.TrimSpace(buildingMatches[2]),
			}
		}

		if len(unitMatches) > 2 {
			unit := Unit{
				Name:   strings.TrimSpace(unitMatches[1]),
				Effect: strings.TrimSpace(unitMatches[2]),
			}

			// This logic only exists for Madagascar, despite the fact that the Civ is terrible
			unitPromotions := unitPromotionRegex.FindAllStringSubmatch(entry, -1)
			for _, promotion := range unitPromotions {
				if len(promotion) > 2 {
					unit.Promotions = append(unit.Promotions, strings.TrimSpace(promotion[1]+" - "+promotion[2]))
				}
			}
		}

		if len(biasMatches) > 1 {
			civ.Bias = strings.TrimSpace(biasMatches[1])
		}

		civilizations = append(civilizations, civ)

	}
	for _, civ := range civilizations {
		fmt.Printf("Civilization: %s\n", civ.Name)
		fmt.Printf("Leader: %s\n", civ.Leader)
		fmt.Printf("Ability: %s - %s\n", civ.UAbility.Name, civ.UAbility.Effect)
		if civ.UImprovement.Name != "" {
			fmt.Printf("Improvement: %s - %s\n", civ.UImprovement.Name, civ.UImprovement.Effect)
		}
		fmt.Printf("Unique Building: %s - %s\n", civ.UBuilding.Name, civ.UBuilding.Effect)
		if civ.UUnit.Name != "" {
			fmt.Printf("Unique Unit: %s - %s\n", civ.UUnit.Name, civ.UUnit.Effect)
		}
		fmt.Printf("Bias: %s\n\n", civ.Bias)
	}
}
