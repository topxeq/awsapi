package awsapi

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"sort"
	"strings"
)

// Sign Signing the url.Values
func Sign(valuesA url.Values, keyA string) string {
	if valuesA == nil {
		return ""
	}

	var bufT strings.Builder

	keysT := make([]string, 0, len(valuesA))

	for k := range valuesA {
		keysT = append(keysT, k)
	}

	sort.Strings(keysT)

	for _, k := range keysT {
		vi := valuesA[k]

		if len(vi) < 1 {
			continue
		}

		bufT.WriteString(k)
		bufT.WriteString(vi[0])
	}

	strT := keyA + bufT.String()

	hmacT := hmac.New(md5.New, []byte(keyA))

	hmacT.Write([]byte(strT))

	return strings.ToUpper(hex.EncodeToString(hmacT.Sum([]byte(""))))
}
