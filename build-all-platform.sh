rm -rf ./build

# change this value to the version you want to build
version="1.0.0"
author="Serghei Gurgurov"
license="ISC"
keywords="[\"onion-kv\", \"key-value\", \"store\", \"database\"]"

cd ./src/

build_platform() {
    env GOOS=$goos GOARCH=$goarch go build -o ../build/$packageos-$packagearch/bin/onion-kv$extension
    cat ../LICENSE > ../build/$packageos-$packagearch/LICENSE
    echo "{}" |
    jq "
      .version=\"${version}\"
    | .name=\"onion-kv-${packageos}-${packagearch}\"
    | .description=\"onion-kv binary for ${packageos} ${packagearch}\"
    | .repository={
        \"type\": \"git\",
        \"url\": \"https://github.com/SergheiGurgurov/onion-kv.git\"
    }
    | .author=\"${author}\"
    | .keywords=${keywords}
    | .os[0]=\"${packageos}\"
    | .cpu[0]=\"${packagearch}\"
    | .license=\"${license}\"
    " > ../build/$packageos-$packagearch/package.json
}

goarch="amd64"
goos="linux"
packagearch="x64"
packageos="linux"
extension=""

build_platform

goarch="arm64"
goos="linux"
packagearch="arm64"
packageos="linux"
extension=""

build_platform

goarch="amd64"
goos="darwin"
packagearch="x64"
packageos="darwin"
extension=""

build_platform

goarch="arm64"
goos="darwin"
packagearch="arm64"
packageos="darwin"
extension=""

build_platform

goarch="amd64"
goos="windows"
packagearch="x64"
packageos="win32"
extension=".exe"

build_platform

cd ..
echo '{
  "name": "onion-kv",
  "description": "a simple key-value store written in golang",
  "bin": {
    "onion-kv": "bin/cli"
  },
  "repository":{
    "type": "git",
    "url": "https://github.com/SergheiGurgurov/onion-kv.git"
  },
  "os": [
    "darwin",
    "linux",
    "win32"
  ],
  "keywords": [
    "onion-kv",
    "key-value",
    "store",
    "database"
  ],
  "cpu": [
    "arm64",
    "x64"
  ],
  "author": "Serghei Gurgurov",
  "license": "ISC"
}' |
jq "
.version=\"${version}\" |
.optionalDependencies={
    \"onion-kv-linux-x64\": \"${version}\",
    \"onion-kv-linux-arm64\": \"${version}\",
    \"onion-kv-darwin-x64\": \"${version}\",
    \"onion-kv-darwin-arm64\": \"${version}\",
    \"onion-kv-win32-x64\": \"${version}\"
}
" > ./build/package.json

cat ./LICENSE > ./build/LICENSE
mkdir ./build/bin
cat ./cli.js > ./build/bin/cli.js

echo "
darwin-arm64
darwin-x64
linux-arm64
linux-x64
win32-x64
" > ./build/.npmignore