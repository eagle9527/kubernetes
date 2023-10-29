package portforward

type Server struct {
	PortForward []PortForward `mapstructure:"PortForward" json:"PortForward" yaml:"PortForward"`
}

type PortForward struct {
	Name       string `yaml:"Name"`
	Resource   string `yaml:"Resource"`
	LocalPort  int    `yaml:"LocalPort"`
	RemotePort int    `yaml:"RemotePort"`
}
