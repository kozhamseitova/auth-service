package mongodb

type Option func(*MongoDB)

func WithHost(host string) Option {
	return func(mongoDB *MongoDB) {
		mongoDB.host = host
	}
}

func WithPort(port string) Option {
	return func(mongoDB *MongoDB) {
		mongoDB.port = port
	}
}

func WithDBName(dbName string) Option {
	return func(mongoDB *MongoDB) {
		mongoDB.dbName = dbName
	}
}