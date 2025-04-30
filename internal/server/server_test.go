package server_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/fabiant7t/hobot/internal/server"
)

func TestMarshalUnmarshal(t *testing.T) {
	jsonDoc := []byte(`[{"server":{"server_ip":"123.123.123.123","server_ipv6_net":"2a01:f48:111:4221::","server_number":321,"server_name":"server1","product":"DS 3000","dc":"NBG1-DC1","traffic":"5 TB","status":"ready","cancelled":false,"paid_until":"2010-09-02","ip":["123.123.123.123"],"subnet":[{"ip":"2a01:4f8:111:4221::","mask":"64"}]}},{"server":{"server_ip":"123.123.123.124","server_ipv6_net":"2a01:f48:111:4221::","server_number":421,"server_name":"server2","product":"X5","dc":"FSN1-DC10","traffic":"2 TB","status":"ready","cancelled":false,"paid_until":"2010-06-11","ip":["123.123.123.124"],"subnet":null}}]`)
	var res []server.ServerListItem
	if err := json.Unmarshal([]byte(jsonDoc), &res); err != nil {
		t.Error(err)
	}
	b, err := json.Marshal(res)
	if err != nil {
		t.Error(err)
	}
	if got, want := b, jsonDoc; !reflect.DeepEqual(got, want) {
		t.Error(cmp.Diff(got, want))
	}
}
