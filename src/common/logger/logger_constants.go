package logger

//Log Levels
const (
	DebugLevel   int = 5 //Level 5
	InfoLevel    int = 4 //Level 4
	TraceLevel   int = 3 //Level 3
	WarningLevel int = 2 //Level 2
	ErrLevel     int = 1 //Level 1
	Skip         int = 4 //Skip level where to end tracing files in stackTrace
	CallingDepth int = 3 //Depth level till the calling method in stackTrace
)
