# Package haveibeenpwned 

## Overview

Package haveibeenpwned provides access to the [Have I been Pwned?](https://haveibeenpwned.com/) API, returning a BreachModel or a PasteModel if any breach/paste is found.

## Types

### type BreachModel

BreachModel Each breach contains a number of attributes describing the incident. In the future, these attributes may expand without the API being versioned.
```
type BreachModel struct {
    Name         string   `json:"Name,omitempty"`
    Title        string   `json:"Title,omitempty"`
    Domain       string   `json:"Domain,omitempty"`
    BreachDate   string   `json:"BreachDate,omitempty"`
    AddedDate    string   `json:"AddedDate,omitempty"`
    ModifiedDate string   `json:"ModifiedDate,omitempty"`
    PwnCount     int      `json:"PwnCount,omitempty"`
    Description  string   `json:"Description,omitempty"`
    DataClasses  []string `json:"DataClasses,omitempty"`
    IsVerified   bool     `json:"IsVerified,omitempty"`
    IsFabricated bool     `json:"IsFabricated,omitempty"`
    IsSensitive  bool     `json:"IsSensitive,omitempty"`
    IsRetired    bool     `json:"IsRetired,omitempty"`
    IsSpamList   bool     `json:"IsSpamList,omitempty"`
    LogoType     string   `json:"LogoType,omitempty"`
}
```

### type PasteModel

PasteModel Each paste contains a number of attributes describing it. In the future, these attributes may expand without the API being versioned.
```
type PasteModel struct {
    Source     string `json:"Source,omitempty"`
    ID         string `json:"Id,omitempty"`
    Title      string `json:"Title,omitempty"`
    Date       string `json:"Date,omitempty"`
    EmailCount int    `json:"EmailCount,omitempty"`
}
```

## Functions

### func BreachedAccount
```
func BreachedAccount(account, domainFilter string, truncate, unverified bool) ([]BreachModel, error)
```
BreachedAccount The most common use of the API is to return a list of all breaches a particular account has been involved in. The API takes a single parameter which is the account to be searched for. The account is not case sensitive and will be trimmed of leading or trailing white spaces. The account should always be URL encoded.

### func Breaches
```
func Breaches(domainFilter string) ([]BreachModel, error)
```
Breaches Getting all breached sites in the system. A "breach" is an instance of a system having been compromised by an attacker and the data disclosed.

### func Breach
```
func Breach(name string) (BreachModel, error)
```
Breach Sometimes just a single breach is required and this can be retrieved by the breach "name". This is the stable value which may or may not be the same as the breach "title" (which can change).

### func PasteAccount
```
func PasteAccount(email string) ([]PasteModel, error)
```
PasteAccount The API takes a single parameter which is the email address to be searched for. Unlike searching for breaches, usernames that are not email addresses cannot be searched for. The email is not case sensitive and will be trimmed of leading or trailing white spaces. The email should always be URL encoded.