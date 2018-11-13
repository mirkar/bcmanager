package prometheus

type Apache struct {
	// environment (production, qa etc)
	Env string `json:"env"`

	// data center where a target is resided
	Datacenter string `json:"datacenter"`

	// host name
	Host string `json:"host"`

	// port number apache instance is running on
	// convert output to a string and rename "ingress"
	Ingress int `json:"ingress, string"`

	// account an apache instance is running under
	Acc string `json:"acc"`

	// the role this apache instance is playing in the environment (wca-cluster-ssl, bsft-trinity etc)
	Role string `json:"role"`

	// IP address and port colon separated combination. IP - host IP,  port - the port apache
	// exporter is running on  (for example 130.4.170.16:9118)
	Instance string `json:"instance"`

	// it is going to be 'apache' always, just as example do not output this field
	Job string `json:"-"`

	// the name is lowercase so is not available outside the package (just for demo purpose)
	password string

	// do not output the field if the value is empty
	InstallPath string `json:"installpath,omitempty"`

	EmptyExample string `json:"expectedempty"`
}
