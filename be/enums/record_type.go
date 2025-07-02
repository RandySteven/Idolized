package enums

type RecordType string

const (
	LiveStreaming    RecordType = `Live`
	PodcastStreaming RecordType = `Podcast`
	AudioStreaming   RecordType = `Audio`
)
