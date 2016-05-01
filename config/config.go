package config

var (
	FrontEndAddr      = ""
	BackEndAddr       = ""
	CryptoMethod      = ""
	LogTo             = ""
	ClientMode        = true
	VerboseTransproxy = false

	RegiServiceAddr      = ":1091"
	PacServiceAddr       = ":1092"
	RemotePacServiceAddr = ""
)

const (
	Version = "0.2.0"
)
