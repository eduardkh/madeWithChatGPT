# go_config_management

> initialize the project

```bash
go mod init go_config_management
cobra-cli init
go mod tidy
go run main.go

# run the app
go run main.go add
# 1. default value - when no value is provided
# 2. configuration file - when a configuration file is provided (config.yaml)
# 3. environment variable - when an environment variable is provided
# 4. CLI flag - when a CLI flag is provided
```
