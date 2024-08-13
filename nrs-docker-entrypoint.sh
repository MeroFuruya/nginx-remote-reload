#!/bin/sh

set -e

if [ -z "$NRS_PORT" ]; then
  export NRS_PORT=5000
fi
echo "NRS_PORT: $NRS_PORT" 
nginx-remote-signal &
