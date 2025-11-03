package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type Flags struct {
	count, repeated, uniq, register bool
	numFields, numChars             int
	inputFile, outputFile           string
}

func (f *Flags) Init() {
	flag.BoolVar(&f.count, "c", false, "count line appearances")
	flag.BoolVar(&f.repeated, "d", false, "show only repeated lines")
	flag.BoolVar(&f.uniq, "u", false, "show unique lines")
	flag.IntVar(&f.numFields, "f", 0, "ignore {num_fields} first fields of each line")
	flag.IntVar(&f.numChars, "s", 0, "ignore {num_chars} first chars of each line")
	flag.BoolVar(&f.register, "i", false, "ignore chars register")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		f.inputFile = args[0]
	}
	if len(args) > 1 {
		f.outputFile = args[1]
	}
}

func (f *Flags) Usage() {
	if (f.count && f.repeated) || (f.count && f.uniq) || (f.repeated && f.uniq) {
		fmt.Println("Usage:\nuniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
		os.Exit(1)
	}
}

func (f *Flags) Input() ([]string, error) {
	var input io.Reader
	if f.inputFile != "" {
		file, err := os.Open(f.inputFile)
		if err != nil {
			return nil, fmt.Errorf("failed to open input file: %v", err)
		}
		defer file.Close()
		input = file
	} else {
		input = os.Stdin
	}
	scanner := bufio.NewScanner(input)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading input failed: %v", err)
	}
	return lines, nil
}

func (f *Flags) Output(lines []string) error {
	var writer io.Writer
	if f.outputFile != "" {
		file, err := os.Create(f.outputFile)
		if err != nil {
			return fmt.Errorf("failed to create output file: %v", err)
		}
		defer file.Close()
		writer = file
	} else {
		writer = os.Stdout
	}
	bufWriter := bufio.NewWriter(writer)
	for _, line := range lines {
		if _, err := bufWriter.WriteString(line + "\n"); err != nil {
			return fmt.Errorf("failed to write to output file: %v", err)
		}
	}
	if err := bufWriter.Flush(); err != nil {
		return fmt.Errorf("failed to flush output file: %v", err)
	}
	return nil
}

func (f *Flags) flagsFSIProcessing(line string) string {
	fields := strings.Fields(line)
	if len(fields) <= f.numFields {
		return line
	}
	ignoredFields := strings.Join(fields[f.numFields:], " ")
	if len(ignoredFields) <= f.numChars {
		return ignoredFields
	}
	ignoredChars := ignoredFields[f.numChars:]
	if f.register {
		ignoredChars = strings.ToLower(ignoredChars)
	}
	return ignoredChars
}

func (f *Flags) flagsCDUProcessing(lines []string, counter map[string]int) []string {
	var result []string
	for _, line := range lines {
		key := f.flagsFSIProcessing(line)
		if f.count && counter[key] > 0 {
			result = append(result, fmt.Sprintf("%d %s", counter[key], line))
			delete(counter, key)
		}
		if f.repeated && counter[key] > 1 {
			delete(counter, key)
			result = append(result, line)
		}
		if f.uniq && counter[key] < 2 {
			result = append(result, line)
		}
	}
	return result
}

func (f *Flags) Processing(lines []string) []string {
	counter := make(map[string]int)
	seen := make(map[string]bool)
	var result []string
	for _, line := range lines {
		key := f.flagsFSIProcessing(line)
		if !seen[key] {
			seen[key] = true
			if !f.count && !f.repeated && !f.uniq {
				result = append(result, line)
			}
		}
		if seen[key] {
			counter[key]++
		}
	}
	if f.count || f.repeated || f.uniq {
		result = f.flagsCDUProcessing(lines, counter)
	}
	return result
}

func main() {
	flags := &Flags{}
	flags.Init()
	flags.Usage()
	lines, err := flags.Input()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := flags.Processing(lines)
	if err := flags.Output(result); err != nil {
		fmt.Println(err)
		return
	}
}
