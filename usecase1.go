package main

import (
	"fmt"
	"slices"
)

func main() {

	nodeTypes := make([]string, 0)
	ProductFamilies := make(map[string][]string)

	ProductFamilies["SR"] = append(nodeTypes, "SR-12", "SR-7")
	ProductFamilies["ESS"] = append(nodeTypes, "ESS-12", "ESS-7")
	ProductFamilies["IXR"] = append(nodeTypes, "IXR-R6", "IXR-e2")
	ProductFamilies["XRS"] = append(nodeTypes, "XRS-20")
	ProductFamilies["IPEdge"] = append(nodeTypes, "SAS-R12", "SAS-R6")

	PrintFamiles(ProductFamilies)

	AddFamilyAndNodeTypes(ProductFamilies, "Sar-hm", []string{"Sar-hm", "Sar-hmc"})

	PrintFamiles(ProductFamilies)

	AddFamilyAndNodeTypes(ProductFamilies, "SR", []string{"SR-7", "SR-12e"})

	DeleteDeprecatedNode(ProductFamilies, "SAS-R6")

	PrintFamiles(ProductFamilies)
}

func DeleteDeprecatedNode(productFamiles map[string][]string, deprecatedNode string) {

	for key, values := range productFamiles {
		nodeTypes := values
		if slices.Contains(nodeTypes, deprecatedNode) {
			for index, nodeType := range nodeTypes {
				if nodeType == deprecatedNode {
					productFamiles[key] = append(nodeTypes[:index], nodeTypes[index+1:]...)
				}
			}
		}
	}
}

func PrintFamiles(productFamiles map[string][]string) {
	fmt.Println("Printing  the product Families ", productFamiles)
}

func AddFamilyAndNodeTypes(productFamiles map[string][]string, Family string, NodeTypes []string) {

	if productFamiles[Family] != nil {
		fmt.Println("Node Family Exist, Update the nodeTypes", Family)
		ExistingNodeTypes := productFamiles[Family]
		for _, nodeType := range NodeTypes {
			if !slices.Contains(ExistingNodeTypes, nodeType) {
				productFamiles[Family] = append(ExistingNodeTypes, nodeType)
			}
		}
	} else {
		fmt.Println("Not Found. Adding the Family ", Family)
		productFamiles[Family] = NodeTypes
	}
}
