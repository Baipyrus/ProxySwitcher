package util

import (
	"fmt"
	"strings"
)

func promptArguments() []string {
	var args []string = []string{}
	var count int

argumentLoop:
	for true {
		var arg string
		fmt.Printf("Arg #%d? ", count+1)
		fmt.Scanln(&arg)

		if arg == "" {
			break argumentLoop
		}

		args = append(args, arg)
		count++
	}

	return args
}

func promptType() VariantType {
	var variantType VariantType

typeLoop:
	for true {
		var input string
		fmt.Print("Type? ")
		fmt.Scanln(&input)

		switch strings.ToLower(input) {
		case "":
			variantType = TEXT
			break typeLoop
		case "variable":
			variantType = VARIABLE
			break typeLoop
		default:
			fmt.Println("Given input not convertable to type!")
		}
	}

	return variantType
}

func PromptVariants() []*Variant {
	var variants []*Variant
	var count int

variantLoop:
	for true {
		fmt.Printf("Prompting 'args' for variant #%d:\n", count+1)
		args := promptArguments()

		var equator string
		fmt.Print("Equator? ")
		fmt.Scanln(&equator)

		if strings.ToLower(equator) == "skip" {
			break variantLoop
		}

		variantType := promptType()

		variant := Variant{Arguments: args, Type: variantType, Equator: equator}
		variants = append(variants, &variant)
		count++
	}

	return variants
}
