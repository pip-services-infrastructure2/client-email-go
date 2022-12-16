package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-infrastructure2/client-email-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type emailHttpClientV1Test struct {
	client  *version1.EmailCommandableHttpClientV1
	fixture *EmailClientFixtureV1
}

func newEmailHttpClientV1Test() *emailHttpClientV1Test {
	return &emailHttpClientV1Test{}
}

func (c *emailHttpClientV1Test) setup(t *testing.T) {
	var HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		HTTP_HOST = "localhost"
	}
	var HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", HTTP_HOST,
		"connection.port", HTTP_PORT,
	)

	c.client = version1.NewEmailCommandableHttpClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewEmailClientFixtureV1(c.client)
}

func (c *emailHttpClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestHttpClientSendEmail(t *testing.T) {
	c := newEmailHttpClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestSendEmailToAddress(t)
	c.fixture.TestSendEmailToRecipient(t)
}
