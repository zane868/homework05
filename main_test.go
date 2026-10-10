package main

import (
	"errors"
	"strings"
	"testing"
)

func TestParseEther(t *testing.T) {
	tests := []struct {
		input string
		wei   string
	}{
		{"1", "1000000000000000000"},
		{"0.1", "100000000000000000"},
		{"0.01", "10000000000000000"},
		{"0.000000000000000001", "1"},
	}
	for _, test := range tests {
		got, err := parseEther(test.input)
		if err != nil || got.String() != test.wei {
			t.Errorf("parseEther(%q) = %v, %v; want %s", test.input, got, err, test.wei)
		}
		if err == nil && formatWei(got) != test.input {
			t.Errorf("formatWei(parseEther(%q)) = %q", test.input, formatWei(got))
		}
	}
	for _, input := range []string{"", "0", "-1", "0.1.2", "1.", ".1", "0.0000000000000000001", "1e3"} {
		if got, err := parseEther(input); err == nil {
			t.Errorf("parseEther(%q) = %v; want error", input, got)
		}
	}
}

func TestSafeError(t *testing.T) {
	t.Setenv("ETHERSCAN_API_KEY", "example-secret")
	message := safeError(errors.New("request apikey=example-secret failed"))
	if strings.Contains(message, "example-secret") {
		t.Fatalf("secret appeared in error: %s", message)
	}
}

func TestCommandDispatch(t *testing.T) {
	tests := []struct {
		args []string
		want string
	}{
		{[]string{"block"}, "请提供 -number"},
		{[]string{"tx", "get"}, "请提供有效的"},
		{[]string{"tx", "send"}, "请提供有效的 -to"},
		{[]string{"account", "txs"}, "请提供有效的 -address"},
		{[]string{"donation", "get"}, "请提供有效的 -donor"},
		{[]string{"donation", "send"}, "无效的 -amount"},
		{[]string{"donation", "events"}, "请提供有效的 -tx"},
		{[]string{"unknown"}, "未知命令"},
	}
	for _, test := range tests {
		err := run(test.args)
		if err == nil || !strings.Contains(err.Error(), test.want) {
			t.Errorf("run(%q) = %v; want error containing %q", test.args, err, test.want)
		}
	}
}
