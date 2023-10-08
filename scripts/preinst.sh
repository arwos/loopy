#!/bin/bash


if ! [ -d /var/lib/loop/ ]; then
    mkdir /var/lib/loop
fi

if [ -f "/etc/systemd/system/loop.service" ]; then
    systemctl stop loop
    systemctl disable loop
    systemctl daemon-reload
fi

if ! [ -d /var/lib/loopcli/ ]; then
    mkdir /var/lib/loopcli
fi

if [ -f "/etc/systemd/system/loopcli.service" ]; then
    systemctl stop loopcli
    systemctl disable loopcli
    systemctl daemon-reload
fi
