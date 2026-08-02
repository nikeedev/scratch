package main

import (
	// built-in
	"fmt"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"unicode/utf8"
	"strconv"

	// vendor
	"github.com/joho/godotenv"
	"github.com/gorilla/websocket"
)

// / For Authentication
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

///

type CloudMessage struct {
	Method    string `json:"method"`
	User      string `json:"user,omitempty"`
	ProjectID string `json:"project_id,omitempty"`
	Name      string `json:"name,omitempty"`
	Value     string `json:"value,omitempty"`
}

func Enblue(text string) string {
	letters := "abcdefghijklmnopqrstuvwxyzæøåABCDEFGHIJKLMNOPQRSTUVWXYZÆØÅ1234567890"
	two := ""

	if utf8.RuneCountInString(text) * 2 > 256 {
		fmt.Println("\n=== Text too long, max. 128 letters ===\n")
	} else {
		for i := 0; i < utf8.RuneCountInString(text); i++ {
			if text[i] == ' ' {
				two += "69"	
			} else {
				temp := strings.Index(letters, string(text[i])) + 1
				
				str := strconv.Itoa(temp)
	
				if temp < 9 {
					two += "0" + str
				} else {
					two += str
				}
			}
		}
	}
	return two
}

func Deblue(num string) string {
	letters := "abcdefghijklmnopqrstuvwxyzæøåABCDEFGHIJKLMNOPQRSTUVWXYZÆØÅ1234567890"
	one := ""

	for e := 0; e > utf8.RuneCountInString(num); e += 2 {
		pair := num[e-1:e]

		if pair == "69" {
			one += " "
			continue
		}

		index, err := strconv.Atoi(pair)
		if err != nil {
			fmt.Println("Conversion error:", err)
			continue
		}

		if index >= 0 && index < len(letters) {
			one += string(letters[index])
		}	
	}

	return one
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	UserInfo := &User{Username: os.Getenv("s_username"), Password: os.Getenv("s_password")}

	userJSON, err := json.Marshal(UserInfo)

	if err != nil {
		log.Fatal("JSON error:", err)
	}

	// log.Println(bytes.NewBuffer(userJSON))

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://scratch.mit.edu/login/", bytes.NewBuffer(userJSON))
	if err != nil {
		log.Fatal("Error: ", err)
	}

	req.Header.Add("Referer", "https://scratch.mit.edu/")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("X-CSRFToken", "a")
	req.Header.Add("Cookie", "scratchcsrftoken=a;")

	resp, err := client.Do(req)
	
	if err != nil {
		log.Fatal("Error: ", err)
	}
	
	csrfToken := strings.Split(strings.TrimSpace(resp.Cookies()[1].Value), ";")[0]
	sessionId := strings.Split(strings.TrimSpace(resp.Cookies()[0].Value), ";")[0]
	
	defer resp.Body.Close()

	// fmt.Println(sessionId)

	// Cloud service
	// wss://clouddata.scratch.mit.edu

	// Handshake: { "method": "handshake", "user": "nikeedev", "project_id": project_id }
	// Message: { "method": "set", "name": "☁ message", "value": input_data.value }
	
	// project_id: 859836142
	//             ^^^^^^^^^ for "Websocket testing" project

	projectID := "1366075816"

	headers := http.Header{}
	headers.Add("X-CSRFToken", csrfToken)
	headers.Add("Cookie", "scratchcsrftoken="+csrfToken+"; scratchsessionsid="+sessionId+";")
	headers.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
	headers.Add("Origin", "https://scratch.mit.edu")
	
	conn, resp, err := websocket.DefaultDialer.Dial("wss://clouddata.scratch.mit.edu/", headers)

	if err != nil {
		log.Fatal(err)
	}
	
	// fmt.Println(resp.Status)

	defer conn.Close()

	// le handshake
	msg := CloudMessage{
		Method: "handshake",
		User: UserInfo.Username,
		ProjectID: projectID,
	}
	data, _ := json.Marshal(msg)
	data = append(data, '\n')

	// fmt.Println(string(data))

	err = conn.WriteMessage(
		websocket.TextMessage,
		data,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the server!")	

	defer conn.Close()

	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}

			// fmt.Println(string(message))

			for _, line := range strings.Split(string(message), "\n") {
				if strings.TrimSpace(line) == "" {
					continue
				}

				var msg CloudMessage

				err := json.Unmarshal([]byte(line), &msg)
				if err != nil {
					fmt.Println("Error unmarshaling JSON:", err)
					continue
				}
				
				if msg.Name == "☁ one" {
					fmt.Printf("[%s] %s\n",
						time.Now().Format("15:04:05"),
						Deblue(msg.Value),
					)
				}
			}
		}
	}()
	
	// always ready for your input
	for {
		var input string
		fmt.Scanln(&input)
		
		input = strings.TrimSpace(input)
		
		enblued := Enblue(input)

		fmt.Println(enblued)

		msg = CloudMessage{
			Method: "set",
			User: UserInfo.Username,
			ProjectID: projectID,
			Name: "☁ two",
			Value: enblued,
		}

		data, _ = json.Marshal(msg)
		data = append(data, '\n')
		err = conn.WriteMessage(
			websocket.TextMessage,
			data,
		)
		if err != nil {
			log.Println(err)
		}
	}
}


