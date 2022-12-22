package customer

import (
	"context"
	"game-store-final-project/project/pkg/mysql_connection"
)

var (
	mysqlConn         = mysql_connection.InitMysqlDB()
	ctx               = context.Background()
	repoMysqlCustomer = NewCustomerRepositoryMysqlInteractor(mysqlConn)
)
