# Entropy Key Scanner (Core)

**Description**
High-performance security engine written in Go with C integration. It uses Shannon Entropy calculation via CGO to detect high-entropy strings, such as private keys and API tokens, that often bypass standard regex-based scanners.

**Key Features**
* Hybrid Detection: Combines optimized regular expressions with statistical entropy analysis.
* Concurrency: Multithreaded file processing using a Worker Pool pattern.
* Low Latency: Core entropy calculations implemented in C for maximum throughput.
* Bilingual: Native support for English and Spanish logs.

**Installation**
Requires Go 1.22+ and a C compiler (GCC/Clang).
```bash
go build -o entropy-key-scanner ./cmd/engine/main.go
```

**Usage**
```bash
./entropy-key-scanner -workers 8 -lang en ./target-directory
```

**JSON Output**
For machine-to-machine integration, use the json flag:
```bash
./entropy-key-scanner -json ./path
```
