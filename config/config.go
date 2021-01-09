package config

type option struct {
	Flag        string
	Description string
}

type man struct {
	Name        string
	Command     string
	Description string
	Options     []option
}

var toolOptions = []option{
	{
		Flag:        "v",
		Description: "Enable verbosity",
	},
	{
		Flag:        "h",
		Description: "Print this message",
	},
}

// Config ...
var Config = man{
	Name:        "Neighborhood",
	Description: "A very simple portknocker!!",
	Command:     "neighborhood",
	Options:     toolOptions,
}
