# Word-Guesser

A command-line tool to help you solve Wordle-style word puzzles by filtering and ranking possible words based on your clues.

## Features

- Filter words by length
- Exclude specific letters
- Specify known letter positions
- Specify letters present but in the wrong position
- Ranks candidates by unique letter count and letter frequency

## Usage

You can run the tool directly using Go without cloning the repository:

```sh
go run github.com/nilspolek/Word-Guesser@latest
```

### Options

- `-f <path>`: Path to a custom dictionary file (optional)
- `-l <length>`: Length of the words to guess (default: 5)

### Example

```sh
go run github.com/nilspolek/Word-Guesser@latest -l 6
```

or with a custom dictionary:

```sh
go run github.com/nilspolek/Word-Guesser@latest -f mywords.txt
```

## How to Use

1. Start the program.
2. When prompted, enter:
   - Letters to exclude (e.g. `abc`)
   - Letters in their correct positions (e.g. `a__c_`)
   - Letters in the word but not in the right position (e.g. `a_b__`)
3. The program will display the best candidate words.

## Requirements

- Go 1.24.3 or newer