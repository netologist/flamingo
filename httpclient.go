package vodka

type HttpClientFactory struct {
	loggerFactory *LoggerFactory
}

func NewHttpClientFactory(loggerFactory *LoggerFactory) *HttpClientFactory {
	return &HttpClientFactory{loggerFactory}
}
