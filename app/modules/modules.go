package modules

import (
	"robbo-assets/package/db_client"
)

type HandlerModule struct {
	//UsersHandler       usershtpp.Handler
}

func SetupHandler(postgresClient db_client.PostgresClient) HandlerModule {
	return HandlerModule{
		//AuthGateway:        authgateway.SetupAuthGateway(postgresClient),
	}
}
