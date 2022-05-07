go build -ldflags="-s -w" -o tmp-code.exe . && upx tmp-code.exe -9
# -H windowsgui 以下命令用于构建无需命令行启动的exe
# go build -ldflags="-s -w -H windowsgui" -o tmp-code.exe . && upx tmp-code.exe -9