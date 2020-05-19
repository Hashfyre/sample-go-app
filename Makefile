dev:
	export GO111MODULE=on
	go run -mod=vendor main.go
review:
	@{\
		echo "Reviewdog needs the following go packages to be installed \n \
					go get -u github.com/golang/lint \n  \
					go get -u github.com/dominikh/go-tools \n \
					go get -u github.com/kisielk/errcheck \n"; \
		reviewdog -diff="git diff master" -level="info"; \
   }
test:
	@{ \
	    set -e; \
	    export GO111MODULE=on;\
	    echo "" > coverage.txt;\
	    GIT_TLD=`git rev-parse --show-toplevel`;\
	    cd $$GIT_TLD;\
      for d in $$(go list -mod=vendor ./... | grep -v "vendor\|bin\|docs"); do \
	        go test -mod=vendor -coverprofile=profile.out -coverpkg=$$d $$d;\
	        if [ -f profile.out ]; then \
	            cat profile.out >> coverage.txt;\
	            rm profile.out;\
	        fi \
	    done \
	}
build:
	export G0111MODULE=on
	mkdir -p ./bin
	go build -mod=vendor -o ./bin/
docker_up:
	docker-compose up
docker_down:
	docker-compose down
