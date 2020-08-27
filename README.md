# cronparser
parsing cron strings from cli

## Prerequisites

- go installed
- GNU make

## Description
This cli tool will parse a cron string e.g.
```
./cronparser "*/15 0 1,15 * 1-5 /usr/bin/find"
```
This tool will not parse single word crons like `@hourly`, `@monthly` etc..

## Run

The tool only uses the standard libary, so you don't need go.mod or any dependency management

- build the binary
```
make build
```

- run the tool and provide the cron string as the first argument
```
./cronparser "* * * * * /bin/true"
```
