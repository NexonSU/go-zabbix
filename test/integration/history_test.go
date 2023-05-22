package integration

import (
	"testing"

	"github.com/fabiang/go-zabbix"
	"github.com/fabiang/go-zabbix/test"
)

func TestHistories(t *testing.T) {
	session := test.GetTestSession(t)

	params := zabbix.HistoryGetParams{}

	histories, err := session.GetHistories(params)
	if err != nil {
		if _, ok := err.(*zabbix.NotFoundError); !ok {
			t.Fatalf("Error getting histories: %v", err)
		}
	}

	if len(histories) == 0 {
		t.Skip("No histories found")
	}

	for i, history := range histories {
		if history.ItemID == 0 {
			t.Fatalf("History %d has no item ID", i)
		}
	}

	t.Logf("Validated %d histories", len(histories))
}
