package pkg_test

import (
	"reflect"
	"testing"

	"github.com/asahnoln/event-notifier/pkg"
)

type stubStore struct {
	events map[pkg.EventType][]pkg.Event
}

func (s *stubStore) Events(when pkg.EventType) []pkg.Event {
	return s.events[when]
}

var db = &stubStore{
	map[pkg.EventType][]pkg.Event{
		pkg.Tomorrow: {
			{
				What:  "Scene 2.5",
				Where: "Dom, Great Hall",
				Who:   []string{"Ivan", "Erkanat"},
			},
		},
		pkg.Today: {
			{
				What:  "Scene 2.6",
				Where: "Dom, Second Room",
				Who:   []string{"Varvara", "Kamila"},
			},
		},
	},
}

func TestGetTomorrowEvents(t *testing.T) {
	es := pkg.TomorrowEvents(db)
	assertSameLength(t, 1, len(es))

	assertSameStruct(t, db.events[pkg.Tomorrow][0], es[0])
}

func TestGetTodayEvents(t *testing.T) {
	es := pkg.TodayEvents(db)

	assertSameStruct(t, db.events[pkg.Today][0], es[0])
}

func assertSameStruct(t testing.TB, want, got interface{}) {
	t.Helper()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want event structure %+v, got %+v", want, got)
	}
}

func assertSameLength(t testing.TB, want, got int) {
	t.Helper()

	if want != got {
		t.Fatalf("want events length %d, got %d", want, got)
	}
}
