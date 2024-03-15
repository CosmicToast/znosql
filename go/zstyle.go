package gozstyle

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// Client connection is just TCP
type Client struct {
	net.Conn
}

// NewZstyle initializes a Zstyle connection
func NewZstyle(dial string) *Client {
	conn, err := net.Dial("tcp", dial)
	if err != nil {
		return nil
	}

	var res Client
	res.Conn = conn

	return &res
}

// ReadLine reads a line. Use the real functions pls.
func (r *Client) ReadLine() (string, error) {
	reader := bufio.NewReader(r)
	s, e := reader.ReadString('\n')
	if e != nil {
		return "", e
	}
	return strings.TrimSuffix(s, "\n"), nil
}

// Exit closes the connection. NEVER FORGET TO DO THIS.
func (r *Client) Exit() {
	fmt.Fprint(r, "exit\n")
	r.Close()
}

// Ping checks for health
func (r *Client) Ping() bool {
	fmt.Fprint(r, "ping\n")
	s, e := r.ReadLine()
	if e != nil {
		return false
	}
	if s != "pong" {
		return false
	}
	return true
}

// Get fetches a key with no default value (empty)
func (r *Client) Get(key string) (string, error) { return r.Getd(key, "") }

// Getd fetches a key with a default value
func (r *Client) Getd(key, defval string) (string, error) {
	query := fmt.Sprintf("getd %s %s\n", key, defval)
	fmt.Fprint(r, query)

	s, e := r.ReadLine()
	if e != nil {
		return "", e
	}
	return s, nil
}

// Put puts a value into a key
func (r *Client) Put(key, val string) {
	query := fmt.Sprintf("put %s %s\n", key, val)
	fmt.Fprint(r, query)
}

func (r *Client) Save() {
	fmt.Fprint(r, "save")
}

func (r *Client) Shutdown() {
	fmt.Fprint(r, "shutdown")
	r.Close()
}
