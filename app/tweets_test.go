package app_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/teitei-tk/go-tweetdel/app"
)

func TestReadTweetsJSON(t *testing.T) {
	assert := assert.New(t)

	sampleJSON := `[
    {
      "tweet" : {
        "retweeted" : false,
        "source" : "<a href=\"http://twitter.com\" rel=\"nofollow\">Twitter Web Client</a>",
        "entities" : {
          "hashtags" : [],
          "symbols" : [],
          "user_mentions" : [],
          "urls" : []
        },
        "display_text_range" : [ "0", "4" ],
        "favorite_count" : "0",
        "id_str" : "1",
        "truncated" : false,
        "retweet_count" : "0",
        "id" : "1",
        "created_at" : "Mon Jan 02 15:04:05 -0700 2006",
        "favorited" : false,
        "full_text" : "test",
        "lang" : "en"
      }
    }
	]`

	twTime, err := time.Parse(time.RubyDate, "Mon Jan 02 15:04:05 -0700 2006")
	assert.NoError(err)

	cAt := app.TwCreatedAt{Time: twTime}

	tests := []struct {
		desc       string
		jsonByte   []byte
		wantTweets *app.Tweets
		creteJSON  bool
		wantError  bool
	}{
		{
			desc:      "file not found",
			creteJSON: false,
			wantError: true,
		},
		{
			desc:      "syntax error",
			jsonByte:  []byte("hogehoge"),
			creteJSON: true,
			wantError: true,
		},
		{
			desc:      "marshal error",
			jsonByte:  []byte("{'id': 1}"),
			creteJSON: true,
			wantError: true,
		},
		{
			desc:       "data not found",
			jsonByte:   []byte("[]"),
			wantTweets: &app.Tweets{},
			creteJSON:  true,
			wantError:  false,
		},
		{
			desc:     "passed",
			jsonByte: []byte(sampleJSON),
			wantTweets: &app.Tweets{
				{
					Tweet: app.Tweet{
						Retweeted: false,
						Source:    "<a href=\"http://twitter.com\" rel=\"nofollow\">Twitter Web Client</a>",
						Entities: struct {
							Hashtags     []interface{} "json:\"hashtags\""
							Symbols      []interface{} "json:\"symbols\""
							UserMentions []interface{} "json:\"user_mentions\""
							Urls         []interface{} "json:\"urls\""
						}{
							Hashtags:     []interface{}{},
							Symbols:      []interface{}{},
							UserMentions: []interface{}{},
							Urls:         []interface{}{},
						},
						DisplayTextRange: []string{"0", "4"},
						FavoriteCount:    "0",
						IDStr:            "1",
						Truncated:        false,
						RetweetCount:     "0",
						ID:               "1",
						CreatedAt:        cAt,
						Favorited:        false,
						FullText:         "test",
						Lang:             "en",
					},
				},
			},
			creteJSON: true,
			wantError: false,
		},
		{
			desc:     "raw js",
			jsonByte: []byte(strings.Join([]string{"window.YTD.tweet.part0 = ", sampleJSON}, "")),
			wantTweets: &app.Tweets{
				{
					Tweet: app.Tweet{
						Retweeted: false,
						Source:    "<a href=\"http://twitter.com\" rel=\"nofollow\">Twitter Web Client</a>",
						Entities: struct {
							Hashtags     []interface{} "json:\"hashtags\""
							Symbols      []interface{} "json:\"symbols\""
							UserMentions []interface{} "json:\"user_mentions\""
							Urls         []interface{} "json:\"urls\""
						}{
							Hashtags:     []interface{}{},
							Symbols:      []interface{}{},
							UserMentions: []interface{}{},
							Urls:         []interface{}{},
						},
						DisplayTextRange: []string{"0", "4"},
						FavoriteCount:    "0",
						IDStr:            "1",
						Truncated:        false,
						RetweetCount:     "0",
						ID:               "1",
						CreatedAt:        cAt,
						Favorited:        false,
						FullText:         "test",
						Lang:             "en",
					},
				},
			},
			creteJSON: true,
			wantError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			jsonPath := filepath.Join(os.TempDir(), time.Now().Format("20060102150405"), "hoge.json")

			if test.creteJSON {
				err := os.MkdirAll(filepath.Dir(jsonPath), 0777)
				if err != nil {
					assert.Error(err)
					return
				}

				defer os.RemoveAll(filepath.Dir(jsonPath))

				err = os.WriteFile(jsonPath, test.jsonByte, 0644)
				if err != nil {
					assert.Error(err)
					return
				}
			}

			tw, err := app.ReadTweetsJSON(jsonPath)
			if test.wantError {
				assert.Error(err)
				assert.Nil(tw)
				return
			}

			assert.NotNil(tw)
			assert.NoError(err)

			tweets := *tw
			wantTweets := *test.wantTweets
			assert.Equal(len(wantTweets), len(tweets))

			for i, tweet := range tweets {
				wantTweet := wantTweets[i]

				assert.Equal(wantTweet.Retweeted, tweet.Retweeted)
				assert.Equal(wantTweet.Source, tweet.Source)
				assert.Equal(wantTweet.Entities, tweet.Entities)
				assert.Equal(wantTweet.DisplayTextRange, tweet.DisplayTextRange)
				assert.Equal(wantTweet.FavoriteCount, tweet.FavoriteCount)
				assert.Equal(wantTweet.IDStr, tweet.IDStr)
				assert.Equal(wantTweet.Truncated, tweet.Truncated)
				assert.Equal(wantTweet.RetweetCount, tweet.RetweetCount)
				assert.Equal(wantTweet.ID, tweet.ID)
				assert.Equal(wantTweet.CreatedAt, tweet.CreatedAt)
				assert.Equal(wantTweet.Favorited, tweet.Favorited)
				assert.Equal(wantTweet.FullText, tweet.FullText)
				assert.Equal(wantTweet.Lang, tweet.Lang)
			}
		})
	}
}
