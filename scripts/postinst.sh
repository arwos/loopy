#!/bin/bash


if [ -f "/etc/systemd/system/loop.service" ]; then
    systemctl start loop
    systemctl enable loop
    systemctl daemon-reload
fi
