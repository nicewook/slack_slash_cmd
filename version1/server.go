// reference link: https://github.com/nlopes/slack/blob/master/examples/slash/slash.go
package main

import (
	"fmt"
	"net/http"
)

// SlashCommand is Slack slash command struct
type SlashCommand struct {
	Token          string `json:"token"`
	TeamID         string `json:"tead_id"`
	TeamDomain     string `json:"team_domain"`
	EnterpriseID   string `json:"enterprise_id,omitempty"`
	EnterPriseName string `json:"enterprise_name,omitempty"`
	ChannelID      string `json:"channel_id"`
	ChannelName    string `json:"channel_name"`
	UserID         string `json:"user_id"`
	UserName       string `json:"user_name"`
	Command        string `json:"command"`
	Text           string `json:"text"`
	ResponseURL    string `json:"response_url"`
	TirggerID      string `json:"trigger_id"`
}

func main() {

	http.HandleFunc("/slash", slashCommandHandler)
	fmt.Println("[INFO] Server listening")
	http.ListenAndServe("", nil)
}

func slashCommandHandler(w http.ResponseWriter, r *http.Request) {
	s, err := slashCommandParse(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: validateToken

	switch s.Command {
	case "/time":
		response := fmt.Sprintf("You requested for KST <-> PST/PDT for %#v", s)
		w.Write([]byte(response))

	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func slashCommandParse(r *http.Request) (s SlashCommand, err error) {
	if err = r.ParseForm(); err != nil {
		return s, err
	}
	s.Token = r.PostForm.Get("token")
	s.TeamID = r.PostForm.Get("team_id")
	s.TeamDomain = r.PostForm.Get("team_domain")
	s.EnterpriseID = r.PostForm.Get("enterprise_id")
	s.EnterPriseName = r.PostForm.Get("enterprise_name")
	s.ChannelID = r.PostForm.Get("channel_id")
	s.ChannelName = r.PostForm.Get("channel_name")
	s.UserID = r.PostForm.Get("user_id")
	s.UserName = r.PostForm.Get("user_name")
	s.Command = r.PostForm.Get("command")
	s.Text = r.PostForm.Get("text")
	s.ResponseURL = r.PostForm.Get("response_url")
	s.TirggerID = r.PostForm.Get("trigger_id")

	return s, nil

}
