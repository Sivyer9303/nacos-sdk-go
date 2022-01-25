package rpc

import (
	"errors"
	"os"
	"testing"

	"github.com/nacos-group/nacos-sdk-go/v2/common/nacos_server"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {

}
func TestRpcClient_ConnectToServer(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	client := NewMockIRpcClient(ctl)
	client.EXPECT().connectToServer(gomock.Eq(ServerInfo{})).Return(nil, errors.New("wrong ip and host"))
	server, err := client.connectToServer(ServerInfo{})
	assert.Nil(t, server)
	assert.NotNil(t, err)
}

func TestRpcClient_GetRpcClient(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	client := NewMockIRpcClient(ctl)
	labels := map[string]string{
		"label": "label",
	}
	clientName := "config-taskId-uid"
	createClient, _ := CreateClient(clientName, GRPC, labels, nil)
	client.EXPECT().GetRpcClient().Return(createClient.GetRpcClient())
	rpcClient := client.GetRpcClient()
	allClient := getAllClient()
	iRpcClient := getClient(clientName)
	assert.NotNil(t, rpcClient)
	assert.NotNil(t, allClient)
	assert.NotNil(t, iRpcClient)
}

func TestRpcClient_GetConnectionType(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	client := NewMockIRpcClient(ctl)
	client.EXPECT().getConnectionType().Return(GRPC)
	connectionType := client.getConnectionType()
	assert.Equal(t, GRPC, connectionType)
}

func TestRpcClient_RpcPortOffset(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	client := NewMockIRpcClient(ctl)
	client.EXPECT().rpcPortOffset().Return(uint64(1000))
	offset := client.rpcPortOffset()
	assert.Equal(t, offset, uint64(1000))
}

func TestNewGrpcClient(t *testing.T) {
	config := &nacos_server.NacosServer{}
	grpcClient := NewGrpcClient("testNewGrpcClient", config)
	assert.NotNil(t, grpcClient)
	assert.NotNil(t, grpcClient.RpcClient)
}

func TestGetMaxCallRecvMsgSize(t *testing.T) {
	t.Run("defaultMsgSizeTest", func(t *testing.T) {
		defaultMsgSize := 10485760 // 10*1024*1024
		size := getMaxCallRecvMsgSize()
		assert.Equal(t, size, defaultMsgSize)
	})

	t.Run("messageSizeTest", func(t *testing.T) {
		msgSize := "20971520" // 20*1024*1024
		err := os.Setenv("nacos.remote.client.grpc.maxinbound.message.size", msgSize)
		if err != nil {
			return
		}
		size := getMaxCallRecvMsgSize()
		assert.Equal(t, size, 20971520)
	})
}
