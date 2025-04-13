package vobj

// SampleMode is the sample mode of the palace.
//
//go:generate stringer -type=SampleMode -linecomment -output=sample_mode.string.go
type SampleMode int8

const (
	For SampleMode = iota // m时间内连续出现n次
	Max                   // 最多出现n次
	Min                   // 最少出现n次
)
