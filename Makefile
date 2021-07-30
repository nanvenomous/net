pkgname="net"
fpath="/usr/local/share/zsh/site-functions"

## build: compile bt executable in current directory from source
build:
	go mod tidy
	./shellscripts/init_conifg.sh
	go build -o "${pkgname}" main.go
	sudo mkdir -p "${fpath}"
	./"${pkgname}" --completion zsh | sudo tee "${fpath}/_${pkgname}" > /dev/null

	sudo cp ./"${pkgname}" "${GOROOT}"/bin/"${pkgname}"


.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
