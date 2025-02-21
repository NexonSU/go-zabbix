package zabbix_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/fabiang/go-zabbix"
	"github.com/fabiang/go-zabbix/test"
	"github.com/fabiang/go-zabbix/types"
)

const (
	fakeURL   = "http://localhost/api_jsonrpc.php"
	fakeToken = "0424bd59b807674191e7d77572075f33"
)

var fakeAPIVersion *types.ZBXVersion

func init() {
	fakeAPIVersion, _ = types.NewZBXVersion("2.0")
}

func prepareTemporaryDir(t *testing.T) (dir string, success bool) {
	tempDir, err := os.MkdirTemp("", "zabbix-session-test")

	if err != nil {
		t.Fatalf("cannot create a temporary dir for session cache: %v", err)
		return "", false
	}

	t.Logf("used %s directory as temporary dir", tempDir)

	return tempDir, true
}

func getTestFileCache(baseDir string) zabbix.SessionAbstractCache {
	sessionFilePath := baseDir + "/" + ".zabbix_session"
	return zabbix.NewSessionFileCache().SetFilePath(sessionFilePath)
}

func TestSessionCache(t *testing.T) {
	// Create a fake session for r/w test
	fakeSession := &zabbix.Session{
		URL:        fakeURL,
		Token:      fakeToken,
		APIVersion: fakeAPIVersion,
	}

	tempDir, success := prepareTemporaryDir(t)

	if !success {
		return
	}

	cache := getTestFileCache(tempDir)

	if err := cache.SaveSession(fakeSession); err != nil {
		t.Errorf("failed to save mock session - %v", err)
		return
	}

	if !cache.HasSession() {
		t.Errorf("session was saved but not detected again by cache")
		return
	}

	// Try to get a cached session
	cachedSession, err := cache.GetSession()

	if err != nil {
		t.Error(err)
		return
	}

	// Check session integrity
	if err := compareSessionWithMock(cachedSession); err != nil {
		t.Error(err)
	}

	testClientBuilder(t, cache)

	if err := cache.Flush(); err != nil {
		t.Error("failed to remove a cached session file")
	}
}

func compareSessionWithMock(session *zabbix.Session) error {
	if session.URL != fakeURL {
		return fmt.Errorf("Session URL '%s' is not equal to '%s'", session.URL, fakeURL)
	}

	if session.Token != fakeToken {
		return fmt.Errorf("Session token '%s' is not equal to '%s'", session.Token, fakeToken)
	}

	if session.APIVersion.String() != fakeAPIVersion.String() {
		return fmt.Errorf(
			"Session version %q is not equal to %q",
			session.APIVersion.String(),
			fakeAPIVersion.String())
	}

	return nil
}

// should started by TestSessionCache
func testClientBuilder(t *testing.T, cache zabbix.SessionAbstractCache) {
	username, password, url := test.GetTestCredentials()

	if !cache.HasSession() {
		t.Errorf("ManualTestClientBuilder test requires a cached session, run TestSessionCache before running this test case")
		return
	}

	// Try to build a session using the session builder
	client, err := zabbix.CreateClient(url).WithCache(cache).WithCredentials(username, password).Connect()

	if err != nil {
		t.Errorf("failed to create a session using cache - %s", err)
		return
	}

	if err := compareSessionWithMock(client); err != nil {
		t.Error(err)
	}
}
