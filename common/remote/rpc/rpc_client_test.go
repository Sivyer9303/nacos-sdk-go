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

//func TestCreateNewConnection(t *testing.T) {
//	go startLocalServer()
//	//create ServerConfig
//	sc := []constant.ServerConfig{
//		*constant.NewServerConfig("127.0.0.1", 8848, constant.WithContextPath("/nacos")),
//	}
//
//	//create ClientConfig
//	cc := *constant.NewClientConfig(
//		constant.WithNamespaceId(""),
//		constant.WithTimeoutMs(5000),
//		constant.WithNotLoadCacheAtStart(true),
//		constant.WithLogDir("/tmp/nacos/log"),
//		constant.WithCacheDir("/tmp/nacos/cache"),
//		constant.WithLogLevel("debug"),
//	)
//	server, err := nacos_server.NewNacosServer(sc, cc, &http_agent.HttpAgent{}, 600, "localhost:8848")
//	if err != nil {
//		t.Fatal(" create nacos server fail")
//	}
//	client := NewGrpcClient("testClient", server)
//	serverInfo := ServerInfo{
//		serverIp:   "localhost",
//		serverPort: 8848,
//	}
//	connection, err := client.createNewConnection(serverInfo)
//	if err != nil {
//		t.Fatal("create connection fail")
//	}
//	assert.NotNil(t, connection)
//}
//
//func TestGrpcClient_connectToServer(t *testing.T) {
//	go startLocalServer()
//	//create ServerConfig
//	sc := []constant.ServerConfig{
//		*constant.NewServerConfig("127.0.0.1", 8848, constant.WithContextPath("/nacos")),
//	}
//
//	//create ClientConfig
//	cc := *constant.NewClientConfig(
//		constant.WithNamespaceId(""),
//		constant.WithTimeoutMs(5000),
//		constant.WithNotLoadCacheAtStart(true),
//		constant.WithLogDir("/tmp/nacos/log"),
//		constant.WithCacheDir("/tmp/nacos/cache"),
//		constant.WithLogLevel("debug"),
//	)
//	server, err := nacos_server.NewNacosServer(sc, cc, &http_agent.HttpAgent{}, 600, "localhost:8848")
//	if err != nil {
//		t.Fatal("start fail")
//	}
//	client := NewGrpcClient("test", server)
//	response, err := client.Request(rpc_request.NewHealthCheckRequest(), 600)
//	if err != nil {
//		t.Fatal("request fail")
//	}
//	fmt.Println(response)
//}
//
//type implRequestServer struct {
//	*pb.UnimplementedRequestServer
//}
//
//// response the same payload
//func (rs *implRequestServer) Request(ct context.Context, pl *pb.Payload) (*pb.Payload, error) {
//	select {
//	case <-ct.Done():
//		return pl, errors.New("context is done")
//	default:
//		return pl, nil
//	}
//}
//
//// start up local grpc server
//func startLocalServer() {
//	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8848))
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//
//	s := grpc.NewServer()
//	pb.RegisterRequestServer(s, &implRequestServer{})
//	log.Printf("server listening at %v", lis.Addr())
//	if err := s.Serve(lis); err != nil {
//		log.Fatalf("failed to serve: %v", err)
//	}
//}
