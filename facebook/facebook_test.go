package facebook

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

var (
	accessToken = flag.String("access_token", "", "Access Token")
)

const usageMsg = `
To obtain an access token you can access https://developers.facebook.com/tools/explorer/
and click the button "Get Access Token"
`

func Test_Facebook(t *testing.T) {
	flag.Parse()
	if *accessToken == "" {
		fmt.Fprint(os.Stderr, usageMsg)
		os.Exit(2)
	}
	fmt.Fprint(os.Stderr, *accessToken)
	api := GraphAPI{*accessToken, "v2.0"}
	resp, err := api.Request("/me", "GET", nil, nil)
	if err != nil {
		t.Error(err.Error())
	}
}
