package dokku

import (
	"crypto/rsa"
	"github.com/texm/dokku-go"
	"golang.org/x/crypto/ssh"
)

type Config struct {
	DebugMode bool

	PrivateKey *rsa.PrivateKey

	Host string
	Port string
}

func Init(cfg Config) (*dokku.SSHClient, error) {
	dCfg := &dokku.SSHClientConfig{
		Host:       cfg.Host,
		Port:       cfg.Port,
		PrivateKey: cfg.PrivateKey,
		// TODO: supply host key / actually check it
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return dokku.NewSSHClient(dCfg)
}
