del gyuttoe-rp.exe
del gyuttoe-rp-release.exe

go build -ldflags "-s -w -H=windowsgui" -tags release 

upx --best -o gyuttoe-rp-release.exe gyuttoe-rp.exe

pause