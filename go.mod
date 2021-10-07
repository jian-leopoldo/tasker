module tasker

go 1.17

require (
	tasker/runner v1.0.0-0000000000000-00000000
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b

) // indirect

replace tasker/runner => ./src/runners
