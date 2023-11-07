package db_dialer

import (
	"golang.org/x/crypto/ssh"
	"net"
	"time"
)

func NewSSHDialer(sshClient *ssh.Client) *SSHClientDialer {
	return &SSHClientDialer{
		sshClient: sshClient,
	}
}

type SSHClientDialer struct {
	sshClient *ssh.Client
}

func (d *SSHClientDialer) Dial(network, address string) (net.Conn, error) {
	return d.sshClient.Dial(network, address)
}

func (d *SSHClientDialer) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	//_, cancel := context.WithTimeout(context.Background(), timeout)
	//defer cancel()
	return d.Dial(network, address)
}
