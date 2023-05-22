package integration

import (
	"testing"

	"github.com/fabiang/go-zabbix"
	"github.com/fabiang/go-zabbix/test"
)

func TestUserMacros(t *testing.T) {
	session := test.GetTestSession(t)

	params := zabbix.UserMacroGetParams{}

	macros, err := session.GetUserMacro(params)

	if err != nil {
		if _, ok := err.(*zabbix.NotFoundError); !ok {
			t.Fatalf("Error getting user macros: %v", err)
		}
	}

	if len(macros) == 0 {
		t.Skip("No usermacro found")
	}

	for i, macro := range macros {
		if macro.HostID == "" {
			t.Fatalf("User macro %d returned in response body has no Host ID", i)
		}
	}

	t.Logf("Validated %d user macros", len(macros))
}
