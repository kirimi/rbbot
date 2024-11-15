package commander

import (
	"context"
	"encoding/json"
	"fmt"
	containertypes "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (c *Commander) ContainersList(message *tgbotapi.Message) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	defer func(cli *client.Client) {
		_ = cli.Close()
	}(cli)

	containers, err := cli.ContainerList(ctx, containertypes.ListOptions{})
	if err != nil {
		return err
	}

	output := "Running Containers\n\n"

	for index, container := range containers {
		var memUsage string
		if stats, err := getContainerStats(ctx, cli, container.ID); err == nil {
			memUsage = bytesToHuman(stats.MemoryStats.Usage)
		} else {
			memUsage = "unknown"
		}

		output += strconv.Itoa(index+1) + ". " + container.Image + " " + container.Status + " " + memUsage + "\n"
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, output)
	_, err = c.bot.Send(msg)

	return err
}

func getContainerStats(ctx context.Context, cli *client.Client, containerId string) (*containertypes.StatsResponse, error) {
	resp, err := cli.ContainerStatsOneShot(ctx, containerId)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var v *containertypes.StatsResponse
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return nil, err
	}

	return v, nil
}

func bytesToHuman(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}
