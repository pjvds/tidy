package logentries_test

import (
	"fmt"
	"testing"

	"github.com/pjvds/tidy"
	"github.com/pjvds/tidy/logentries"
	"github.com/stvp/go-udp-testing"
)

func TestLogentriesBackend(t *testing.T) {
	address := "127.0.0.1:8125"
	token := "2bfbea1e-10c3-4419-bdad-7e6435882e1f"
	udp.SetAddr(address)

	expected := fmt.Sprintf("%vDEBUG (module): foobar\n", token)
	udp.ShouldReceiveOnly(t, expected, func() {
		backend := logentries.Configure(token).Address(address).Build()
		log := tidy.NewLogger("module", backend)

		log.Debug("foobar")
	})
}

/*
func getLogTail(key string) (string, error) {
	url := fmt.Sprintf("https://pull.logentries.com/%v/hosts/ManualHost/tidy_tests/?start=-100000", key)
	response, err := http.Get(url)

	if err != nil {
		return "", err
	}

	all, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	return string(all), nil
}

func BenchmarkBackend(b *testing.B) {
	token := os.Getenv("LE_TOKEN")
	if len(token) == 0 {
		b.Skip("LE_TOKEN not set")
	}

	key := os.Getenv("LE_ACCOUNT_KEY")
	if len(key) == 0 {
		b.Skip("LE_ACCOUNT_KEY not set")
	}

	backend := logentries.Configure(token).TCP().Build()
	entry := tidy.Entry{
		Timestamp: time.Now(),
		Module:    tidy.Module("test"),
		Level:     tidy.FATAL,
		Message:   "log message",
	}

	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		backend.Log(entry)
	}
}

func TestBackendRoundtrip(t *testing.T) {
	token := os.Getenv("LE_TOKEN")
	if len(token) == 0 {
		t.Skip("LE_TOKEN not set")
	}

	key := os.Getenv("LE_ACCOUNT_KEY")
	if len(key) == 0 {
		t.Skip("LE_ACCOUNT_KEY not set")
	}

	backend := logentries.Configure(token).UDP().Build()
	log := tidy.NewLogger("foobar", backend)

	id, _ := uuid.NewV4()
	secret := id.String()

	for n := 0; n < 10; n++ {
		log.WithFields(tidy.Fields{
			"secret": secret,
			"n":      n,
		}).Debug("hello world")
	}

	var lastBody atomic.Value
	lastBody.Store("<empty>")

	done := make(chan struct{})
	matched := make(chan struct{})

	go func() {
		t.Logf("looking for secret: %v", secret)
		for {
			select {
			case <-done:
				return
			default:
				tail, err := getLogTail(key)
				if err != nil {
					t.Fatalf("failed to get tail from logentries: %v", err)
				}

				lastBody.Store(tail)
				if strings.Contains(tail, secret) {
					close(matched)
					return
				}

				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	select {
	case <-matched:
		// great
	case <-time.After(45 * time.Second):
		t.Logf("body: %v", lastBody.Load())
		t.Fatalf("entry not found in tail")
		close(done)
	}
}
*/
