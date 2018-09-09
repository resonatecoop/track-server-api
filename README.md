# Track Server API

The track-server-api delivers streaming track downloads, and accepts track uploads. 

There is no documentation yet, except for this README. 

The following commands install the track-server-api in a temporary test directory and run all existing tests:

```
export GOPATH=/tmp/test
mkdir -p /tmp/test/src
cd /tmp/test/src
git clone git@gitlab.resonate.ninja:resonate/track-server-api.git
cd track-server-api
dep ensure
retool build
retool do protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. ./rpc/service.proto
go run internal/database/migrations/*.go
go test -v track-server-api/internal/server --ginkgo.v="true"
```
