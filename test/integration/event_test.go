package integration

import (
	"testing"

	"github.com/cavaliercoder/go-zabbix"
	"github.com/cavaliercoder/go-zabbix/test"
)

func TestEvents(t *testing.T) {
	session := test.GetTestSession(t)

	params := zabbix.EventGetParams{
		SelectAcknowledgements: zabbix.SelectExtendedOutput,
		SelectAlerts:           zabbix.SelectExtendedOutput,
		SelectHosts:            zabbix.SelectExtendedOutput,
		SelectRelatedObject:    zabbix.SelectExtendedOutput,
	}

	events, err := session.GetEvents(params)
	if err != nil {
		if _, ok := err.(*zabbix.NotFoundError); !ok {
			t.Fatalf("Error getting events: %v", err)
		}
	}

	if len(events) == 0 {
		t.Skip("No events found")
	}

	for i, event := range events {
		if event.EventID == "" {
			t.Fatalf("Event %d has no Event ID", i)
		}

		if event.Timestamp().IsZero() {
			t.Fatalf("Event %d has no timestamp", i)
		}
	}

	t.Logf("Validated %d Events", len(events))
}
