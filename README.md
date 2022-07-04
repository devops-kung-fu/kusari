![](img/kusari128x128.png)

***NOTE: There's some working code in here, but don't bet the farm on it just yet...***

# kusari (鎖)

[![GoDoc](https://godoc.org/github.com/devops-kung-fu/kusari?status.svg)](https://pkg.go.dev/github.com/devops-kung-fu/kusari)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/devops-kung-fu/kusari) 
[![Go Report Card](https://goreportcard.com/badge/github.com/devops-kung-fu/kusari)](https://goreportcard.com/report/github.com/devops-kung-fu/kusari) 
[![codecov](https://codecov.io/gh/devops-kung-fu/kusari/branch/main/graph/badge.svg?token=P9WBOBQTOB)](https://codecov.io/gh/devops-kung-fu/kusari) 
[![SBOM](https://img.shields.io/badge/CyloneDX-SBoM-informational)](kusari-sbom.json)

A simple blockchain module for Golang.

## Overview

[DKFM](https://github.com/devops-kung-fu) is experimenting with a few ideas around using block chain concepts for creating irrefutable evidence that security controls are being executed during build and deployments of code in CI/CD pipelines. 

This module encapsulates functionality to create and manage blockchains.

## Using kusari

Well, buyer beware here. Right now this module is in a really early state. It's also really dang noisy and outputs a ton of logs. Ensure that you have good log management in any consumer of this module. 

If you want to suppress logging (you'll have to do this in your app) you can do this:

``` go

log.SetOutput(ioutil.Discard)

```

## Development

## Overview

In order to use contribute and participate in the development of ```kusari``` you'll need to have an updated Go environment. Before you start, please view the [Contributing](CONTRIBUTING.md) and [Code of Conduct](CODE_OF_CONDUCT.md) files in this repository.

## Prerequisites

This project makes use of [DKFM](https://github.com/devops-kung-fu) tools such as [Hookz](https://github.com/devops-kung-fu/hookz), [Hinge](https://github.com/devops-kung-fu/hinge), and other open source tooling. Install these tools with the following commands:

``` bash

go install github.com/devops-kung-fu/hookz@latest
go install github.com/devops-kung-fu/lucha@latest
go install github.com/devops-kung-fu/hinge@latest
go install github.com/kisielk/errcheck@latest
go install golang.org/x/lint/golint@latest
go install github.com/fzipp/gocyclo@latest

```

## Software Bill of Materials

```kusari``` uses the CycloneDX to generate a Software Bill of Materials in CycloneDX format (v1.4) every time a developer commits code to this repository (as long as [Hookz](https://github.com/devops-kung-fu/hookz) is being used and is has been initialized in the working directory). More information for CycloneDX is available [here](https://cyclonedx.org)

The current SBoM for ```kusari``` is available [here](kusari-sbom.json).

## Credits

A big thank-you to our friends at [Freepik](https://www.freepik.com) for the ```kusari``` logo.