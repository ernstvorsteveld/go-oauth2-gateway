package oauth2_gateway

type GatewayClient map[string]interface{}

type GatewayInputMapper interface {
	Map(input interface{}) (interface{}, error)
}

type ClientGateway interface {
	Create(client interface{}) (string, error)
	DeleteByClient(client interface{}) (bool, error)
	DeleteById(clientId string) (bool, error)
	Get(clientId string) (interface{}, error)
	List(page int, size int) ([]interface{}, error)
	Update(clientId string, client interface{}) (interface{}, error)
}
