cd ./build/

publish_platform(){
    npm publish ./$packageos-$packagearch/
}

goarch="amd64"
goos="linux"
packagearch="x64"
packageos="linux"
extension=""

publish_platform

goarch="arm64"
goos="linux"
packagearch="arm"
packageos="linux"
extension=""

publish_platform

goarch="amd64"
goos="darwin"
packagearch="x64"
packageos="darwin"
extension=""

publish_platform

goarch="arm64"
goos="darwin"
packagearch="arm"
packageos="darwin"
extension=""

publish_platform

goarch="amd64"
goos="windows"
packagearch="x64"
packageos="win32"
extension=".exe"

publish_platform

npm publish
