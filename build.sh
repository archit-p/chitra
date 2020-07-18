#!/bin/bash

# configure the directories
BUILD_BASE=build
CLIENT_BASE=chitra-client

# create absolute paths
SOURCE_DIR=($(pwd))
BUILD_DIR="${SOURCE_DIR}/${BUILD_BASE}"
CLIENT_DIR="${SOURCE_DIR}/${CLIENT_BASE}"

# go source file
SRC=Chitra.go

# declare GOOS as linux
GOOS=linux

# build the go module
go build ${SOURCE_DIR}/${SRC}

# build the react frontend
npm run build --prefix ${CLIENT_DIR}

# delete the build directory if it exists
if [ -d "${BUILD_DIR}" ]; then
	rm -rf ${BUILD_DIR}
fi

# move to build dir
mkdir ${BUILD_DIR}
cp ${SOURCE_DIR}/Chitra ${BUILD_DIR}
cp ${CLIENT_DIR}/build ${BUILD_DIR}/client -r

echo "All build files at ${BUILD_DIR}"
echo "Run ${BUILD_DIR}/Chitra to start server"
