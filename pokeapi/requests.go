package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetURl(url string) ([]byte, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("Failed to create request: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, fmt.Errorf("Request Failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("Request returned non-200 status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return []byte{}, fmt.Errorf("Failed to pasrse stream: %w", err)
	}

	return data, nil
}

func UnmarshalData[T any](jsonData []byte) (T, error) {
	var deserializedData T
	err := json.Unmarshal(jsonData, &deserializedData)
	if err != nil {
		return deserializedData, fmt.Errorf("Could not Unmarshal Data: %w", err)
	}
	return deserializedData, nil
}

func DecodeData[T any](stream io.Reader) (T, error) {
	var deserializedData T
	decoder := json.NewDecoder(stream)
	err := decoder.Decode(&deserializedData)
	if err != nil {
		return deserializedData, fmt.Errorf("Error Decoding Data: %w", err)
	}
	return deserializedData, nil
}
