package vodka

type DbClientFactory struct {
	loggerFactory *LoggerFactory
}

func NewDbClientFactory(loggerFactory *LoggerFactory) *DbClientFactory {
	return &DbClientFactory{loggerFactory}
}
