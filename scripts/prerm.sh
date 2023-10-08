#!/bin/bash


if [ -f "/etc/systemd/system/loop.service" ]; then
    systemctl stop loop
    systemctl disable loop
    systemctl daemon-reload
fi

if [ -f "/etc/systemd/system/loopcli.service" ]; then
    systemctl stop loopcli
    systemctl disable loopcli
    systemctl daemon-reload
fi
