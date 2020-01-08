package dms

import (
	"net/http"
    "fmt"

	"github.com/anacrolix/dms/upnp"
)

type QPlayRenderingControlService struct {
	*Server
	upnp.Eventing
}

func (me *QPlayRenderingControlService) Handle(action string, argsXML []byte, r *http.Request) (map[string]string, error) {

    fmt.Println("QPlayRenderingControlService::Handle()! for %s\n", action)

    switch action {
    // Params to this are: "InstanceID", "Channel", returns "Mute" bool
	case "GetMute":
		return map[string]string{
			"Mute": "true",
		}, nil
    // Params to this are: "InstanceID", "Channel", "Mute"
    case "SetMute":
		return map[string]string{
			"": "",
		}, nil
    // Params to this are: "InstanceID", "Channel", returns "CurrentVolume"
    case "GetVolume":
		return map[string]string{
			"CurrentVolume": "50",
		}, nil
    // Params to this are: "InstanceID", "Channel", "DesiredVolume"
    case "SetVolume":
		return map[string]string{
			"": "",
		}, nil
	default:
		return nil, upnp.InvalidActionError
	}
}

