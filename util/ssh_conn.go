package util

import (
	"net"
	"time"
)

type SSHChanConn struct {
	Conn net.Conn
}

func (t *SSHChanConn) Read(b []byte) (n int, err error) {
	return t.Conn.Read(b)
}

func (t *SSHChanConn) Write(b []byte) (n int, err error) {
	return t.Conn.Write(b)
}

func (t *SSHChanConn) Close() error {
	return t.Conn.Close()
}

func (t *SSHChanConn) LocalAddr() net.Addr {
	return t.Conn.LocalAddr()
}

func (t *SSHChanConn) RemoteAddr() net.Addr {
	return t.Conn.RemoteAddr()
}

func (t *SSHChanConn) SetDeadline(deadline time.Time) error {
	if err := t.SetReadDeadline(deadline); err != nil {
		return err
	}
	return t.SetWriteDeadline(deadline)
}

func (t *SSHChanConn) SetReadDeadline(deadline time.Time) error {
	return nil
}

func (t *SSHChanConn) SetWriteDeadline(deadline time.Time) error {
	return nil
}
