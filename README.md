# Pex code challenge: Senior Go Engineer

By Michal Gondar, Nov 2020

The optimal number of workers will depend on the hardware of the machine on which this program will run. 


## Assignment:
`https://gist.github.com/ehmo/e736c827ca73d84581d812b3a27bb132`

## Assumptions
- Multiple downloads of the same URL are acceptable. The code does not perform duplicity checks.
  - Considering that input can be "more than a billion URLs" and this program will run on a machine with limited
    resources, duplicity checks should be done outside of this program, if desired.
- Download failures are ignored. No retries are performed in the code.
  - `DownloadImageWithRetry` can be used for a limited number of retries, if desired.
  - Failed attempts (except 404) can be pushed back to the `urls` channel for retry.
- If less than 3 colors are extracted, the CSV file is filled with empty strings to preserve the CSV format.

## Build

`go build -o bin/pex cmd/pex/main.go`

## Usage

`bin/pex data/input.txt data/output.csv`
`bin/pex -d -n=200 data/input.txt data/output.csv`

## Tests

`go test ./...`

## Future optimizations:
- Spawn workers per host (extracted from url). 

## Profiling

```
func main() {
	// CPU profiling by default
	//defer profile.Start(profile.ProfilePath("profiling")).Stop()

	// Memory profiling
	//defer profile.Start(profile.ProfilePath("profiling"), profile.MemProfile).Stop()

    ...
}
```

`go tool pprof --pdf profiling/cpu.pprof > profiling/cpu.pdf`
`go tool pprof --pdf profiling/mem.pprof > profiling/mem.pdf`
