package wordnik

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type WordnikResponse struct {
	Word        string `json:"word"`
	Definitions []struct {
		Text         string `json:"text"`
		PartOfSpeech string `json:"partOfSpeech"`
		Source       string `json:"source"`
	} `json:"definitions"`
	Examples []struct {
		Text  string `json:"text"`
		Title string `json:"title"`
	} `json:"examples"`
}

type WordDefinition struct {
	Text         string       `json:"text"`
	PartOfSpeech string       `json:"partOfSpeech"`
	ExampleUses  []ExampleUse `json:"exampleUses"`
}

type ExampleUse struct {
	Text string `json:"text"`
}

type WordPronunciations struct {
	
}

func GetWordOfTheDay() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("[GetWordOfTheDay] Error loading .env file: %v", err)
	}

	wordnikURL := os.Getenv("WORDNIK_URL")
	if wordnikURL == "" {
		log.Fatalf("[GetWordOfTheDay] Missing WORDNIK_URL in environment variables.")
	}

	wordnikToken := os.Getenv("WORDNIK_TOKEN")
	if wordnikToken == "" {
		log.Fatalf("[GetWordOfTheDay] Missing WORDNIK_TOKEN in environment variables.")
	}

	apiURL := fmt.Sprintf("%s/words.json/wordOfTheDay?api_key=%s", wordnikURL, wordnikToken)

	resp, err := http.Get(apiURL)
	if err != nil {
		log.Printf("[GetWordOfTheDay] Error making API request: %v", err)
		return "⚠️ Could not fetch the word of the day. Please try again later."
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[GetWordOfTheDay] Error reading API response: %v", err)
		return "⚠️ Could not read response from Wordnik API."
	}

	log.Printf("[GetWordOfTheDay] Raw API response: %s", body)

	var result WordnikResponse
	if err := json.Unmarshal(body, &result); err != nil {
		log.Printf("[GetWordOfTheDay] Error parsing JSON response: %v", err)
		return "⚠️ Could not parse the Wordnik API response."
	}

	if result.Word == "" {
		log.Println("[GetWordOfTheDay] No word found in API response.")
		return "⚠️ No word of the day found."
	}

	var definition string
	if len(result.Definitions) > 0 {
		definition = fmt.Sprintf("_Definition_: %s (%s)", result.Definitions[0].Text, result.Definitions[0].PartOfSpeech)
	} else {
		definition = "⚠️ No definition found."
	}

	var example string
	if len(result.Examples) > 0 {
		example = fmt.Sprintf("_Example_: \"%s\" - %s", result.Examples[0].Text, result.Examples[0].Title)
	} else {
		example = "⚠️ No example found."
	}

	return fmt.Sprintf("📖 **Word of the Day**: **%s**\n%s\n%s", result.Word, definition, example)
}

func GetWord(word string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("[GetWord] Error loading .env file: %v", err)
	}

	wordnikURL := os.Getenv("WORDNIK_URL")
	if wordnikURL == "" {
		log.Fatalf("[GetWord] Missing WORDNIK_URL in environment variables.")
	}

	wordnikToken := os.Getenv("WORDNIK_TOKEN")
	if wordnikToken == "" {
		log.Fatalf("[GetWord] Missing WORDNIK_TOKEN in environment variables.")
	}

	apiURL := fmt.Sprintf("%s/word.json/%s/definitions?limit=3&api_key=%s", wordnikURL, word, wordnikToken)

	resp, err := http.Get(apiURL)
	if err != nil {
		log.Printf("[GetWord] Error making request to Wordnik: %v", err)
		return "⚠️ Error fetching word definition."
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("[GetWord] Wordnik API returned status: %d", resp.StatusCode)
		return "⚠️ No definition found."
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[GetWord] Error reading response body: %v", err)
		return "⚠️ Error reading response."
	}

	var definitions []WordDefinition
	err = json.Unmarshal(body, &definitions)
	if err != nil {
		log.Printf("[GetWord] Error parsing JSON: %v", err)
		return "⚠️ Error parsing response."
	}

	if len(definitions) == 0 {
		log.Println("[GetWord] No definitions found.")
		return "⚠️ No definition found."
	}

	var result strings.Builder
	result.WriteString(fmt.Sprintf("📖 **%s**\n", word))

	for i, def := range definitions {
		if i >= 3 {
			break
		}
		result.WriteString(fmt.Sprintf("**%d.** *(%s)* %s\n", i+1, def.PartOfSpeech, def.Text))
		if len(def.ExampleUses) > 0 {
			result.WriteString(fmt.Sprintf("_Example:_ \"%s\"\n", def.ExampleUses[0].Text))
		}
	}

	return result.String()
}

func GetEtymologies(word string) ([]string, error) {
	err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("[GetEtymologies] Error loading.env file: %v", err)
    }

    wordnikURL := os.Getenv("WORDNIK_URL")
    if wordnikURL == "" {
        log.Fatalf("[GetEtymologies] Missing WORDNIK_URL in environment variables.")
    }

    wordnikToken := os.Getenv("WORDNIK_TOKEN")
    if wordnikToken == "" {
        log.Fatalf("[GetEtymologies] Missing WORDNIK_TOKEN in environment variables.")
    }

    apiURL := fmt.Sprintf("%s/word.json/%s/etymologies?useCanonical=false&api_key=%s", wordnikURL, word, wordnikToken)

	resp, err := http.Get(apiURL)
	if err != nil {
        log.Printf("[GetEtymologies] Error making request to Wordnik: %v", err)
        return nil, fmt.Errorf("error fetching etymologies")
    }


    if resp.StatusCode != http.StatusOK {
        log.Printf("[GetEtymologies] Wordnik API returned status: %d", resp.StatusCode)
        return nil, fmt.Errorf("no etymologies found")
    }

	body, err := io.ReadAll(resp.Body)
	if err != nil {
        log.Printf("[GetEtymologies] Error reading response body: %v", err)
        return nil, fmt.Errorf("error reading response")
    }

	var etymologies []string
	err = json.Unmarshal(body, &etymologies)
	if err != nil {
        log.Printf("[GetEtymologies] Error parsing JSON: %v", err)
        return nil, fmt.Errorf("error parsing response")
    }

	return etymologies, nil
}
