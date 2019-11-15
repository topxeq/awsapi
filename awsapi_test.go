package awsapi

import (
	"net/url"
	"testing"
	"time"

	"github.com/topxeq/tk"
)

func Test001(t *testing.T) {
	timeStampT := tk.GetTimeStampMid(time.Now())
	uidT := "user1"

	urlStrT := "http://oa.example.com/portal/openapi"

	postDataT := url.Values{}

	postDataT.Set("timestamp", timeStampT)
	postDataT.Set("sig_method", "HmacMD5")
	postDataT.Set("cmd", "org.user.get")
	postDataT.Set("appId", "com.awspaas.user.apps.cms_viwe")
	postDataT.Set("access_key", "abcde12345")
	postDataT.Set("format", "json")
	postDataT.Set("uid", uidT)

	postDataT.Set("sig", Sign(postDataT, "12345abcde"))

	rs := tk.DownloadPageUTF8(urlStrT, postDataT, "", 15)

	tk.Pl("rs: %v", rs)

}
