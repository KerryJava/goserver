#!/bin/bash
go test -test.bench=".*" -count=5 -cpuprofile=cpu.profile .
#go tool pprof --text mybin http://0.0.0.0:8082:/debug/pprof/profile
