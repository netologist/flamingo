package vodka

type Services struct {
	DbClientFactory   *DbClientFactory
	HttpClientFactory *HttpClientFactory
	LoggerFactory     *LoggerFactory
}
