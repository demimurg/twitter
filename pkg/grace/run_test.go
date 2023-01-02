package grace

import (
	"context"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestRun(t *testing.T) {
	grpcSrv := grpc.NewServer()
	httpSrv := &http.Server{Handler: http.NewServeMux()}

	var testCases = []struct {
		name      string
		terminate func() error // with execute process id
	}{
		{
			name: "http server out",
			terminate: func() error {
				return httpSrv.Shutdown(context.Background())
			},
		},
		{
			name: "grpc server out",
			terminate: func() error {
				grpcSrv.Stop()
				return nil
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			waitForExecute := make(chan struct{})
			grpcPort, httpPort := getPort(), getPort()
			go func() {
				Run(
					GRPC(grpcSrv, grpcPort),
					HTTP(httpSrv, httpPort),
				)
				close(waitForExecute)
			}()

			<-time.After(10 * time.Millisecond)
			err := test.terminate()
			require.NoError(t, err)

			select {
			case <-waitForExecute:
				for _, port := range []string{grpcPort, httpPort} {
					// if server isn't stop, it will "address already in use" error
					_, err := net.Listen("tcp", port)
					require.NoError(t, err)
				}
			case <-time.After(2 * time.Second):
				t.Error("routines doesn't shutdown in two second")
			}
		})
	}
}

// getPort guaranteed to return open port in ":8080" format
func getPort() string {
	// if "0" port automatically chosen
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}

	addr := lis.Addr().String()
	index := strings.Index(addr, ":")
	return addr[index:]
}

func TestGetPort(t *testing.T) {
	for i := 0; i < 10; i++ {
		port := getPort()
		t.Log("Available port ", port)
		lis, err := net.Listen("tcp", port)
		require.NoError(t, err)
		require.NoError(t, lis.Close())
	}
}
