// Go Api server
// @jeffotoni
// 2019-01-04

package email

var (
	URL_SEND_EMAIL = "https://sendmail.s3apis.com"
	END_POINT_SEND = URL_SEND_EMAIL + "/v1/api/send"
	X_KEY          = "MTIzNDU2Nzg5MA=="
	AUTORIZATION   = "Bearer MTIzNDU2Nzg5MGF1dG9yaXphdGlvbmJlYXJlZDIwMjA="
)

type DataEmail struct {
	To      string `json:to`
	From    string `json:from`
	FromMsg string `json:frommsg`
	Titulo  string `json:titulo`
	MsgHtml string `json:msghtml`
}
