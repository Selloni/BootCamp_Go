#!/bin/bush

go test -bench . -benchmem -cpuprofile=profile.pro
go tool pprof profile.pro

rm integration*
rm profile.pro