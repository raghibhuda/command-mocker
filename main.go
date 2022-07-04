package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	listCommand := flag.NewFlagSet("ls", flag.ExitOnError)
	pwdCommand := flag.NewFlagSet("pwd", flag.ExitOnError)
	cpCommand := flag.NewFlagSet("cp", flag.ExitOnError)
	mvCommand := flag.NewFlagSet("mv", flag.ExitOnError)
	rmCommand := flag.NewFlagSet("rm", flag.ExitOnError)
	mkdirCommand := flag.NewFlagSet("mkdir", flag.ExitOnError)
	wcCommand := flag.NewFlagSet("head", flag.ExitOnError)

	rmRecursiveParameter := rmCommand.Bool("r", false, "Delete dir/files recursively")
	mkdirNestedParameter := mkdirCommand.Bool("p", false, "Make directory with nested children")

	if len(os.Args) < 2 {
		fmt.Println("Any command is required: ")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "ls":
		listCommand.Parse(os.Args[2:])
	case "pwd":
		pwdCommand.Parse(os.Args[2:])
	case "cp":
		cpCommand.Parse(os.Args[2:])
	case "mv":
		mvCommand.Parse(os.Args[2:])
	case "rm":
		rmCommand.Parse(os.Args[2:])
	case "mkdir":
		mkdirCommand.Parse(os.Args[2:])
	case "wc":
		wcCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if pwdCommand.Parsed() {
		fmt.Println(runPWD())
	}

	if rmCommand.Parsed() {

		if *rmRecursiveParameter {
			rmArgs := os.Args[3:]
			runRM(rmArgs, true)
		} else {
			rmArgs := os.Args[2:]
			err := runRM(rmArgs, false)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}

	}

	if mkdirCommand.Parsed() {
		if *mkdirNestedParameter {
			mkdirArgs := os.Args[3:]
			runMKDIR(mkdirArgs, true)

		} else {
			mkdirArgs := os.Args[2:]
			runMKDIR(mkdirArgs, false)
		}

	}

}
