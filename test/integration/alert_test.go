package integration

import (
	"testing"

	"github.com/fabiang/go-zabbix"
	"github.com/fabiang/go-zabbix/test"
)

func TestAlerts(t *testing.T) {
	session := test.GetTestSession(t)

	params := zabbix.AlertGetParams{
		SelectHosts: zabbix.SelectExtendedOutput,
	}

	alerts, err := session.GetAlerts(params)
	if err != nil {
		if _, ok := err.(*zabbix.NotFoundError); !ok {
			t.Fatalf("Error getting alerts: %v", err)
		}
	}

	if len(alerts) == 0 {
		t.Skip("No alerts found")
	}

	for i, alert := range alerts {
		if alert.AlertID == "" {
			t.Fatalf("Alert %d has no Alert ID", i)
		}
	}

	t.Logf("Validated %d Alerts", len(alerts))
}
