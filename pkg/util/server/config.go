package server

var defConfig = Config{
	Address: ":8888",
}

type Config struct {
	Address            string
	SwaggerUIPath      string
	SwaggerFilePath    string
	OpenAPIPath        string
	CORSAllowedHeaders []string
	CORSAllowedMethods []string
	WriteMIMEs         []string
}

//func fillConfigWithDefaultValues(config *Config) {
//	if config.SwaggerUIPath == "" {
//		config.SwaggerUIPath = `/swagger`
//	}
//	if config.SwaggerFilePath == "" {
//		config.SwaggerFilePath = `manifests/swagger`
//	}
//	if config.OpenAPIPath == "" {
//		config.OpenAPIPath = `/swagger.json`
//	}
//	if len(config.CORSAllowedHeaders) == 0 {
//		config.CORSAllowedHeaders = []string{"Content-Type", "Accept"}
//	}
//	if len(config.CORSAllowedMethods) == 0 {
//		config.CORSAllowedMethods = []string{"GET", "PUT", "POST", "DELETE"}
//	}
//	if len(config.WriteMIMEs) == 0 {
//		config.WriteMIMEs = []string{"application/json"}
//	}
//}
