# haveibeenpwned

A CLI for [Troy Hunt's Have I Been Pwned? API](https://haveibeenpwned.com/API/v2) using Go

## Usage

```
SubCommands:

Please indicate subcommand:
account                 Get breaches for a particular account
breach                  Get information of a particular breach
breaches                Get all breaches
pastes                  Get all pastes for a particular account
```

Each command has its own -h option and each mandatory fields.

```
$ go run main.go account -h
Usage of account:
  -domain string
        Filters the result set to only breaches against the domain specified
  -name string
        Account to be validated (required)
  -truncate
        Returns only the name of the breach
  -unverified
        Returns breaches that have been flagged as unverified

$ go run main.go breach -h
Usage of breach:
  -name string
        Breach name (required)

$ go run main.go breaches -h
Usage of breaches:
  -domain string
        Filters the result set to only breaches against the domain specified

$ go run main.go pastes -h
Usage of pastes:
  -email string
        Email to be searched (required)
```

## License

This tool is distributed under the [MIT License](LICENSE.txt).