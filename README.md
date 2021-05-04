# Json-validator with schemas
This repository show how to validate complex rules in a Json instance from a Json Schema

## How to run
From the current directory execute in the terminal:   
```
$ go run main.go
```
This going to display the json tests (located in [Test Folder](json-test))   
- Lines for files with `*-ok` suffix should display `The document is valid`   
- Lines for files with `*-error` suffix should display each condition failed in the json tested.