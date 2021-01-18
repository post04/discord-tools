package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
)

var used = []string{}
var (
	giftRegex = regexp.MustCompile("(discord.com/gifts/|discordapp.com/gifts/|discord.gift/)([a-zA-Z0-9]+)")
)

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

var (
	token, err = ioutil.ReadFile("token.txt")
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content != "" && giftRegex.Match([]byte(m.Content)) {
		gifts := giftRegex.FindStringSubmatch(m.Content)
		if len(gifts) < 2 {
			return
		}
		if len(gifts[2]) != 16 {
			color.Red("Fake code %s sent from %s!", gifts[2], m.Author.String())
			return
		}
		usedalreadty := contains(used, gifts[2])
		if usedalreadty {
			color.Red("Code %s wont retry, sent from %s", gifts[2], m.Author.String())
			return
		}
		start := makeTimestamp()
		client := &http.Client{}
		var link = "https://discordapp.com/api/v6/entitlements/gift-codes/" + gifts[2] + "/redeem"
		var body = []byte(`{"channel_id":` + m.ChannelID + "}")
		req, err := http.NewRequest("POST", link, bytes.NewBuffer(body))
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("authorization", string(token))
		res, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		bodyBytes, err := ioutil.ReadAll(res.Body)
		bodyString := string(bodyBytes)
		end := makeTimestamp()
		used = append(used, gifts[2])
		if strings.Contains(bodyString, "This gift has been redeemed already.") {
			color.Red("code %s already used, sent by %s, took %d ms", gifts[2], m.Author.String(), end-start)
		} else if strings.Contains(bodyString, "nitro") {
			color.Green("code %s used, sent by %s, took %d ms", gifts[2], m.Author.String(), end-start)
		} else if strings.Contains(bodyString, "Unknown Gift Code") {
			color.Red("code %s isn't nitro, sent by %s, took %d ms", gifts[2], m.Author.String(), end-start)
		} else {
			color.Yellow("%s has not attributes, sent by %s, took %d ms", gifts[2], m.Author.String(), end-start)
		}
	}
}

func main() {
	dg, err := discordgo.New(string(token))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	dg.AddHandler(messageCreate)
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	color.Cyan("Sniping on " + strconv.Itoa(len(dg.State.Guilds)) + " Servers. \n")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	_ = dg.Close()
}
