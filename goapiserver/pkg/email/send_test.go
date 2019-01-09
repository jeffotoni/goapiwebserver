// Go Api server
// @jeffotoni
// 2019-01-04

package email

import "testing"

func TestSendEmail(t *testing.T) {
	type args struct {
		To      string
		From    string
		FromMsg string
		Nome    string
		Projeto string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SendEmail(tt.args.To, tt.args.From, tt.args.FromMsg, tt.args.Nome, tt.args.Projeto)
		})
	}
}
