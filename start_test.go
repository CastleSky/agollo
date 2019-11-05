package agollo

import (
	"testing"
	"time"

	"github.com/livbarn/agollo/test"
	// "github.com/livbarn/agollo/test"
)

func TestStart(t *testing.T) {
	go runMockConfigServer(onlyNormalConfigResponse)
	go runMockNotifyServer(onlyNormalResponse)
	defer closeMockConfigServer()

	appconfig := &AppConfig{
		AppId:            "123",
		Cluster:          "default",
		NamespaceName:    "application",
		NextTryConnTime:  3,
		Ip:               "http://service-apollo-config-server-dev.sre:8080",
		BackupConfigPath: "backConfig",
	}

	Start(appconfig)

	value := getValue("key1")
	test.Equal(t, "value1", value)
}

func TestErrorStart(t *testing.T) {
	server := runErrorResponse()
	newAppConfig := getTestAppConfig()
	newAppConfig.Ip = server.URL

	time.Sleep(1 * time.Second)

	appconfig := &AppConfig{
		AppId:            "test",
		Cluster:          "dev",
		NamespaceName:    "application",
		NextTryConnTime:  3,
		Ip:               "127.0.0.1:8888",
		BackupConfigPath: "backConfig",
	}

	Start(appconfig)

	value := getValue("key1")
	test.Equal(t, "value1", value)

	value2 := getValue("key2")
	test.Equal(t, "value2", value2)

}
