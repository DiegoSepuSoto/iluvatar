package ainulindale

import sql "iluvatar/src/infrastructure/db/postgresql"

type ainulindalePostgresqlRepository struct {
	dbConnection sql.Connection
}

func NewAinulindalePostgresqlSQLRepository(dbConnection sql.Connection) *ainulindalePostgresqlRepository {
	return &ainulindalePostgresqlRepository{dbConnection: dbConnection}
}

