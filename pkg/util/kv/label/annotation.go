package label

import (
	"encoding/json"

	"github.com/moon-monitor/moon/pkg/util/cnst"
	"github.com/moon-monitor/moon/pkg/util/kv"
	"github.com/moon-monitor/moon/pkg/util/template"
)

func NewAnnotation(summary, description string) *Annotation {
	return &Annotation{
		kvMap: kv.NewStringMap(map[string]string{
			cnst.AnnotationKeySummary:     summary,
			cnst.AnnotationKeyDescription: description,
		}),
	}
}

type Annotation struct {
	kvMap kv.StringMap
}

func (a *Annotation) String() string {
	bs, _ := a.MarshalBinary()
	return string(bs)
}

func (a *Annotation) MarshalBinary() (data []byte, err error) {
	return json.Marshal(a.kvMap)
}

func (a *Annotation) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &a.kvMap)
}

func (a *Annotation) GetSummary() string {
	summary, ok := a.kvMap.Get(cnst.AnnotationKeySummary)
	if !ok {
		return ""
	}
	return summary
}

func (a *Annotation) SetSummary(summary string) {
	a.kvMap.Set(cnst.AnnotationKeySummary, summary)
}

func (a *Annotation) GetDescription() string {
	description, ok := a.kvMap.Get(cnst.AnnotationKeyDescription)
	if !ok {
		return ""
	}
	return description
}

func (a *Annotation) SetDescription(description string) {
	a.kvMap.Set(cnst.AnnotationKeyDescription, description)
}

func (a *Annotation) Format(data interface{}) *Annotation {
	for k, v := range a.kvMap {
		a.kvMap.Set(k, template.TextFormatterX(v, data))
	}
	return a
}
