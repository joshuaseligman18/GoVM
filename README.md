# GoVM

## Contents
1. [Overview](#overview)
2. [Setup Instructions](#setup-instructions)
3. [Instruction Set](#instruction-set)
## Overview
GoVM is a virtual ARM processor written in Go. The project was inspired by Professor Gormanly's 6502 project from his Computer Organization and Architecture class at Marist College. The instruction set is based off of that described in *Computer Organization and Design ARM Edition: The Hardware Software Interface* by David A. Patterson and John L. Hennessy. By using the ARM architecture, GoVM is able to take advantage of more modern CPU design practices such as scalar processing, which improves performance and maximizes the utilization of the hardware of the CPU.
## Setup Instructions
1. Install Go from the official website [here](https://go.dev/)
2. Create a Go module by running `go mod init <project-name>` in the directory of your project
3. Add GoVM to your project dependencies by running `go get github.com/joshuaseligman/GoVM@v0.6.0`
4. Import GoVM at the top of your code: `import "github.com/joshuaseligman/GoVM"`

## Instruction Set