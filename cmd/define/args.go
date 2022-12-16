package main

type ProgramArgs struct {
	Word string `arg:"positional, required"`
}

func (args ProgramArgs) Description() string {
	return "A simple command line application that finds the definition to a word"
}
