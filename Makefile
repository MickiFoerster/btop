GO_FILES=main.go \
		 cpu.go \
		 http.go
TEMPLATES=$(shell ls -1 templates/*.gotpl)

btop: $(GO_FILES) Makefile $(TEMPLATES)
	go build -o $@ $(GO_FILES)

