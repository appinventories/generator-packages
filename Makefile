test_app: csv_to_json gen_locale_keys app_path_config

csv_to_json:
	go run github.com/appinventories/generator-packages CsvToJson -c _example/config.yaml

gen_locale_keys:	
	go run github.com/appinventories/generator-packages GenLocaleKeys -c _example/config.yaml 

app_path_config:
	go run github.com/appinventories/generator-packages GenAppPath -c _example/config.yaml