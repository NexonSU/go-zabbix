package integration

import (
	"testing"

	"github.com/fabiang/go-zabbix"
	"github.com/fabiang/go-zabbix/test"
)

func TestItemsIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	session := test.GetTestSession(t)

	params := zabbix.ItemGetParams{}

	items, err := session.GetItems(params)
	if err != nil {
		if _, ok := err.(*zabbix.NotFoundError); !ok {
			t.Fatalf("Error getting items: %v", err)
		}
	}

	if len(items) == 0 {
		t.Skip("No items found")
	}

	for i, item := range items {
		if item.ItemID == "" {
			t.Fatalf("Item %d has no item ID", i)
		}
	}

	t.Logf("Validated %d items", len(items))
}
