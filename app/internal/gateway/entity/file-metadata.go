package entity

type FileMetadata struct {
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
	Content  []byte `json:"content"`
}
