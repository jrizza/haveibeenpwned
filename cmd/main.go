package main

import (
	"flag"
	"fmt"
	"log"
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
		fmt.Println("account, breach, breaches or pastes subcommand is required")
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
			log.Fatal(err)
		}
		if data == nil {
			log.Printf("No breaches for account: " + *accountNameFlag)
		}
		for _, b := range data {
			log.Printf("%s", b.Title)
		}
	case "breaches":
		breachesCommand.Parse(os.Args[2:])
		_, err := pwn.Breaches(*breachesDomainFlag)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	case "breach":
		breachCommand.Parse(os.Args[2:])
		if *breachNameFlag == "" {
			breachCommand.PrintDefaults()
			os.Exit(1)
		}
		_, err := pwn.Breach(*breachNameFlag)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	case "pastes":
		pastesCommand.Parse(os.Args[2:])
		if *pastesNameFlag == "" {
			pastesCommand.PrintDefaults()
			os.Exit(1)
		}
		_, err := pwn.PasteAccount(*pastesNameFlag)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	default:
		fmt.Println("account, breach, breaches or pastes subcommand is required")
		os.Exit(1)
	}
}
