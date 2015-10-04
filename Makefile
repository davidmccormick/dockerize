.SILENT :
.PHONY : template clean fmt

TAG:=`1.0`
LDFLAGS:=-X main.buildVersion $(TAG)

all: template

template:
	echo "Building template"
	go install -ldflags "$(LDFLAGS)"

dist-clean:
	rm -rf dist
	rm -f template-linux-*.tar.gz

dist: dist-clean
	mkdir -p dist/linux/amd64 && GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o dist/linux/amd64/template
	mkdir -p dist/linux/armel && GOOS=linux GOARCH=arm GOARM=5 go build -ldflags "$(LDFLAGS)" -o dist/linux/armel/template
	mkdir -p dist/linux/armhf && GOOS=linux GOARCH=arm GOARM=6 go build -ldflags "$(LDFLAGS)" -o dist/linux/armhf/template

release: dist
	tar -cvzf template-linux-amd64-$(TAG).tar.gz -C dist/linux/amd64 template
	tar -cvzf template-linux-armel-$(TAG).tar.gz -C dist/linux/armel template
	tar -cvzf template-linux-armhf-$(TAG).tar.gz -C dist/linux/armhf template
