# protoc-gen-enum-desc (pged)

**inspired by [PGV](https://github.com/bufbuild/protoc-gen-validate)**

Generate custom description for proto enum

### Install

```shell
go install github.com/windmeup/protoc-gen-enum-desc/cmd/protoc-gen-enum-desc@latest
```

# Usage

```protobuf
import "pged/enumdesc.proto"; // proto files in GO_HOME/pkg/mode/github.com/windmeup/...
enum ClientEvent {
  option (pged.default_description) = "unknown client event"; // default description, optional
  CLIENT_EVENT_UNSPECIFIED = 0;
  CLIENT_EVENT_START = 1 [(pged.description) = "client start"]; // description, optional
  CLIENT_EVENT_STOP = 2;
}
```

## Generate code file with protoc

```shell
protoc \
  -I . \
  -I path/to/pged/ \
  --go_out=":../generated" \
  --enum-desc_out="lang=go:../generated" \
  example.proto
```

## Generate code with [buf](https://buf.build/), recommend

```yaml
# buf.yaml
version: v1
deps:
  - buf.build/windmeup/pged
breaking:
  use:
    - FILE
lint:
  use:
    - DEFAULT
```

```yaml
# buf.gen.yaml
version: v1
plugins:
  - name: go
    out: .
    opt: paths=source_relative
  - name: enum-desc
    out: .
    opt:
      - paths=source_relative
      - lang=go
```

```shell
buf mod update
buf generate
```

# Golang

- go module

```shell
  go get github.com/windmeup/protoc-gen-enum-desc
```

- get enum's description with Description() methods, which in generated files xxxxxx.pb.enumdesc.go

```go
package example

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_ClientEvent_Description(t *testing.T) {
	require.Equal(t, "client start", Client_CLIENT_EVENT_START.Description())
}
```

# Other languages

TODO
