 #!/usr/bin/env bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )


export GOPATH="$SCRIPT_DIR/curator-back"

echo "Set GOPATH env var to $GOPATH"