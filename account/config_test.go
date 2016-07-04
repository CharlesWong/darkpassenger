package account

import (
	"encoding/json"
	"testing"
)

func TestConfig(t *testing.T) {
	c := &Config{
		AdminToken: "token",
		DataFile:   "dp.dat",
		ListenAddr: ":9123",
	}
	b, _ := json.Marshal(c)
	t.Log(string(b))
}
