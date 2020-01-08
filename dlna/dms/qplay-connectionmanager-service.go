package dms

import (
	"net/http"
    "fmt"

	"github.com/anacrolix/dms/upnp"
)

type QPlayConnectionManagerService struct {
	*Server
	upnp.Eventing
}

func (me *QPlayConnectionManagerService) Handle(action string, argsXML []byte, r *http.Request) (map[string]string, error) {

    fmt.Println("QPlayConnectionManagerService::Handle() %s\n", action)

    //TODO: Implement the actual behavior the QPlay service wants
	switch action {
	case "GetProtocolInfo":
		return map[string]string{
			"Source": "A Source",
            "Sink": "A Sink",
		}, nil
    default:
		return nil, upnp.InvalidActionError
	}
}

