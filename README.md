# Entropy Library for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/chrisjchandler/entropy.svg)](https://pkg.go.dev/github.com/chrisjchandler/entropy)
[![Go Report Card](https://goreportcard.com/badge/github.com/chrisjchandler/entropy)](https://goreportcard.com/report/github.com/chrisjchandler/entropy)

A lightweight Go library to calculate the Shannon entropy of strings. This library is useful for detecting randomness in strings, such as in security applications like anomaly detection, DNS analysis, and more.

## Features

- Compute Shannon entropy for any string.
- Lightweight and easy to use.
- Includes unit tests for reliability.

## Installation

Install the library using `go mod init entropy`:

Usage
Import the library and calculate the entropy of a string:

package main

import (
    "fmt"
    "github.com/chrisjchandler/entropy"
)

func main() {
    input := "www.google.com"
    result := entropy.Calculate(input)
    fmt.Printf("Entropy of '%s' is: %.4f\n", input, result)
}

Example Output
For the input www.google.com

Entropy of 'www.google.com' is: 3.1808

API Reference
func Calculate(input string) float64
Calculates the Shannon entropy of a given string.

Parameter: input - The string for which entropy will be calculated.
Returns: float64 - The Shannon entropy value.

Testing
Run the included tests using go test:

Contributing
Contributions are welcome! Feel free to open an issue or submit a pull request.

Acknowledgments
Inspired by the need for efficient entropy calculations in DNS and anomaly detection scenarios.
