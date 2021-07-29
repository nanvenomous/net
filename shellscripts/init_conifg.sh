#!/usr/bin/env bash

exampleConfig=".config/net.yaml"
# TODO: this should be synced with config in cmd/root.go 
configLocation="${HOME}/.config/net.yaml"

if [ ! -f "${configLocation}" ]; then
	echo "Making config file at ${configLocation}"
	cp "${exampleConfig}" "${configLocation}"
else
	echo "Config file found"
fi
