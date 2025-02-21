package xmlconverter

import (
	"encoding/xml"
	"log"
	"strings"
)

type Etymology struct {
	XMLName xml.Name `xml:"ety"`
	Content []byte   `xml:",innerxml"`
}

func FormatEtymology(xmlData string) string {
	var ety Etymology
	err := xml.Unmarshal([]byte(xmlData), &ety)
	if err != nil {
		log.Printf("Error parsing XML: %v", err)
		return "⚠️ Failed to parse etymology."
	}

	output := string(ety.Content)

	output = strings.ReplaceAll(output, "<ets>", "**")
	output = strings.ReplaceAll(output, "</ets>", "**")

	output = strings.ReplaceAll(output, "<er>", "*")
	output = strings.ReplaceAll(output, "</er>", "*")

	output = strings.ReplaceAll(output, "\n", " ")

	return output
}
