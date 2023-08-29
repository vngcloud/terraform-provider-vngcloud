package client

import (
	"github.com/vngcloud/terraform-provider-vngcloud/client/authen"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vdb"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vloadbalancing"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vserver"
)

type Client struct {
	AuthenClient  *authen.AuthenClient
	VserverClient *vserver.APIClient
	VdbClient     *vdb.APIClient
	VlbClient     *vloadbalancing.APIClient
	ProjectId     string
}

func NewClient(vdbBaseURL string, vserverBaseURL string, vlbBaseURL string, projectId string, userId string, ClientID string, ClientSecret string, TokenURL string) (*Client, error) {
	authenConfig := authen.NewConfiguration(ClientID, ClientSecret, TokenURL)
	authenClient, err := authen.NewAuthenClient(authenConfig)
	if err != nil {
		return nil, err
	}

	vserverConfig := vserver.NewConfiguration(vserverBaseURL, authenClient.Client)
	vserverClient := vserver.NewAPIClient(vserverConfig)

	vlbConfig := vloadbalancing.NewConfiguration(vlbBaseURL, authenClient.Client)
	vlbClient := vloadbalancing.NewAPIClient(vlbConfig)

	vdbConfig := vdb.NewConfiguration(vdbBaseURL, authenClient.Client)
	vdbClient := vdb.NewAPIClient(vdbConfig)

	vdbConfig.AddDefaultHeader("user_id", userId)

	client := &Client{
		AuthenClient:  authenClient,
		VserverClient: vserverClient,
		VdbClient:     vdbClient,
		VlbClient:     vlbClient,
		ProjectId:     projectId,
	}
	return client, nil
}

func NewClientV2(vserverBaseURL string, vlbBaseURL string, ClientID string, ClientSecret string, TokenURL string) (*Client, error) {
	authenConfig := authen.NewConfiguration(ClientID, ClientSecret, TokenURL)
	authenClient, err := authen.NewAuthenClient(authenConfig)
	if err != nil {
		return nil, err
	}

	vserverConfig := vserver.NewConfiguration(vserverBaseURL, authenClient.Client)
	vserverClient := vserver.NewAPIClient(vserverConfig)

	vlbConfig := vloadbalancing.NewConfiguration(vlbBaseURL, authenClient.Client)
	vlbClient := vloadbalancing.NewAPIClient(vlbConfig)

	client := &Client{
		AuthenClient:  authenClient,
		VserverClient: vserverClient,
		VdbClient:     nil,
		VlbClient:     vlbClient,
		ProjectId:     "projectId",
	}
	return client, nil
}
