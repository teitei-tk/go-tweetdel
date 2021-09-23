package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
)

type Tweets []struct {
	Tweet `json:"tweet"`
}

type Tweet struct {
	Retweeted bool   `json:"retweeted"`
	Source    string `json:"source"`
	Entities  struct {
		Hashtags     []interface{} `json:"hashtags"`
		Symbols      []interface{} `json:"symbols"`
		UserMentions []interface{} `json:"user_mentions"`
		Urls         []interface{} `json:"urls"`
	} `json:"entities"`
	DisplayTextRange []string    `json:"display_text_range"`
	FavoriteCount    string      `json:"favorite_count"`
	IDStr            string      `json:"id_str"`
	Truncated        bool        `json:"truncated"`
	RetweetCount     string      `json:"retweet_count"`
	ID               string      `json:"id"`
	CreatedAt        TwCreatedAt `json:"created_at"`
	Favorited        bool        `json:"favorited"`
	FullText         string      `json:"full_text"`
	Lang             string      `json:"lang"`
}

type TwCreatedAt struct {
	time.Time
}

func (m *TwCreatedAt) UnmarshalJSON(data []byte) error {
	t, err := time.Parse(fmt.Sprintf(`"%s"`, time.RubyDate), string(data))
	*m = TwCreatedAt{Time: t}
	return err
}

func ReadTweetsJSON(path string) (*Tweets, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "cannot read tweets.js")
	}

	twByte := bytes.Replace(raw, []byte("window.YTD.tweet.part0 = "), []byte(""), -1)
	tweets := &Tweets{}
	if err := json.Unmarshal(twByte, tweets); err != nil {
		return nil, err
	}

	return tweets, nil
}
