package do

type LabelNotices struct {
	Key            string   `json:"key"`
	Value          string   `json:"value"`
	ReceiverRoutes []string `json:"receiverRoutes"`
}

func (l *LabelNotices) GetKey() string {
	if l == nil {
		return ""
	}
	return l.Key
}

func (l *LabelNotices) GetValue() string {
	if l == nil {
		return ""
	}
	return l.Value
}

func (l *LabelNotices) GetReceiverRoutes() []string {
	if l == nil {
		return nil
	}
	return l.ReceiverRoutes
}
