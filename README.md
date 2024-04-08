# onion-kv

this is a simple key-value store written in golang

# installation

### npm (coming soon)

```sh
npm install -g onion-kv
```

### build for yourself (golang compiler required)

```sh
git clone https://github.com/SergheiGurgurov/onion-kv.git
./build.sh
# the binary will be installed in ~/.local/bin/onion-kv
```

# usage

```sh
# launch db, it will create all the necessary files in the current directory
onion-kv
```

to connect to the db you can use the **onion-kv-client** package from **npm**, documentation for building a custom client will be released in the future.
