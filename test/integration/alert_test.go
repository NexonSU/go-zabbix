package integration

import (
	"testing"

	"github.com/NexonSU/go-zabbix"
	"github.com/NexonSU/go-zabbix/test"
)

func TestAlertsIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	session := test.GetTestSession(t)

	params := zabbix.AlertGetParams{
		GetParameters: zabbix.GetParameters{
			OutputFields: zabbix.SelectExtendedOutput,
		},
		SelectHosts:      zabbix.SelectExtendedOutput,
		SelectUsers:      zabbix.SelectExtendedOutput,
		SelectMediaTypes: zabbix.SelectExtendedOutput,
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
