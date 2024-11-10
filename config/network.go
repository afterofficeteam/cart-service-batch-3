package config

import "net"

func NetworkListener(network, address string) (net.Listener, error) {
	listen, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}

	return listen, nil
}
