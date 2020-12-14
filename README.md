## Running solutions

[![Go Report Card](https://goreportcard.com/badge/github.com/sshilin/aoc2020)](https://goreportcard.com/report/github.com/sshilin/aoc2020) [![License](https://img.shields.io/badge/license-MIT-brightgreen)](./LICENSE)

Both parts

    go test -v github.com/sshilin/aoc2020/day1

Specific test

    go test -v github.com/sshilin/aoc2020/day1 -run Test_part1

Benchmark

    go test -v github.com/sshilin/aoc2020/day1 -run=^$ -benchmem -bench BenchmarkPart1
