package dms

import (
	"net/http"
    "fmt"

	"github.com/anacrolix/dms/upnp"
)

type QPlayAVTransportService struct {
	*Server
	upnp.Eventing
}

func (me *QPlayAVTransportService) Handle(action string, argsXML []byte, r *http.Request) (map[string]string, error) {

    fmt.Println("QPlayAVTransportService::Handle() called %s\n", action)

	switch action {
    // Params to this are: "InstanceID", "CurrentURI", "CurrentURIMetaData" 
	case "SetAVTransportURI":
		return map[string]string{
			"": "",
		}, nil
    // Params to this are: "InstanceID", returns (CurrentURI string, CurrentURIMetaData string)
    case "GetMediaInfo":
		return map[string]string{
			"CurrentURI": "something",
			"CurrentURIMetadata": "something",
		}, nil
    // Params to this are: "InstanceID", returns "PositionInfoResponse"
    case "GetPositionInfo":
		return map[string]string{
			"PositionInfoResponse": "",
		}, nil
    // Params to this are: "InstanceID"
    case "Stop":
		return map[string]string{
			"": "",
		}, nil
    // Params to this are: "InstanceID", "Speed"
    case "Play":
		return map[string]string{
			"": "",
		}, nil
    // Params to this are: "InstanceID"
    case "Pause":
		return map[string]string{
			"": "",
		}, nil
    // Params to this are: "InstanceID", "Speed", "Target"
    case "Seek":
		return map[string]string{
			"": "",
		}, nil
    // Params to this are: "InstanceID", "NewPlayMode"
    case "SetPlayMode":
		return map[string]string{
			"": "",
		}, nil
	default:
		return nil, upnp.InvalidActionError
	}
}

