package ocr

// Extractor implements video data extraction related operations.
type OcrEngine interface {
	// Extract is the main function to extract the data.
	Recognize(imgbase64 string) string
}
