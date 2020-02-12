run:
	go get -u github.com/koron/go-github
	go get -u gopkg.in/yaml.v2
	go run _scripts/vim_jp-redirects-update/main.go
