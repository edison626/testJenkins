package main

import (
	"fmt"
	"os"
)

func main() {
	varJmsServerURL := os.Getenv("JmsServerURL")
	fmt.Printf("JmsServerURL : %s\n", varJmsServerURL)
	varJMSToken := os.Getenv("JMSToken")
	fmt.Printf("JMSToken : %s\n", varJMSToken)
	varBatch := os.Getenv("JmsServerURL")
	fmt.Printf("Batch : %s\n", varBatch)
	varAssetNote := os.Getenv("AssetNote")
	fmt.Printf("AssetNote : %s\n", varAssetNote)
	varAssetNodeDisplay := os.Getenv("AssetNodeDisplay")
	fmt.Printf("AssetNodeDisplay : %s\n", varAssetNodeDisplay)
}
