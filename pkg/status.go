package pkg

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	prettyTime "github.com/andanhm/go-prettytime"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/nlopes/slack"
)

type Status struct {
	ID      int
	URL     string
	Syncing *ethereum.SyncProgress
	Block   *types.Block
}

func NewStatus(id int, url string, syncing *ethereum.SyncProgress, block *types.Block) *Status {
	s := new(Status)
	s.ID = id
	s.URL = url
	s.Syncing = syncing
	s.Block = block
	return s
}

func (s *Status) Send(url string) error {
	// Pick pretext and color
	color := "#2ecc40"
	text := "Node is synchronized"
	if s.Syncing != nil {
		color = "#0074d9"
		text = fmt.Sprintf("Syncing block %d of %d", s.Syncing.CurrentBlock, s.Syncing.HighestBlock)
	}

	msg := &slack.WebhookMessage{
		Attachments: []slack.Attachment{
			slack.Attachment{
				Color: color,
				Fields: []slack.AttachmentField{
					slack.AttachmentField{
						Title: "Latest Block",
						Value: s.Block.Number().String(),
						Short: true,
					},
					slack.AttachmentField{
						Title: "Received At",
						Value: prettyTime.Format(time.Unix(s.Block.Time().Int64(), 0)),
						Short: true,
					},
				},
				Pretext:   "Ethereum Node Status",
				Title:     fmt.Sprintf("Node %d", s.ID+1),
				TitleLink: s.URL,
				Text:      text,
				Ts:        json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
			},
		},
	}

	if err := slack.PostWebhook(url, msg); err != nil {
		return err
	}
	return nil
}
