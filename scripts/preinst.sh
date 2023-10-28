#!/bin/bash


if ! [ -d /var/lib/loop/ ]; then
    mkdir /var/lib/loop
fi

if [ -f "/etc/systemd/system/loop.service" ]; then
    systemctl stop loop
    systemctl disable loop
    systemctl daemon-reload
fi
