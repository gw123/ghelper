package hashids

import (
	"testing"
)

func TestHashIdsHelper_DecodeString(t *testing.T) {

	type args struct {
		input string
	}
	tests := []struct {
		name    string
		slat    string
		args    args
		wantErr bool
	}{
		{
			name: "000",
			slat: "123456",
			args: args{
				input: "hello",
			},
		},
		{
			name: "001",
			slat: "990209292",
			args: args{
				input: "hello world",
			},
		},
		{
			name: "002",
			slat: "xsjioklxiiwkugxpwqxvm",
			args: args{
				input: "this is input string.",
			},
		},
		{
			name: "003",
			slat: "xsjioklxiiwkugxpwqxvm%?.!",
			args: args{
				input: "this is input string.",
			},
		},
		{
			name: "004",
			slat: "xsjioklxiiwkugxpwqxvm%?.!#()",
			args: args{
				input: "2",
			},
		},
		{
			name: "004",
			slat: "xsjioklxiiwkugxpwqxvm%?.!#()",
			args: args{
				//输入不能为空
				input: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h, err := NewHashIdsHelper(tt.slat)
			if err != nil {
				t.Error(err)
			}

			temp, err := h.EncodeString(tt.args.input)
			if err != nil && !tt.wantErr {
				t.Error(err)
			}
			t.Logf("encodeString %s => %s", tt.args.input, temp)
			deInput, err := h.DecodeString(temp)
			if err != nil {
				t.Error(err)
			}

			if tt.args.input != deInput {
				t.Errorf("test:%s ,input: %s , deInput:%s", tt.name, tt.args.input, deInput)
			}
		})
	}
}

func TestEncodeString(t *testing.T) {
	type args struct {
		input string
		salt  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "000",
			args: args{
				input: "hello world",
				salt:  "123456",
			},
			want: "2jozKvwcGU5cjfmcYs4c5sLcwXvQBO",
		},
		//{
		//	name: "001",
		//	args: args{
		//		input: "this is input string.",
		//		salt:  "xsjioklxiiwkugxpwqxvm",
		//	},
		//	want: "k8ULtacyslcxHMUVFzSVTwcbHAUrFxSpTzcrHrcmTeUVTxU9IKUytxS3TJU8FwUxtWUPS8cXHXc3T2cRUQSB",
		//},
		//{
		//	name: "002",
		//	args: args{
		//		input: "this 990x9l2x1l00210.",
		//		salt:  "xsjioklxiiwkugxpwqxvm",
		//	},
		//},
		//{
		//	name: "003",
		//	args: args{
		//		input: "this ^&*()(_+=!@#><>??/\\|nput string.",
		//		salt:  "xsjioklxiiwkugxpwqxvm",
		//	},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncodeString(tt.args.input, tt.args.salt)
			t.Logf("EncodeString %s ,%s ,%s", tt.args.input, tt.args.salt, got)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != "" && got != tt.want {
				t.Errorf("EncodeString() got = %v, want %v", got, tt.want)
			}

			deStr, err := DecodeString(got, tt.args.salt)
			if deStr != tt.args.input || err != nil {
				t.Errorf("DecodeString() got = %v, want: %v ,err = %v", deStr, tt.args.input, err)
			}
		})
	}
}
