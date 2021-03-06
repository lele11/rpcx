// +build kcp

package server

import (
	"errors"
	"net"

	kcp "github.com/xtaci/kcp-go"
)

func init() {
	makeListeners["kcp"] = kcpMakeListener
}

func kcpMakeListener(s *Server, address string) (ln net.Listener, err error) {
	if s.Options == nil || s.Options["BlockCrypt"] == nil {
		return nil, errors.New("KCP BlockCrypt must be configured in server.Options")
	}

	return kcp.ListenWithOptions(address, s.Options["BlockCrypt"].(kcp.BlockCrypt), 10, 3)
}
