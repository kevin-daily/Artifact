package main

import "fmt"

func printHelp() {
	fmt.Println()
	fmt.Println("rest - Recover HP by resting" +
		"\nfight - Fight a single monster" +
		"\nfight -auto - Fight monster a specified number of times. Auto rests after each fight" +
		"\nmove - Move to a specified location (Uses X Y coordinates)" +
		"\ngather - Gather a single resource" +
		"\ngather -auto - Gather a specified number of resources" +
		"\nequip - Equip specified item" +
		"\nunequip - Unequip specified item" +
		"\ncraft - Craft an item" +
		"\ninventory - List the items in your inventory" +
		"\nresource - Give details about specified item" +
		"\nresource -all - Give list of all items" +
		"\nmonster - Give details about specified monster" +
		"\nmonster -all - Give list of all monsters" +
		"\nitem - Give details about specified item" +
		"\nitem -all - Give list of all items")
	fmt.Println()
}
