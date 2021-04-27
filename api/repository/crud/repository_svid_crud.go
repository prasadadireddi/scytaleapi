package crud

import (
    _ "github.com/prasadadireddi/scytaleapi/api/models"
    "github.com/prasadadireddi/scytaleapi/api/utils/channels"

    "bufio"
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
)

type RepositorySvidCRUD struct {
	svid *string
}

const (
	socketPath    = "unix:///var/run/spire/sockets/agent.sock"
	serverAddress = "spire-test.sds.local:8081"
)

// NewRepositoryPostsCRUD returns a new repository with DB connection
func NewRepositorySvidCRUD() *RepositorySvidCRUD {
	return &RepositorySvidCRUD{}
}


func (r *RepositorySvidCRUD) ValidateSpiffeID(sid string) (int, error) {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	spiffeID := spiffeid.Must("spire-test.sds.local", "myworkload")
	fmt.Println(spiffeID)

	conn, err := spiffetls.DialWithMode(ctx, "tcp", serverAddress,
		spiffetls.MTLSClientWithSourceOptions(
			tlsconfig.AuthorizeID(spiffeID),
			workloadapi.WithClientOptions(workloadapi.WithAddr(socketPath)),
		))
	if err != nil {
		log.Fatalf("Unable to create TLS connection: %v", err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "Hello server\n")

	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		status, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatalf("Unable to read server response: %v", err)
			ch <- false
			return
		}
		log.Printf("Server says: %q", status)
		ch <- true
	}(done)
	
	if channels.OK(done) {
		return 200, nil
	}
	return 401, err
}
