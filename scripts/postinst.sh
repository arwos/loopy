#!/bin/bash


if [ -f "/etc/systemd/system/loop.service" ]; then
    systemctl start loop
    systemctl enable loop
    systemctl daemon-reload
fi

if [ -f "/etc/systemd/system/loopcli.service" ]; then
    systemctl start loopcli
    systemctl enable loopcli
    systemctl daemon-reload
fi
