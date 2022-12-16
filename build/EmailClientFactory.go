package build

import (
	clients1 "github.com/pip-services-infrastructure2/client-email-go/version1"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type EmailClientFactory struct {
	*cbuild.Factory
}

func NewEmailClientFactory() *EmailClientFactory {
	c := &EmailClientFactory{
		Factory: cbuild.NewFactory(),
	}

	cmdHttpClientDescriptor := cref.NewDescriptor("service-email", "client", "commandable-http", "*", "1.0")
	grpcClientDescriptor := cref.NewDescriptor("service-email", "client", "grpc", "*", "1.0")
	cmdGrpcClientDescriptor := cref.NewDescriptor("service-email", "client", "commandable-grpc", "*", "1.0")

	c.RegisterType(cmdHttpClientDescriptor, clients1.NewEmailCommandableHttpClientV1)
	c.RegisterType(grpcClientDescriptor, clients1.NewEmailGrpcClientV1)
	c.RegisterType(cmdGrpcClientDescriptor, clients1.NewEmailCommandableGrpcClientV1)

	return c
}
