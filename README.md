# Chitra
A simple Go app to expose a user-specified directory for media-streaming over the local network. A simple frontend web app can be found at [Chitra-client](https://github.com/archit-p/chitra-client).

## Submodules
This repository includes the client app as a submodule.
```bash
git submodule sync
git submodule update --init
```

## Building
An included script `build.sh` can be used to build the server along with the sample frontend app.
```bash
./build.sh
```
Running this creates a `build/` dir and moves the client app to `build/client/' dir.

## Running
The app can be run from inside the build directory.
```bash
./build/Chitra <options>
```

## Configuration Options
The app includes configuration options to choose how it functions.
```bash
-sport		: port to start the server on
-cport		: port to serve the client app on
-sdir		: directory to server media from
-cdir		: directory for the client app
```

