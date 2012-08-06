TARGET := does_hsts

all: $(TARGET)

$(TARGET): $(TARGET).go
	go build -o $(TARGET)

clean:
	rm $(TARGET)

.PHONY: clean all

