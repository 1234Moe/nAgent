package structType

type ComicDownload struct {
	ComicId           int
	OutputDir         string
	DirNameUseComicId bool
	ThreadCount       int
	ProxyUrl          string
	Retry             bool
	Zip               bool
}
