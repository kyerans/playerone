CURRENT_DIR=$(pwd)
GOPATH_DIR=${GOPATH}

if [ -z "$GOPATH_DIR" ]; then
  echo "GOPATH is not set. Please set it before running this script."
  exit 1
fi

if [[ "$CURRENT_DIR" != *"$GOPATH_DIR"* ]]; then
  echo "Current directory ($CURRENT_DIR) does not contain GOPATH ($GOPATH_DIR)."
  exit 1
fi

GEN_PATH=$GOPATH/src/github.com/kyerans/playerone
GEN_OUT=$GOPATH/src

protoc \
  -I"$GEN_PATH" \
  --grpc-gateway_out="$GEN_OUT" \
  --go_out="$GEN_OUT" \
  --go-grpc_out=require_unimplemented_servers=false:"$GEN_OUT" \
  --validate_out="lang=go,paths=:$GEN_OUT" \
  "$(pwd)"/*.proto || exit 1
