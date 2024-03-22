#!/bin/sh

# start dlv server on port 4000
/go/bin/dlv debug \
    ./cmd/server \
    --continue \
    --listen=:4000 \
    --headless=true \
    --log=true \
    --accept-multiclient \
    --api-version=2 \
    --output ./__debug_bins/server &

# wait for code changes to rebuild dlv's code
watchexec -p -e go dlv connect :4000 --init init-dlv.txt
