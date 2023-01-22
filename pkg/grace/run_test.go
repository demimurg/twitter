package grace

import (
	"context"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"google.golang.org/grpc/reflection"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestRun(t *testing.T) {
	t.Parallel()
	grpcSrv := grpc.NewServer()
	reflection.Register(grpcSrv)
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

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			waitForExecute := make(chan struct{})
			grpcPort, grpcuiPort, httpPort, prometheusPort := getPort(), getPort(), getPort(), getPort()
			go func() {
				Run(
					GRPC(grpcSrv, grpcPort),
					GRPCUI(grpcuiPort, "localhost"+grpcPort),
					HTTP(httpSrv, httpPort),
					Prometheus(prometheusPort),
				)
				close(waitForExecute)
			}()

			<-time.After(100 * time.Millisecond)
			err := tc.terminate()
			require.NoError(t, err)

			select {
			case <-waitForExecute:
				for _, port := range []string{grpcPort, grpcuiPort, httpPort, prometheusPort} {
					// if server isn't stop, it will "address already in use" error
					_, err := net.Listen("tcp", port)
					require.NoError(t, err)
				}
			case <-time.After(2 * time.Second):
				t.Error("processes doesn't shutdown in two second")
			}
		})
	}
}

// getPort guaranteed to return open port in ":8080" format
func getPort() string {
	// if "0" port automatically chosen
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	// close listener because this port should be free to use
	defer lis.Close()

	addr := lis.Addr().String()
	i := strings.LastIndex(addr, ":")
	return addr[i:]
}

func TestGetPort(t *testing.T) {
	for i := 0; i < 10; i++ {
		port := getPort()
		_, err := net.Listen("tcp", port)
		require.NoError(t, err)

		t.Log("Available port", port)
	}
}
