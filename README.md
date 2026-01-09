# RPC forward 
RPC request forward utils

# Usage 

```go
func (ws *WebServer) GetStationList(ctx context.Context, req *webpb.GetStationListReq) (reply *webpb.GetStationListReply, err error) {
	reply, _, err = rpcforward.Call[
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