package main

import (
	"os"
    "fmt"
    "github.com/elliotwutingfeng/go-fasttld"
    "encoding/json"
)


func main() {
    // Initialise fasttld extractor
    extractor, _ := fasttld.New(fasttld.SuffixListParams{})

    // Extract URL subcomponents
    url := os.Args[1]
    res, _ := extractor.Extract(fasttld.URLParams{URL: url, ConvertURLToPunyCode: true})
    
    // Display results
    // fasttld.PrintRes(url, res) // Pretty-prints res.Scheme, res.UserInfo, res.SubDomain etc.
    // print(res.SubDomain,res.RegisteredDomain)
    
    data := map[string]string{
    	"sub_domain": res.SubDomain,
    	"reg_domain": res.RegisteredDomain,
        "root":       res.Domain,
        "tld":        res.Suffix,
        "full_domain":res.SubDomain+"."+res.RegisteredDomain,
    }

	// 编码为 JSON 格式的字节片
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON encoding error:", err)
		return
	}

	fmt.Println(string(jsonBytes))
}