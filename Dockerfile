FROM golang:1.22-alpine AS builder
LABEL authors="edgar(chore)macias"
# Instalar dependencias necesarias para CGO (gcc, musl-dev para C)
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Copiar archivos de dependencia primero para aprovechar el cache
COPY go.mod ./
# RUN go mod download (si tuvieras dependencias externas)

# Copiar el resto del código
COPY . .

# Compilar usando el Makefile que creamos
RUN apk add --no-cache make
RUN make build

# --- ETAPA 2: Imagen Final (Producción) ---
FROM alpine:latest

RUN apk add --no-cache libc6-compat

WORKDIR /root/

# Copiar solo el binario final desde la etapa anterior
COPY --from=builder /app/high-perf-scanner .

# Comando por defecto
ENTRYPOINT ["./high-perf-scanner"]