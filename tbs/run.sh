#!/bin/sh
rm tbs.exe
go build -o tbs.exe
tbs -conf=tbs.toml