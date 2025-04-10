package event_test

import (
	"testing"
	"time"

	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/event"
	"github.com/moon-monitor/moon/pkg/api/houyi/common"
	"github.com/moon-monitor/moon/pkg/util/kv/label"
	"github.com/moon-monitor/moon/pkg/util/pointer"
)

func Test_AlertJob(t *testing.T) {
	alert := &event.AlertJob{
		Status: common.EventStatus_pending,
		Labels: label.NewLabel(map[string]string{
			"alertname": "TestAlert",
			"severity":  "warning",
			"job":       "test-job",
		}),
		Annotations:  label.NewAnnotation("summary", "this is a test alert"),
		StartsAt:     pointer.Of(time.Now().UTC()),
		EndsAt:       nil,
		GeneratorURL: "",
		Fingerprint:  "",
	}
	marshalBinary, err := alert.MarshalBinary()
	if err != nil {
		t.Error(err)
	}
	t.Log(string(marshalBinary))
	var newAlert event.AlertJob
	err = newAlert.UnmarshalBinary(marshalBinary)
	if err != nil {
		t.Error(err)
	}
	newAlert.StatusNext()
	marshalBinary, err = newAlert.MarshalBinary()
	if err != nil {
		t.Error(err)
	}
	t.Log(string(marshalBinary))
	newAlert.StatusNext()
	marshalBinary, err = newAlert.MarshalBinary()
	if err != nil {
		t.Error(err)
	}
	t.Log(string(marshalBinary))
	t.Log(newAlert.GetFingerprint())
}
