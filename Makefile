export GO111MODULE=on

all: frp_plugin_allowed_ports

frp_plugin_allowed_ports:
	go build -o ./bin/frp_plugin_allowed_ports ./cmd/frp_plugin_allowed_ports
