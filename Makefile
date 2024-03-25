# main target
all: build-css build-go

# build css with tailwind
build-css:
	pnpm run tw:prod

# build Go application
build-go:
	go build -o ./build/potoforio ./main.go

# cleans up prod css
clean-css:
	rm -f ./static/styles.css

# clean all
clean: clean-css
	rm -f ./build/potoforio
