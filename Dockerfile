FROM golang:1.22.1-alpine

# Instalar o Air
RUN go install github.com/air-verse/air@latest

# Definir o diretório de trabalho
WORKDIR /app

# Criar um diretório temporário
RUN mkdir /app/tmp

# Copiar os arquivos go.mod e go.sum
COPY go.* ./

# Baixar as dependências
RUN go mod download

# Copiar o código-fonte
COPY . .

# Expor a porta 8080
EXPOSE 8080

# Executar o Air com o arquivo de configuração .air.toml
CMD ["air", "-c", ".air.toml"]
