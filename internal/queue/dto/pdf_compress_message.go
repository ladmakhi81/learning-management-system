package queuedto

type PDFCompressMessage struct {
	FileName        string
	FileDestination string
}

func NewPDFCompressMessage(fileName, fileDestination string) PDFCompressMessage {
	return PDFCompressMessage{FileName: fileName, FileDestination: fileDestination}
}
