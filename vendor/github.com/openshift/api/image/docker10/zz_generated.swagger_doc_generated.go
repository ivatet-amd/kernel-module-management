package docker10

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_DockerConfig = map[string]string{
	"": "DockerConfig is the list of configuration options used when creating a container.",
}

func (DockerConfig) SwaggerDoc() map[string]string {
	return map_DockerConfig
}

var map_DockerImage = map[string]string{
	"": "DockerImage is the type representing a container image and its various properties when retrieved from the Docker client API.\n\nCompatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
}

func (DockerImage) SwaggerDoc() map[string]string {
	return map_DockerImage
}

// AUTO-GENERATED FUNCTIONS END HERE