package internal

type CLI struct{}

func (cli *CLI) ParseArgs(args []string) {
    // Implement argument parsing
}

var ArgDefOptions = map[string]map[string]string{
    "verbose": {"short": "v", "desc": ""},
    "debug":   {"short": "d", "desc": ""},
    "quiet":   {"short": "q", "desc": ""},
    "dryrun":  {"short": "dr", "desc": ""},
}
