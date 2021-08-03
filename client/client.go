package client

import (
	"github.com/vngcloud/terraform/client/authen"
	"github.com/vngcloud/terraform/client/vdb"
	"github.com/vngcloud/terraform/client/vserver"
)

type Client struct {
	AuthenClient  *authen.AuthenClient
	VserverClient *vserver.APIClient
	VdbClient     *vdb.APIClient
	ProjectId     string
}

func NewClient(baseURL string, projectId string, userId string, ClientID string, ClientSecret string, TokenURL string) (*Client, error) {
	authenConfig := authen.NewConfiguration(ClientID, ClientSecret, TokenURL)
	authenClient, err := authen.NewAuthenClient(authenConfig)
	if err != nil {
		return nil, err
	}

	vserverConfig := vserver.NewConfiguration("https://vserverapi.vngcloud.vn/vserver-gateway", authenClient.Client)
	vserverClient := vserver.NewAPIClient(vserverConfig)

	vdbConfig := vdb.NewConfiguration("https://vcmc-internal.vngcloud.tech/vdb-gateway", authenClient.Client)
	vdbClient := vdb.NewAPIClient(vdbConfig)

	vdbConfig.AddDefaultHeader("user_id", userId)

	client := &Client{
		AuthenClient:  authenClient,
		VserverClient: vserverClient,
		VdbClient:     vdbClient,
		ProjectId:     projectId,
	}
	return client, nil
}
