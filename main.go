package main

import (
	"fmt"
	"log"
	"os"
)

func TestDisplay(varJmsServerURL string, varJMSToken string, varBatch string, varAssetNote string, varAssetNodeDisplay string) {
	fmt.Printf("JmsServerURL : %s\n", varJmsServerURL)
	fmt.Printf("JMSToken : %s\n", varJMSToken)
	fmt.Printf("Batch : %s\n", varBatch)
	fmt.Printf("AssetNote : %s\n", varAssetNote)
	fmt.Printf("AssetNodeDisplay : %s\n", varAssetNodeDisplay)
}

func main() {
	varJmsServerURL := os.Getenv("JmsServerURL")
	varJMSToken := os.Getenv("JMSToken")
	varBatch := os.Getenv("Batch")
	varAssetNote := os.Getenv("AssetNote")
	varAssetNodeDisplay := os.Getenv("AssetNodeDisplay")

	if varJmsServerURL == "" || varJMSToken == "" || varBatch == "" || varAssetNote == "" || varAssetNodeDisplay == "" {
		log.Fatalf("值不能为空")
	}

	TestDisplay(varJmsServerURL, varJMSToken, varBatch, varAssetNote, varAssetNodeDisplay)
	fmt.Printf("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
}
