# Local Serve - Simple http server

This is a simple http server for local testing and quick experiments.

Uses go modules, so using go > 1.11:

```
$ git clone https://github.com/matheusd/localserve
$ cd localserve
$ go install .
```

## Running

Just run `localserve` on the directory you want to serve over http. Use `-h` to see options.

## Simulate Long Running Op

Localserve can simulate a long-running operation by accessing a special route: `/__sleep__?duration=10s`. This will cause the server to sleep for 10 seconds, then return.

You can use any duration-like string to specify the delay.

