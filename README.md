# exopulse epggen tool
Golang tool is a simple code generator aimed to remedy golang's lack of generics.

[![CircleCI](https://circleci.com/gh/exopulse/epggen.svg?style=svg)](https://circleci.com/gh/exopulse/epggen)
[![Build Status](https://travis-ci.org/exopulse/epggen.svg?branch=master)](https://travis-ci.org/exopulse/epggen)
[![GitHub license](https://img.shields.io/github/license/exopulse/epggen.svg)](https://github.com/exopulse/epggen/blob/master/LICENSE)

# Overview

This project contains command line tool _eppgen_ for code generation. No additional libraries are needed or exposed.

# Features

At the moment only map generator is available. Map template provides common map operations while providing easier way to handle maps as types.

# Using epggen package

## Installing package

Use _go install_ to install the latest version of this tool.

    $ go install github.com/exopulse/epggen
    
## Use command line help switch to display available features

    $ epggen --help

## Use epggen tool in go generate go files from command line

    $ epggen map ID:Person
    $ epggen map int:Account

This will generate two new files:
- person_map.go
- account_map.go
    
in directory where command was invoked.
     
## Use epggen tool in go generate tool chain

```go
package model

//go:generate epggen map ID:Person
//go:generate epggen map ID:Gender

// ID is an alias for all IDs.
type ID int
```

Run go generate to generate required files:

    $ go generate ./...

## Zero values

Some templates require that you write zero value function for a type. This function is invoked when some of the generated functions return default value.
For example, methods FindKey() and FindValue() from _map_ template require such functions. 

```go
// GetIDZeroValue returns zero value for ID.
func GetIDZeroValue() ID {
	return 0
}

// GetPersonZeroValue returns zero value for Person.
func GetPersonZeroValue() Person {
	return Person{}
}
```
     
# About the project

## Contributors

* [exopulse](https://github.com/exopulse)

## License

Epggen is released under the MIT license. See
[LICENSE](https://github.com/exopulse/epggen/blob/master/LICENSE)
