package randomOrg

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
)

func GetDiceRoll(n uint8) ([]uint8, error) {
	url := os.Getenv("RO_URI")
	apiKey := os.Getenv("RO_APIKEY")
	id := rand.Intn(10000)
	var jsonStr = []byte(fmt.Sprintf(`
		{
			"jsonrpc": "2.0",
			"method": "generateIntegers",
			"params": {
				"apiKey": "%s",
				"n": %d,
				"min": 1,
				"max": 6,
				"replacement": true
			},
			"id": %d
		}
	`, apiKey, n, id))

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// fmt.Println("response Body:", string(body))

	var result DiceResponse
	json.Unmarshal(body, &result)
	if result.ID != id {
		return nil, errors.New("result id different than request - tampering suspected")
	}
	fmt.Println("Random.org Requests left in the next 24 hours: ", result.Result.RequestsLeft)
	return []uint8(result.Result.Random.Data), nil
}
