#!/bin/bash


if [ -f "/etc/systemd/system/loop.service" ]; then
    systemctl stop loop
    systemctl disable loop
    systemctl daemon-reload
fi
