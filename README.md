# tetris-optimizer

tetris-optimizer is a Go program and a solution to an exercise of the same name in the 01-edu curriculum.

## usage

The program reads tetronominoes (tetris blocks) from a provided text file, arranges them in as small a square as possible and prints out the resulting solution.

The program can be used without compiling it:

```bash
go run . "sample.txt"
```

Compiling it to an executable is another option:

```bash
go build
./tetris sample.txt
```
