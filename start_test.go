package agollo

import (
	"testing"
	"time"
	// "github.com/CastleSky/agollo/test"
)

func TestStart(t *testing.T) {
	go runMockConfigServer(onlyNormalConfigResponse)
	go runMockNotifyServer(onlyNormalResponse)
	defer closeMockConfigServer()

	appconfig := &agollo.AppConfig{
		AppId:            "123",
		Cluster:          "default",
		NamespaceName:    "application",
		NextTryConnTime:  3,
		Ip:               "http://service-apollo-config-server-dev.sre:8080",
		BackupConfigPath: "backConfig",
	}

	Start(appconfig)

	value := getValue("key1")
	// test.Equal(t, "value1", value)
}

func TestErrorStart(t *testing.T) {
	server := runErrorResponse()
	newAppConfig := getTestAppConfig()
	newAppConfig.Ip = server.URL

	time.Sleep(1 * time.Second)

	Start()

	value := getValue("key1")
	// test.Equal(t, "value1", value)

	value2 := getValue("key2")
	// test.Equal(t, "value2", value2)

}
