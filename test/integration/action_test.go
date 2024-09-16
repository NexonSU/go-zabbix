package integration

import (
	"testing"

	"github.com/fabiang/go-zabbix"
	"github.com/fabiang/go-zabbix/test"
)

func TestActionsIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	session := test.GetTestSession(t)

	params := zabbix.ActionGetParams{}

	actions, err := session.GetActions(params)

	if err != nil {
		if _, ok := err.(*zabbix.NotFoundError); !ok {
			t.Fatalf("Error getting actions: %v", err)
		}
	}

	if len(actions) == 0 {
		t.Skip("No actions found")
	}

	for i, action := range actions {
		if action.ActionID == "" {
			t.Fatalf("Action %d has no Action ID", i)
		}

		if action.Name == "" {
			t.Fatalf("Action %d has no name", i)
		}
	}

	t.Logf("Validated %d Actions", len(actions))
}
