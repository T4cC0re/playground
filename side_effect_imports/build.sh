#!/usr/bin/env bash

go build -buildmode=plugin -o demo.so ../demo_plugin
go build
