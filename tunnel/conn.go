package tunnel

import (
	"github.com/CharlesWong/darkpassenger/crypto"
	"net"
	"time"
)

type Conn struct {
	conn   net.Conn
	cipher *crypto.Cipher
	pool   *recycler
}

func NewConn(conn net.Conn, cipher *crypto.Cipher, pool *recycler) *Conn {
	return &Conn{
		conn:   conn,
		cipher: cipher,
		pool:   pool,
	}
}

func (c *Conn) Read(b []byte) (int, error) {
	c.conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	if c.cipher == nil {
		return c.conn.Read(b)
	}
	n, err := c.conn.Read(b)
	if n > 0 {
		c.cipher.Decrypt(b[0:n], b[0:n])
	}
	return n, err
}

func (c *Conn) Write(b []byte) (int, error) {
	c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if c.cipher == nil {
		return c.conn.Write(b)
	}
	c.cipher.Encrypt(b, b)
	return c.conn.Write(b)
}

func (c *Conn) Close() {
	c.conn.Close()
}

func (c *Conn) CloseRead() {
	if conn, ok := c.conn.(*net.TCPConn); ok {
		conn.CloseRead()
	}
}

func (c *Conn) CloseWrite() {
	if conn, ok := c.conn.(*net.TCPConn); ok {
		conn.CloseWrite()
	}
}
