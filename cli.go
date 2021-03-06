package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		n bool
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.BoolVar(&n, "dry-run", false, "show what would have been transferred")
	flags.BoolVar(&n, "n", false, "show what would have been transferred(Short)")

	flVersion := flags.Bool("version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if *flVersion {
		fmt.Println(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	// Check count of args
	parsedArgs := flags.Args()
	if len(parsedArgs) < 3 {
		fmt.Println("[Usage]\n$ rrn regexp replacement file|dir")
		return ExitCodeError
	}
	regexpStr := parsedArgs[0]
	replacement := parsedArgs[1]
	srcPath := parsedArgs[2]

	// Compile regexp
	re, err := regexp.Compile(regexpStr)
	if err != nil {
		fmt.Println("arg 1 must be regular expression")
		return ExitCodeError
	}

	// Check source path
	if _, err := os.Lstat(srcPath); err != nil {
		fmt.Println("arg 3 must be path string")
		return ExitCodeError
	}

	// Rename recursive
	err = filepath.Walk(srcPath,
		func(srcPath string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			// rename file only
			dir, srcFile := path.Split(srcPath)
			dstFile := re.ReplaceAllString(srcFile, replacement)

			if srcFile == dstFile {
				return nil
			}
			dstPath := dir + dstFile

			if n {
				fmt.Println(fmt.Sprintf("'%s' would be renamed to '%s'", srcPath, dstPath))
			} else {
				err = os.Rename(srcPath, dstPath)
				if err != nil {
					return err
				}
			}

			return nil
		})

	if err != nil {
		fmt.Println(err)
		return ExitCodeError
	}

	return ExitCodeOK
}
