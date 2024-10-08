package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ApiUrl 		= ""
	Porta 		= 0
	HashKey		[]byte
	Block		[]byte
)

func Carregar() {
	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}
	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}

	ApiUrl = os.Getenv("API_URL")

	HashKey = []byte(os.Getenv("HASH_KEY"))
	if len(HashKey) != 16 && len(HashKey) != 24 && len(HashKey) != 32 {
		log.Fatalf("HashKey deve ter 16, 24 ou 32 bytes, mas tem %d bytes", len(HashKey))
	}

	Block = []byte(os.Getenv("BLOCK_KEY"))
	if len(Block) != 16 && len(Block) != 24 && len(Block) != 32 {
		log.Fatalf("BlockKey deve ter 16, 24 ou 32 bytes, mas tem %d bytes", len(Block))
	}
}