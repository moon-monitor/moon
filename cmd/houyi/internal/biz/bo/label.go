package bo

type LabelNotices struct {
	Key            string
	Value          string
	ReceiverRoutes []string
}

type Label map[string]string

type Annotation map[string]string
