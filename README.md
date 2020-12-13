## Running solutions

Both parts

    go test -v github.com/sshilin/aoc2020/day1

Specific test

    go test -v github.com/sshilin/aoc2020/day1 -run Test_part1

Benchmark

    go test -v github.com/sshilin/aoc2020/day1 -run=^$ -benchmem -bench BenchmarkPart1