#!/bin/bash

go build

echo
echo "==========================="
echo "Running the test......"
echo
./bufferManager
echo "==========================="
echo
echo
echo "==========================="
echo "pprof - with bufferManager:"
echo
go tool pprof --text ./bufferManager ./with_bufferManager.prof
echo "==========================="
echo
echo
echo "==========================="
echo "pprof - without bufferManager:"
echo
go tool pprof --text ./bufferManager ./without_bufferManager.prof
echo "==========================="
