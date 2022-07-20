#!/bin/bash
export NODE_ENV=development
rm -rf build
mkdir -p build/ui

cd server
env GOOS=windows GOARCH=amd64 go build -o ../build/retro.exe
go build -o ../build/retro
chmod +x ../build/retro
cd ..
./build/retro &

cd frontend
npm run serve

killall retro