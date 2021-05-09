package crud

import (
	"context"
        "strings"
	"fmt"
	_ "io"
	"log"
	"time"
	_ "github.com/spiffe/go-spiffe/v2/spiffeid"
	_ "github.com/spiffe/go-spiffe/v2/spiffetls"
	_ "github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
)

type RepositorySvidCRUD struct {
	svid *string
}

const (
	socketPath   = "unix:///var/run/spire/sockets/agent.sock"
	serverURL    = "http://0.0.0.0:8081"
)

// NewRepositoryPostsCRUD returns a new repository with DB connection
func NewRepositorySvidCRUD() *RepositorySvidCRUD {
	return &RepositorySvidCRUD{}
}

func (r *RepositorySvidCRUD) ValidateSpiffeID(sid string) (int, error) {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

        x509Source, err := workloadapi.NewX509Source(ctx, workloadapi.WithClientOptions(workloadapi.WithAddr(socketPath)))
	if err != nil {
		log.Fatalf("Unable to create X509Source %v", err)
	}
	defer x509Source.Close()

        x509, err := workloadapi.FetchX509SVID(ctx, workloadapi.WithAddr(socketPath))
        if err != nil {
                log.Fatalf("Unable to get X509Source %v", err)
                return 401, err
        }

        //fmt.Println(x509.ID)
        spiffeid := x509.ID.String()
        serSan := strings.Split(spiffeid, "/")
        fmt.Println("Source SAN: ", serSan[2])
        cliSan := strings.Split(sid, "/")
        fmt.Println("Passed SAN: ", cliSan[2])
        if serSan[2] == cliSan[2] {
                return 200, nil
        } else {
                return 401, err
        }

	//spiffeID := spiffeid.Must("spire-test.sds.local", "myworkload")
	//fmt.Println(spiffeID)

	//conn, err := spiffetls.DialWithMode(ctx, "tcp", serverAddress,
	//	spiffetls.MTLSClientWithSourceOptions(
	//		tlsconfig.AuthorizeID(spiffeID),
	//		workloadapi.WithClientOptions(workloadapi.WithAddr(socketPath)),
	//	))
	//if err != nil {
	//	log.Fatalf("Unable to create TLS connection: %v", err)
	//}
	//defer conn.Close()
 
        //tlsConfig := tlsconfig.MTLSClientConfig(source, source, tlsconfig.AuthorizeID(spiffeID))
        client, err := workloadapi.New(ctx, workloadapi.WithAddr(socketPath))
	if err != nil {
		log.Fatalf("Unable to create workload API client: %v", err)
	}
	defer client.Close()
        
        //req, err := client.Get(serverURL)
        //if err != nil {
	//	log.Fatalf("Error connecting to %q: %v", serverURL, err)
	//}

        //defer req.Body.Close()
	//body, err := ioutil.ReadAll(req.Body)
	//if err != nil {
	//	log.Fatalf("Unable to read body: %v", err)
	//      return 401, err
	//}

	//log.Printf("%s", body)

	return 200, nil
}
