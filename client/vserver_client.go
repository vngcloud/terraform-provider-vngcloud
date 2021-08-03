package client

import (
	"github.com/vngcloud/terraform/client/authen"
	"github.com/vngcloud/terraform/client/vserver"
)

type VSRClient struct {
	AuthenClient  *authen.AuthenClient
	VserverClient *vserver.APIClient
}

func NewVServerClient(ClientID string, ClientSecret string, TokenURL string, baseURL string) (*VSRClient, error) {
	authenConfig := authen.NewConfiguration(ClientID, ClientSecret, TokenURL)
	authenClient, err := authen.NewAuthenClient(authenConfig)
	if err != nil {
		return nil, err
	}
	vserverConfig := vserver.NewConfiguration(baseURL, authenClient.Client)
	vserverClient := vserver.NewAPIClient(vserverConfig)
	client := &VSRClient{
		AuthenClient:  authenClient,
		VserverClient: vserverClient,
	}
	return client, nil
}
