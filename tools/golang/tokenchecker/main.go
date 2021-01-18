package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var (
	red    = color("\033[1;31m%s\033[0m")
	green  = color("\033[1;32m%s\033[0m")
	yellow = color("\033[1;33m%s\033[0m")
)

func color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

//https://stackoverflow.com/a/10510783
func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

var client = &http.Client{}
var link = "https://discordapp.com/api/v7/users/@me"

func checkToken(token string) ([]byte, bool) {
	var dead = false

	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		fmt.Println(red("ERROR:", err))
		return []byte("lol"), true
	}
	req.Header.Add("Authorization", token)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(red("ERROR:", err))
		return []byte("lol"), true
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return bodyBytes, true
	}
	return bodyBytes, dead
}

func getTokens() []string {
	var filename string
	fmt.Print(green("What file am I reading for tokens: "))
	fmt.Scanln(&filename)
	check := exists(filename)
	if !check {
		fmt.Println(red("The file " + filename + " doesn't exist!"))
		os.Exit(0)
	}
	input, err := ioutil.ReadFile("./" + filename)
	if err != nil {
		fmt.Println(red("Error reading "+filename+"\n", err))
		os.Exit(0)
	}
	return strings.Split(string(input), "\r\n")
}

func logTokens(text string) {
	var option string
	fmt.Print(green("Do you want to save the tokens to a file (y / yes): "))
	fmt.Scanln(&option)
	if strings.ToLower(option) == "y" || strings.ToLower(option) == "yes" {
		fmt.Print(green("Great! What file should I save to: "))
		var file string
		fmt.Scanln(&file)
		check := exists(file)
		if !check {
			fmt.Println(red("The file " + file + " doesn't exist!"))
			os.Exit(0)
		}
		f, err := os.Create(file)
		if err != nil {
			fmt.Println(red(fmt.Sprintf("Error reading %s\n%s", file, err)))
			f.Close()
			os.Exit(0)
		}
		f.WriteString(text)
		f.Close()
		fmt.Println(green(fmt.Sprintf("Tokens written to %s!", file)))
	}
}

type discordResponse struct {
	Verified bool   `json:"verified"`
	Username string `json:"username"`
	Discrim  string `json:"discriminator"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func logToConsole(data discordResponse, option bool, token string, dead bool) {
	if option {
		if dead {
			fmt.Println(red("Dead: ", token))
		} else {
			// var phone = "none"
			// if _, ok := data["Phone"]; ok {
			// 	phone = data.Phone
			// }
			// var email = "none"
			// if _, ok := data.Email; ok {
			// 	email = data.Email
			// }
			var phone = "none"
			if data.Phone != "" {
				phone = data.Phone
			}
			var email = "none"
			if data.Email != "" {
				email = data.Email
			}
			fmt.Println(green(fmt.Sprintf("Working: %s | %s#%s | %s | %s | %s", token, data.Username, data.Discrim, phone, email, fmt.Sprint(data.Verified))))
		}
	}
}

func main() {
	var working = []string{}
	tokens := getTokens()
	if len(tokens) < 1 {
		fmt.Println(red("I was unable to find any tokens!"))
		return
	}
	fmt.Println(green(fmt.Sprintf("Pulled %s tokens!", fmt.Sprint(len(tokens)))))
	var option string
	var logToConsoleOption bool
	fmt.Print(green("Should I log the resposnes to console (y / yes): "))
	fmt.Scanln(&option)
	option = strings.ToLower(option)
	logToConsoleOption = false
	if option == "y" || option == "yes" {
		logToConsoleOption = true
		fmt.Println(yellow("\n\n		Console logging layout"))
		fmt.Println(green("\nWorking: token | username#discrim | phone | email | verified"))
		fmt.Println(red("Dead: token\n"))
	}
	for _, token := range tokens {
		data, dead := checkToken(token)
		var resp discordResponse
		if dead {
			logToConsole(resp, logToConsoleOption, token, dead)
			continue
		}
		working = append(working, token)
		err := json.Unmarshal(data, &resp)
		if err != nil {
			fmt.Println(red("Error", err))
			continue
		}
		logToConsole(resp, logToConsoleOption, token, dead)
	}
	if len(working) > 0 {
		logTokens(strings.Join(working, "\n"))
	}

}
