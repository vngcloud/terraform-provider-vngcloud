package client

/*import (
	"github.com/vngcloud/terraform/client/authen"
	"github.com/vngcloud/terraform/client/vdb"
)

type VDBClient struct {
	vDatabaseClient *vdb.APIClient
	ProjectId string
}

func NewVDBClient(baseURL string, projectId string, userId string, clientID string, clientSecret string, tokenURL string) (*VSRClient, error) {
	// for local testing, use this to initialize vdbConfig
	//vdbConfig := vdb.NewConfiguration(baseURL, nil)
	authenConfig := authen.NewConfiguration(clientID, clientSecret, tokenURL)
	authenClient, err := authen.NewAuthenClient(authenConfig)
	if err != nil {
		return nil, err
	}
	vdbConfig := vdb.NewConfiguration(baseURL, authenClient.Client)
	vdbConfig.AddDefaultHeader("user_id/home/lap13015/Documents/GitProject/vdb-terraform", userId)
	vDatabaseClient := vdb.NewAPIClient(vdbConfig)
	client := &VDBClient{
		vDatabaseClient: vDatabaseClient,
		ProjectId: projectId,
	}
	return client, nil
}*/
