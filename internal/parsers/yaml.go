package parsers

type Config struct {
	CsvToJsonConfigs    []CsvToJsonConfig    `yaml:"csv_to_json_configs,flow"`
	GenLocaleKeyConfigs []GenLocaleKeyConfig `yaml:"gen_locale_key_configs,flow"`
	AppPathConfigs      []AppPathConfig      `yaml:"app_path_configs,flow"`
}

type CsvToJsonConfig struct {
	Key       *string `yaml:"key"`
	InputPath string  `yaml:"input_path"`
	OutputDir string  `yaml:"output_dir"`
}

type GenLocaleKeyConfig struct {
	InputDir  string `yaml:"input_dir"`
	OutputDir string `yaml:"output_dir"`
	Namespace string `yaml:"namespace"`
}

type AppPathConfig struct {
	ClassName  string `yaml:"class_name"`
	InputDir   string `yaml:"input_dir"`
	OutputFile string `yaml:"output_file"`
}
