# RPC forward 
RPC request forward utils

# Usage 

```go
package web
import (
     "github.com/civet148/rpcforw"
     "web/pbs/webpb"
     "web/pbs/corepb"
)
type WebServer struct {
     CoreCli *corepb.CoreServiceClient
}
func (ws *WebServer) GetStationList(ctx context.Context, req *webpb.GetStationListReq) (reply *webpb.GetStationListReply, err error) {
	reply, _, err = rpcforw.Call[
		webpb.GetStationListReq,
		webpb.GetStationListReply,
		corepb.GetStationListReq,
		corepb.GetStationListReply,
	](ctx, req, ws.CoreCli.GetStationList)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
```