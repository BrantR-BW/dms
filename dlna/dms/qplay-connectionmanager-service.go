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

//TODO: Update the handler to handle the events listed in the QPlay services doc
func (me *QPlayConnectionManagerService) Handle(action string, argsXML []byte, r *http.Request) (map[string]string, error) {
	//host := r.Host
	//userAgent := r.UserAgent()

    fmt.Println("Handle called! for %s\n", action)
    return map[string]string{
        "ThingA": "ThingB",
    }, nil

    /*
	switch action {
	case "GetSystemUpdateID":
		return map[string]string{
			"Id": me.updateIDString(),
		}, nil
	case "GetSortCapabilities":
		return map[string]string{
			"SortCaps": "dc:title",
		}, nil
	case "Browse":
		var browse browse
		if err := xml.Unmarshal([]byte(argsXML), &browse); err != nil {
			return nil, err
		}
		obj, err := me.objectFromID(browse.ObjectID)
		if err != nil {
			return nil, upnp.Errorf(upnpav.NoSuchObjectErrorCode, err.Error())
		}
		switch browse.BrowseFlag {
		case "BrowseDirectChildren":
			objs, err := me.readContainer(obj, host, userAgent)
			if err != nil {
				return nil, upnp.Errorf(upnpav.NoSuchObjectErrorCode, err.Error())
			}
			totalMatches := len(objs)
			objs = objs[func() (low int) {
				low = browse.StartingIndex
				if low > len(objs) {
					low = len(objs)
				}
				return
			}():]
			if browse.RequestedCount != 0 && int(browse.RequestedCount) < len(objs) {
				objs = objs[:browse.RequestedCount]
			}
			result, err := xml.Marshal(objs)
			if err != nil {
				return nil, err
			}
			return map[string]string{
				"TotalMatches":   fmt.Sprint(totalMatches),
				"NumberReturned": fmt.Sprint(len(objs)),
				"Result":         didl_lite(string(result)),
				"UpdateID":       me.updateIDString(),
			}, nil
		case "BrowseMetadata":
			fileInfo, err := os.Stat(obj.FilePath())
			if err != nil {
				if os.IsNotExist(err) {
					return nil, &upnp.Error{
						Code: upnpav.NoSuchObjectErrorCode,
						Desc: err.Error(),
					}
				}
				return nil, err
			}
			upnp, err := me.cdsObjectToUpnpavObject(obj, fileInfo, host, userAgent)
			if err != nil {
				return nil, err
			}
			buf, err := xml.Marshal(upnp)
			if err != nil {
				return nil, err
			}
			return map[string]string{
				"TotalMatches":   "1",
				"NumberReturned": "1",
				"Result":         didl_lite(func() string { return string(buf) }()),
				"UpdateID":       me.updateIDString(),
			}, nil
		default:
			return nil, upnp.Errorf(upnp.ArgumentValueInvalidErrorCode, "unhandled browse flag: %v", browse.BrowseFlag)
		}
	case "GetSearchCapabilities":
		return map[string]string{
			"SearchCaps": "",
		}, nil
	default:
		return nil, upnp.InvalidActionError
	}
    */
}

