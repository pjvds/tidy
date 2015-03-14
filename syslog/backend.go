package syslog

//import(sys "log/syslog")
//Only one call to Dial is necessary. On write failures, the syslog client will attempt to reconnect to the server and write again.

type backend struct {
	// The network (udp, tcp), or empty for local syslog socket. If the network is set
	// you also need to set the address.
	Network string
	// The remote address, or empty for local syslog socket. If the network is set
	// you also need to set the address.
	Address string

	// The TAG field value for the syslog entries. If empty the process name is used.
	Tag string
}
