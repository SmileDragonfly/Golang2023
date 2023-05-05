package main

import (
	"testing"
)

func TestReadPrivateKey(t *testing.T) {
	type args struct {
		sPrivateKey string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"TestReadPrivateKey",
			args{"MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCVEi/qLwGQWsJk+BH9vrcavcrwaV/m8n4adWt0MPhlWvEqv7bcLoWYu2o+3whAiLY4ZybN6acfCqttAs2+rHcF88unh6w1PLqaCB5JErdeRmgXnYlCu6jIZMYhhvpl5yBfCTHDB1FGF47Bd/EfztIu21NScXddsN0tiaB1uT3/RPon68wokvEzOMMQl89Ik79XDIJPhBnsNU9GX838NaD3Chzdn7eFGUqUdOjbyoIVkHrXj4qfx1935YLX7d/g9auFr7k3ClE7iXJjRCJHmkF8c0seyD2N4TzYhPyAj7Ku76Uls83Tu7KrmQEkPUL2tRGyRC6wimnOlBPA1Ad0BCtPAgMBAAECggEADh6gYjacl8847XZfweCQGFzUvYvFlSrvzdLEYEeJJ5SDFlD9YYKYjdxmlljqv64TUGlY0BUrCrIohZVH5qEQTwSGDDK6GXrMk+1j7Pj0XF4f2ujFiQgcVwrQh3lw+zj2pHnK+FWXmaN+lo2lTaV53A1TKZJsw3eOo5MPG0d1+1kotGrZeOFELRnjLkGFw4SntahhMHByXKAl8pFGUyRJqHZBXw4erXbRLMDg7TAYn2VJdHrhSeEzFGmKA7QHuTRq52SZg71xDLiHmFVSeP7tVrj60jN7pjJTjJR6C2fFxphEpKaQ02zRzTjmqsAMbP77GmlWtyP/+crw70P4q/yHXQKBgQDFO0xn32CwlsTb+QIZnhfYO46+lfJP6MFuK/WIu0p2tee8xGDQYkWZ4xETkJb604ArhQ9WATDKfZkloCtk8pcRqVp6Q0xFjL+pfnZ4W9CUOwxpf1kPn8u6Dcay5CYLtbibrbaA5XR2HjsdxE+lxl29nfwabYEqCCmmPEPF2YaPZQKBgQDBfTrfa59KLzRHGAz55YXk6KNcjHzzRYT4b7Y4Hzh1OCB27tnVZjGBp8IUrXSnodPmFhcia/dQJJCZdciQ8Z8oj3XuyY4iB0QsXxBmFomjH/FYI+1hnhQehQvAqAeIpYlNRXYYPxqAfn3AQJOZebDrT/V4JaS4g9YjcQT4IqiGowKBgQCqZKeG2cIj7a1XSZJZ5W4+Pn39A3hbNv/dmZa/sOcNFeyF9baacTwmTbiUCYeWXSDO+F6ec9reJZIoom67AKYo+QGUvQ1ozMdMvFfHdbMGTNlVT1L3H5uXOo2eQWLpHO7HeFVCmHl8DnQOLGqPEogr6BBEGLTNRk4NMuVuSZZpzQKBgET23MgLdRAc+RYp9V4QuAOaA7gV/uc6rSVbs+gXAKmPIsshYRUVwqmC4MM7++tP29YTo5VKRDEVh1CbUayP4nmzgIZm4rkwO9VQ4OhyOgaheQVAcPitPmCObVzyxxSmY+Td0DTeMRUBgNLIcZNvc2a77jMvv6FgpC+ntey3dbffAoGAPvAnfdpFlUlRDdbsgn6HeiXcErCETeUPnCuiMjrb0zI1mnaJEPTJukQEFoHNmtFNgK4t1l0UEApYJ1PPPcyKVmvJSVwaXVxnLHvIj0vpNjTs9gdG37EZHFKXQjjhn+TD2iOXv1kqnyLJQjFBw+xgDtzp8qITXmJt3qKdpA7KQBw="},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadPrivateKey(tt.args.sPrivateKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadPrivateKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}

func TestGetRequestDateTime(t *testing.T) {
	time := GetRequestDateTime()
	t.Log(time)
}
