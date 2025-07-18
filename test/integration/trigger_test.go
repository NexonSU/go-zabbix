package integration

import (
	"testing"

	"github.com/NexonSU/go-zabbix"
	"github.com/NexonSU/go-zabbix/test"
)

func TestTriggersIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	session := test.GetTestSession(t)

	params := zabbix.TriggerGetParams{}

	triggers, err := session.GetTriggers(params)
	if err != nil {
		if _, ok := err.(*zabbix.NotFoundError); !ok {
			t.Fatalf("Error getting triggers: %v", err)
		}
	}

	if len(triggers) == 0 {
		t.Skip("No triggers found")
	}

	for i, trigger := range triggers {
		if trigger.TriggerID == "" {
			t.Fatalf("Trigger %d has no trigger ID", i)
		}
	}

	t.Logf("Validated %d triggers", len(triggers))
}
