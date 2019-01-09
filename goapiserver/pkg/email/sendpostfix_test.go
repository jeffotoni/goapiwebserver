// Back-end in Go server
// @jeffotoni
// 2019-01-09

package email

import "testing"

func TestSendMail(t *testing.T) {
	type args struct {
		fromName  string
		fromEmail string
		toNames   []string
		toEmails  []string
		subject   string
		body      string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SendMail(tt.args.fromName, tt.args.fromEmail, tt.args.toNames, tt.args.toEmails, tt.args.subject, tt.args.body)
		})
	}
}
