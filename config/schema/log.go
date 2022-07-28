package schema

type Log struct {
	Level   string         `json:"level"`
	Console ConsoleLogging `json:"console"`
	File    FileLogging    `json:"file"`
}

type ConsoleLogging struct {
	Enable bool `json:"enable"`
	Json   bool `json:"json"`
}

type FileLogging struct {
	Enable   bool   `json:"enable"`
	Filename string `json:"filename"`
	MaxSize  int    `json:"max_size"`
	MaxFiles int    `json:"max_files"`
	MaxAge   int    `json:"max_age"`
}
