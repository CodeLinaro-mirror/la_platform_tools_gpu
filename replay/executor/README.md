# executor
--
    import "android.googlesource.com/platform/tools/gpu/replay/executor"

Package executor contains the Execute function for sending a replay to a device.

## Usage

#### func  Execute

```go
func Execute(
	payload protocol.Payload,
	decoder builder.ResponseDecoder,
	connection io.ReadWriteCloser,
	database database.Database,
	logger log.Logger,
	architecture device.Architecture) error
```
Execute sends the replay payload for execution on the target replay device
communicating on connection. decoder will be used for decoding all postback
reponses. Once a postback response is decoded, the corresponding handler in the
handlers map will be called.
