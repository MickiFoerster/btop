GO_FILES=main.go \
		 cpu.go \
		 http.go
TEMPLATES=$(shell ls -1 templates/*.gotpl)

gotopws: $(GO_FILES) Makefile
	go build -o $@ $(GO_FILES)

http.go: $(TEMPLATES)
