all: pepino
.PHONY: clean

pepino:
	go build

clean:
	test -f pepino && rm pepino

