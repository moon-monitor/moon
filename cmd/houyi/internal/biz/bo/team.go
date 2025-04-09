package bo

type TeamItem struct {
	TeamId uint32
	Uuid   string
}

type LabelNotices struct {
	Key            string
	Value          string
	ReceiverRoutes []string
}
