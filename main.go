package main

import (
	"flag"
	"fmt"
	"neighborhood/config"
	"neighborhood/helpers"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/*
Probe the host with the protocol and the port provided
in the sequence
*/
func probe(host string, proto string, port int) error {
	d := net.Dialer{
		Timeout: 200 * time.Millisecond,
	}
	c, err := d.Dial(proto, host+":"+strconv.Itoa(port))
	if err != nil {
		return err
	}
	defer c.Close()
	return nil
}

/*
User can pass a sequence of ports to probe.
The port can be followed by the protocol
in the following format:
<PORT>:<TCP | UDP>
*/
func parsePort(portArgument string) (string, int) {
	splittedPortArgument := strings.Split(portArgument, ":")
	if len(splittedPortArgument) == 0 {
		fmt.Println("INVALID PORT PARAMETER")
		os.Exit(1)
	}

	var proto string = "TCP"
	var port int
	if len(splittedPortArgument) > 1 {
		proto = splittedPortArgument[1]
	}

	port, err := strconv.Atoi(splittedPortArgument[0])
	if err != nil {
		fmt.Println("INVALID PORT NUMBER", splittedPortArgument[0])
		os.Exit(1)
	}

	return strings.ToLower(proto), port
}

func parseArguments(args []string) []string {
	var expectedFlags []string
	for _, option := range config.Config.Options {
		expectedFlags = append(expectedFlags, option.Flag)
	}

	joinedFlags := strings.Join(expectedFlags, "|")

	argsString := strings.Join(args, ",")
	regexString := "(\\s|,|)-(" + joinedFlags + ")(\\s|,|$|)"

	rgx := regexp.MustCompile(regexString)
	argsWithoutFlags := rgx.ReplaceAllString(argsString, "")

	return strings.Split(argsWithoutFlags, ",")
}

/*
Main function
*/
func main() {
	arguments := os.Args

	isVerbose := flag.Bool("v", false, "")
	showHelp := flag.Bool("h", false, "")
	flag.Parse()

	if *showHelp == true {
		helpers.Man()
		os.Exit(0)
	}

	args := parseArguments(arguments[1:])
	if len(args) < 2 {
		fmt.Println("Invalid arguments length!\nEnsure to provide a valid host and a valid sequence of ports")
		os.Exit(1)
	}

	host := args[0]
	portSequence := args[1:]

	for _, p := range portSequence {
		proto, port := parsePort(p)
		probe(host, proto, port)
		if *isVerbose == true {
			fmt.Println("PROBE:", proto, "->", host, ":", port)
		}
	}

}
