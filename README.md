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

## Usage

```
neighborhood -v <host> <port:protocol> <port:protocol>

# example of sequence:
# 80:tcp 81:udp
# 80 81:tcp

```
