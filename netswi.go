package netswi

import (
	"net"
	"net/http"
	"time"
)

func NewClientBindedToIP(ip string) (*http.Client, error) {
	if trans, err := NewTransportBindedToIP(ip); err != nil {
		return nil, err
	} else {
		return &http.Client{Transport: trans}, nil
	}
}

func NewTransportBindedToIP(ip string) (*http.Transport, error) {
    ip_addr, err := net.ResolveIPAddr("ip", ip)
	if err != nil { return nil, err }

	trans := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			LocalAddr: &net.TCPAddr{IP: ip_addr.IP},
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns: 100,
		IdleConnTimeout: 90 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	return trans, nil
}
