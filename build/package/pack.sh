#!/usr/bin/env bash
set -euo pipefail

mkdir -p out
make DESTDIR=out/deb/cityu-cp37-keep-online prefix=/ exec_prefix=/usr install-deb
dpkg-deb --build out/deb/cityu-cp37-keep-online
