#!/bin/bash

set -ex

./PalServer.sh -port "$PORT" -useperfthreads -NoAsyncLoadingThread -UseMultithreadForDS "$@"
