package servercmd

import "testing"

func TestMatchServer(t *testing.T) {
	for _, tc := range []struct {
		testName   string
		srvName    string
		srvIPs     []string
		queryName  string
		queryIP    string
		ignoreCase bool
		want       bool
	}{
		{"Search by name", "marvin", []string{"1.2.3.4"}, "marvin", "", false, true},
		{"Search by IP", "marvin", []string{"1.2.3.4"}, "", "1.2.3.4", false, true},
		{"Search by neither name nor IP", "marvin", []string{"1.2.3.4"}, "", "", false, false},
		{"Search by case-insensitive name", "macOS", []string{"1.2.3.4"}, "macos", "", true, true},
		{"Search by IP is always case-insensitive", "marvin", []string{"afaf:face:cafe:fade:deaf:beef:babe:f00d"}, "", "AFAF:FACE:CAFE:FADE:DEAF:BEEF:BABE:F00D", false, true},
		{"Search by IP and name, both match", "marvin", []string{"1.2.3.4"}, "marvin", "1.2.3.4", false, true},
		{"Search by IP and name, only ip matches", "marvin", []string{"1.2.3.4"}, "deep thought", "1.2.3.4", false, false},
		{"Search by IP and name, only name matches", "marvin", []string{"1.2.3.4"}, "marvin", "44.33.22.11", false, false},
	} {
		if got, want := matchServer(tc.srvName, tc.srvIPs, tc.queryName, tc.queryIP, tc.ignoreCase), tc.want; got != want {
			t.Errorf("%s: Got %t, want %t", tc.testName, got, want)
		}
	}
}
