package pkg

import (
	"context"
	"errors"
	"net/url"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

type Notifier struct {
	Interval   uint
	WebhookURL string
	Nodes      []string
	Context    context.Context
	Clients    []*ethclient.Client
}

func NewNotifier(nodes []string, webhookURL string, interval uint) (*Notifier, error) {
	if len(nodes) == 0 {
		return nil, errors.New("Node list is empty")
	}

	if len(webhookURL) == 0 {
		return nil, errors.New("Webhook URL is empty")
	}

	if _, err := url.Parse(webhookURL); err != nil {
		return nil, err
	}

	n := new(Notifier)
	n.Context = context.Background()
	n.Nodes = nodes
	n.Interval = interval
	n.WebhookURL = webhookURL
	n.Clients = []*ethclient.Client{}

	for _, node := range nodes {
		client, err := ethclient.Dial(node)
		if err != nil {
			return nil, err
		}

		n.Clients = append(n.Clients, client)
	}
	return n, nil
}

func (n *Notifier) NodeStatus(id int) (*Status, error) {
	syncing, err := n.Clients[id].SyncProgress(n.Context)
	if err != nil {
		return nil, err
	}

	latest, err := n.Clients[id].BlockByNumber(n.Context, nil)
	if err != nil {
		return nil, err
	}

	return NewStatus(id, n.Nodes[id], syncing, latest), nil
}

func (n *Notifier) Run() {
	ticker := time.NewTicker(time.Duration(n.Interval) * time.Millisecond)
	quit := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			log.Info("Notifier tick")

			// Fetch
			statuses := []*Status{}
			for id, _ := range n.Nodes {
				log.Infof("Fetching node status %d", id+1)

				status, err := n.NodeStatus(id)
				if err != nil {
					log.Errorf("Error fetching node status %d: %s", id+1, err.Error())
				}

				statuses = append(statuses, status)
			}

			// Send
			for id, status := range statuses {
				log.Infof("Sending status %d", id+1)

				if err := status.Send(n.WebhookURL); err != nil {
					log.Errorf("Error sending nstatus %d: %s", id+1, err.Error())
				}
			}

		case <-quit:
			ticker.Stop()
			return
		}
	}
}
