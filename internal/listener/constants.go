package listener

const (
	// EnvPort is the key of the default port for the server
	EnvPort = "PORT"
)

var (
	// Port is the default port for the server
	Port string
)

// Load loads the listener constants
func Load() {
	// Load the port
	if err := Loader.LoadVariable(
		EnvPort,
		&Port,
	); err != nil {
		panic(err)
	}
}
