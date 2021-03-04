pre:
	@echo "=== Bund === [Preinstallation of some stuff]"
	go get github.com/client9/misspell/cmd/misspell
	go mod download golang.org/x/tools
	go get github.com/llorllale/go-gitlint/cmd/go-gitlint
	go get github.com/psampaz/go-mod-outdated
	go mod download golang.org/x/tools
	go get gotest.tools/gotestsum
	go get golang.org/x/tools/internal/imports@v0.0.0-20200522201501-cb1345f3a375
	go get golang.org/x/tools/internal/imports@v0.0.0-20200522201501-cb1345f3a375
