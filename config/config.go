package config

var (
	FrontEndAddr      = ""
	BackEndAddr       = ""
	CryptoMethod      = ""
	LogTo             = ""
	ClientMode        = true
	VerboseTransproxy = false

	ProxyServerAddr      = ":8888"
	RegiServiceAddr      = ":1091"
	PacServiceAddr       = ":1092"
	RemotePacServiceAddr = ""
)

const (
	Version = "0.2.0"
)
