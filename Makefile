install:
	go build -o  ~/.steampipe/plugins/hub.steampipe.io/plugins/ygpark80/clickup@latest/steampipe-plugin-clickup.plugin *.go

local:
	go build -o  ~/.steampipe/plugins/local/clickup/clickup.plugin *.go
