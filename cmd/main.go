package main

import (
	"flag"
	"fmt"
	"os"

	pwn "github.com/jrizza/haveibeenpwned"
)

func main() {

	//Flags sets
	var accountCommand = flag.NewFlagSet("account", flag.ExitOnError)
	var breachCommand = flag.NewFlagSet("breach", flag.ExitOnError)
	var breachesCommand = flag.NewFlagSet("breaches", flag.ExitOnError)
	var pastesCommand = flag.NewFlagSet("pastes", flag.ExitOnError)

	//Flags for account subcommand
	var accountNameFlag = accountCommand.String("name", "", "Account to be validated (required)")
	var accountTruncateFlag = accountCommand.Bool("truncate", false, "Returns only the name of the breach")
	var accountDomainFlag = accountCommand.String("domain", "", "Filters the result set to only breaches against the domain specified")
	var accountUnverifiedFlag = accountCommand.Bool("unverified", false, "Returns breaches that have been flagged as unverified")

	//Flags for breaches subcommand
	var breachesDomainFlag = breachesCommand.String("domain", "", "Filters the result set to only breaches against the domain specified")

	//Flags for breach subcommand
	var breachNameFlag = breachCommand.String("name", "", "Breach name (required)")

	//Flags for pastes subcommand
	var pastesNameFlag = pastesCommand.String("email", "", "Email to be searched (required)")

	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("Please indicate subcommand:")
		fmt.Println("account			Get breaches for a particular account")
		fmt.Println("breach			Get information of a particular breach")
		fmt.Println("breaches		Get all breaches")
		fmt.Println("pastes 			Get all pastes for a particular account")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "account":
		accountCommand.Parse(os.Args[2:])
		if *accountNameFlag == "" {
			accountCommand.PrintDefaults()
			os.Exit(1)
		}
		data, err := pwn.BreachedAccount(*accountNameFlag, *accountDomainFlag, *accountTruncateFlag, *accountUnverifiedFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
		checkBreachesData(data)
	case "breaches":
		breachesCommand.Parse(os.Args[2:])
		data, err := pwn.Breaches(*breachesDomainFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
		checkBreachesData(data)
	case "breach":
		breachCommand.Parse(os.Args[2:])
		if *breachNameFlag == "" {
			breachCommand.PrintDefaults()
			os.Exit(1)
		}
		data, err := pwn.Breach(*breachNameFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
		checkBreachData(data)
	case "pastes":
		pastesCommand.Parse(os.Args[2:])
		if *pastesNameFlag == "" {
			pastesCommand.PrintDefaults()
			os.Exit(1)
		}
		data, err := pwn.PasteAccount(*pastesNameFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
		checkPasteData(data)
	default:
		fmt.Println("account, breach, breaches or pastes subcommand is required")
		os.Exit(1)
	}
	os.Exit(0)
}

func checkBreachesData(data []pwn.BreachModel) {
	if data == nil {
		fmt.Println("Good!, no breaches found!")
	}
	fmt.Printf("Breaches found: %d...\n", len(data))
	fmt.Println("----------------------------")

	for _, b := range data {
		fmt.Printf("Name: %s\n", b.Name)
		if b.Title != "" {
			fmt.Printf("Title: %s\nDomain: %s\nBreach Date: %s\nAdded Date: %s\nModified Date: %s\n", b.Title, b.Domain, b.BreachDate, b.AddedDate, b.ModifiedDate)
			fmt.Printf("Pwn Count: %d\nDescription: %s\nData Classes: %s\nIs Verified?: %v\nIs Fabricated?: %v\nIs Sensitive?: %v\nIs Retired?: %v\nIs Spam List?: %v\n", b.PwnCount, b.Description, b.DataClasses, b.IsVerified, b.IsFabricated, b.IsSensitive, b.IsRetired, b.IsSpamList)
		}
		fmt.Println("----------------------------")
	}
	os.Exit(0)
}

func checkBreachData(data pwn.BreachModel) {
	if data.Name == "" {
		fmt.Println("Good!, no breaches found!")
	} else {
		fmt.Println("Breach found...")
		fmt.Println("----------------------------")
		fmt.Printf("Title: %s\nDomain: %s\nBreach Date: %s\nAdded Date: %s\nModified Date: %s\n", data.Title, data.Domain, data.BreachDate, data.AddedDate, data.ModifiedDate)
		fmt.Printf("Pwn Count: %d\nDescription: %s\nData Classes: %s\nIs Verified?: %v\nIs Fabricated?: %v\nIs Sensitive?: %v\nIs Retired?: %v\nIs Spam List?: %v\n", data.PwnCount, data.Description, data.DataClasses, data.IsVerified, data.IsFabricated, data.IsSensitive, data.IsRetired, data.IsSpamList)
		fmt.Printf("Logo Type: %s\n", data.LogoType)
		fmt.Println("----------------------------")
	}
	os.Exit(0)
}

func checkPasteData(data []pwn.PasteModel) {
	if data == nil {
		fmt.Println("Good!, no pastes found!")
	}
	fmt.Printf("Pastes found: %d...\n", len(data))
	fmt.Println("----------------------------")
	for _, b := range data {
		fmt.Printf("ID: %s\nTitle: %s\nSource: %s\nDate: %s\nEmailCount: %d\n", b.ID, b.Title, b.Source, b.Date, b.EmailCount)
		fmt.Println("----------------------------")
	}
	os.Exit(0)
}
