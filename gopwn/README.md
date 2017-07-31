# haveibeenpwned

A CLI for [Troy Hunt's Have I Been Pwned? API](https://haveibeenpwned.com/API/v2) using Go

## Instalation

```
$ git clone https://github.com/jrizza/haveibeenpwned
$ cd haveibeenpwned/gopwn
$ go build
```

## Usage

```
$ ./gopwn
Please indicate subcommand:
account                 Get breaches for a particular account
breach                  Get information of a particular breach
breaches                Get all breaches
pastes                  Get all pastes for a particular account
```

Each command has its own -h option and each mandatory fields.

```
$ ./gopwn account -h
Usage of account:
  -domain string
        Filters the result set to only breaches against the domain specified
  -name string
        Account to be validated (required)
  -truncate
        Returns only the name of the breach
  -unverified
        Returns breaches that have been flagged as unverified

$ ./gopwn breach -h
Usage of breach:
  -name string
        Breach name (required)

$ ./gopwn breaches -h
Usage of breaches:
  -domain string
        Filters the result set to only breaches against the domain specified

$ ./gopwn pastes -h
Usage of pastes:
  -email string
        Email to be searched (required)
```

## License

This tool is distributed under the [MIT License](LICENSE).