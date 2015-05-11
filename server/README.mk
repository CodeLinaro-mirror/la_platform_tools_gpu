# server
--
    import "android.googlesource.com/platform/tools/gpu/server"

Package server implements the rpc gpu debugger service, queriable by the
clients, along with some helpers exposed via an http listener.

## Usage

#### func  Run

```go
func Run(config Config, rpcReady chan<- struct{})
```
Run listens on the HTTP and RPC TCP ports given in config, initializes the
resource database, the replay manager and handles HTTP and RPC requests on
incoming connections. If not nil, the rpcReady channel is closed as soon as
incoming RPC requests can start being issued.

#### type Config

```go
type Config struct {
	HttpAddress string
	RpcAddress  string
	DataPath    string
	LogfilePath string
}
```
