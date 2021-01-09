# Neighborhood
> Good neighbors knock on the door to enter

This tool is a small golang program to perform port knocking.

## Requires
Golang 1.15 +

## Install
After cloning or downloading the repository, follow these istructions
```
go mod vendor
go build -ldflags "-s -w"
cp neighborhood /bin # or another binary folder
```

You can also download the latest release binary.

## Usage

```
neighborhood -v <host> <port:protocol> <port:protocol>

# example of sequence:
# 80:tcp 81:udp
# 80 81:tcp

```
## Have you found a bug?

Please open a new issue on:

https://github.com/CasvalDOT/neighborhood/issues

## License

Copyright (c) Fabrizio Gortani

[MIT License](http://en.wikipedia.org/wiki/MIT_License)
