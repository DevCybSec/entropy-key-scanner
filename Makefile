BINARY_NAME=entropy-key-scanner
GO=go

build:
	# Go detectará los archivos .c en la carpeta del paquete y los compilará automáticamente
	CGO_ENABLED=1 $(GO) build -o $(BINARY_NAME) ./cmd/engine/main.go

clean:
	rm -f $(BINARY_NAME)