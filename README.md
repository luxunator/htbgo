<p align="center">
  <img src="https://user-images.githubusercontent.com/50147562/183465332-52660dfa-d905-43b1-b44f-4394c717d5b3.png" alt="htbgo banner" title="htbgo" width="500" height="350">
</p>

# htbgo
Golang Bindings for Hack The Box v4 API

## Getting Started
### Installation
To install htbgo you can use the standard method of installing packages with `go get`
```
go get github.com/luxunator/htbgo
```

### Usage
To start using the package in your program you first need to import it
```golang
import "github.com/luxunator/htbgo"
```

You will need to create a session with an application token that you can generate in your Hack The Box profile settings
```golang
session, err := htbgo.New("HTB_APP_TOKEN")
```

After creating a session you will be able to use methods to interact with the API

## Contributing
Contributions are welcome, we just ask that you please follow these guidelines:

* First open an issue for the bug or feature enhancement and be as descriptive as possible.
* Try to match current naming conventions as closely as possible.
* Create a Pull Request with your changes against the main branch.

## Notes

* This package was created with community documentation as no official documentation of the API was ever provided to the public so endpoints may be missing.
* As the API has been seen to change frequently keep in mind field types may change. We are sorry for this but we are at the hands of Hack The Box.
* Interfaces are being used to interpret dynamic types, for example Rank may be 0 or "unranked"

## Kudos
* [Angelo Geant Gaviola](https://github.com/KiDxS) - Helping create and maintain the package
* [LunatiX](https://github.com/Lunatix01) - Helping create and maintain the package
* [Kris Stanley](https://github.com/Propolisa) - Creating the HTB v4 API community documentation that was a major resource for this package
