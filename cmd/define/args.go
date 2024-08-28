package main

import (
	"fmt"

	"github.com/carlmjohnson/versioninfo"
)

type ProgramArgs struct {
	Word    string `arg:"positional, required"`
	ListAll bool   `arg:"-a, --list-all" default:"false" help:"List all definitions for a word"`
	Copy    bool   `arg:"-c,--copy" default:"false" help:"Copy definition to the clipboard. Does not work with --list-all"`
	MaxWidth int    `arg:"-w,--width" default:"80" help:"The max width of the output table"`
}

func (args ProgramArgs) Description() string {
	return "A simple command line application that finds the definition to a word"
}

func (ProgramArgs) Version() string {
	return fmt.Sprintln("Version: ", versioninfo.Short())
}
