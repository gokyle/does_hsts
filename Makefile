TARGET := does_hsts
PREFIX ?= /usr/local

all: $(TARGET)

$(TARGET): $(TARGET).go
	go build -o $(TARGET)

clean:
	rm $(TARGET)

install: $(TARGET)
	install -m 0755 -d $(PREFIX)/bin
	install -m 0755 $(TARGET)  $(PREFIX)/bin/$(TARGET)

uninstall:
	-rm -f $(PREFIX)/bin/$(TARGET)

.PHONY: clean all

