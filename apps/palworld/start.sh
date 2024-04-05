#!/bin/bash

set -ex

./app/PalServer.sh -port "$PORT" -useperfthreads -NoAsyncLoadingThread -UseMultithreadForDS
