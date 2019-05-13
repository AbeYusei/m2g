package m2g

// Err error message ID
type Err int

const (
	// InvalidArgLength if arguments specified was too much.
	InvalidArgLength Err = iota
	// InvalidArgFile if specified file was unexpected.
	InvalidArgFile
	// NotExistsInputFile if specified file was not found.
	NotExistsInputFile
	// AlraedyExistsOutputFile if output file was already existed.
	AlraedyExistsOutputFile
	// InternalProcess if other lib process throws error.
	InternalProcess
)

// Error get error message.
func (m Err) Error() string {
	switch m {
	case InvalidArgLength:
		return "引数に変換対象の.movファイルを指定してください。"
	case InvalidArgFile:
		return ".movファイルを指定してください。"
	case NotExistsInputFile:
		return "指定された.movファイルが存在しません。"
	case AlraedyExistsOutputFile:
		return ".gifファイルが既に存在します。"
	case InternalProcess:
		return "ffmpegのプロセス内でエラーが発生しました。"
	default:
		return "Unknown Error."
	}
}
