# Neighborhood
> Good neighbors knock on the door to enter

![screenshot](https://raw.githubusercontent.com/CasvalDOT/neighborhood/main/screenshot.png)

This tool is a small golang program to perform port knocking.

## Requires
[GO](https://golang.org/dl/) is required to develop or to compile source code. 
Otherwise you can simple download the latest release.

## Install
After cloning or downloading the repository, follow these istructions
```
go mod vendor
go build -ldflags "-s -w"
mv neighborhood /bin # or another binary folder
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

## Support the project
If you want, you can also freely donate to fund the project development!
[donate](https://paypal.me/FGortani)

## License

Copyright (c) Fabrizio Gortani

[MIT License](http://en.wikipedia.org/wiki/MIT_License)
