package costlogger

import (
	"net/http"

	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
)

const clientId = "78836484100-c7tmmk03qn4ujpk2b4ea434p4oqtomto.apps.googleusercontent.com"

var (
	scopes	  = []string{endpoints.EmailScope}
	clientIds = []string{clientId, endpoints.APIExplorerClientID}
	// in case we'll want to use TicTacToe API from an Android app
	audiences = []string{clientId}
)

type CostListReq struct {
	Limit int `json:"limit"`
}

type CostListResp struct {
	Items []string `json:"items"`
}

// CostLogger API service
type CostLoggerAPI struct {
}

// ScoresList queries scores for the current user.
// Exposed as API endpoint
func (cl *CostLoggerAPI) CostList(r *http.Request,
	req *CostListReq, resp *CostListResp) error {

	resp.Items = []string{"a", "b"}
	return nil
}

func RegisterService() (*endpoints.RPCService, error) {
	api := &CostLoggerAPI{}
	rpcService, err := endpoints.RegisterService(api,
			"costlogger", "v1", "Cost Logger API", true)
	if err != nil {
		return nil, err
	}

	info := rpcService.MethodByName("CostsList").Info()
	info.Path, info.HTTPMethod, info.Name = "costs", "GET", "costs.list"
	info.Scopes, info.ClientIds, info.Audiences = scopes, clientIds, audiences

	return rpcService, nil
}
