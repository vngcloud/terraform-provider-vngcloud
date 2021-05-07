package client

import (
	"github.com/vngcloud/terraform/client/authen"
	"github.com/vngcloud/terraform/client/vserver"
)

type Client struct {
	AuthenClient  *authen.AuthenClient
	VserverClient *vserver.APIClient
}

func NewClient(ClientID string, ClientSecret string, TokenURL string, baseURL string) (*Client, error) {
	authenConfig := authen.NewConfiguration(ClientID, ClientSecret, TokenURL)
	authenClient, err := authen.NewAuthenClient(authenConfig)
	if err != nil {
		return nil, err
	}
	vserverConfig := vserver.NewConfiguration(baseURL, authenClient.Client)
	vserverClient := vserver.NewAPIClient(vserverConfig)
	client := &Client{
		AuthenClient:  authenClient,
		VserverClient: vserverClient,
	}
	return client, nil
}
