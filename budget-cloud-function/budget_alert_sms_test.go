package budgetalertsms

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/cloudevents/sdk-go/v2/event"
)

func TestBudgetAlertSms(t *testing.T) {
	tests := []struct {
		data string
		want string
	}{
		{want: "Hello, World!\n"},
		{data: "Go", want: "Hello, Go!\n"},
	}
	for _, test := range tests {
		r, w, _ := os.Pipe()
		log.SetOutput(w)
		originalFlags := log.Flags()
		log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

		m := PubSubMessage{
			Data: []byte(test.data),
		}
		msg := MessagePublishedData{
			Message: m,
		}
		e := event.New()
		e.SetDataContentType("application/json")
		e.SetData(e.DataContentType(), msg)

		budgetAlertSms(context.Background(), e)

		w.Close()
		log.SetOutput(os.Stderr)
		log.SetFlags(originalFlags)

		out, err := ioutil.ReadAll(r)
		if err != nil {
			t.Fatalf("ReadAll: %v", err)
		}
		if got := string(out); got != test.want {
			t.Errorf("BudgetAlertSms(%q) = %q, want %q", test.data, got, test.want)
		}
	}
}
