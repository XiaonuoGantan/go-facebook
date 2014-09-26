package facebook

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	url "net/url"
)

const (
	ValidAPIVersions = "1.0,2.0,2.1"
	GraphHost        = "https://graph.facebook.com/"
)

type GraphAPI struct {
	AccessToken string
	Version     string
}

type GraphAPIError struct {
	Reason string
}

func (err GraphAPIError) Error() string {
	return err.Reason
}

func (api *GraphAPI) Request(path string, method string, args url.Values, postArgs url.Values) (map[string]interface{}, error) {
	endpoint := GraphHost + api.Version + path + "?access_token=" + api.AccessToken
	client := &http.Client{}
	if args != nil {
		q := args.Encode()
		endpoint = endpoint + "&" + q
	}
	if postArgs != nil {

	}
	req, err := http.NewRequest(method, endpoint, nil)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, GraphAPIError{err.Error()}
	}
	resp, err := client.Do(req)
	data, bodyErr := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		if bodyErr != nil {
			return nil, GraphAPIError{resp.Status}
		}
		return nil, GraphAPIError{string(data)}
	}
	var fbResp map[string]interface{}
	err = json.Unmarshal(data, &fbResp)
	if err != nil {
		return nil, GraphAPIError{err.Error()}
	}
	return fbResp, nil
}
