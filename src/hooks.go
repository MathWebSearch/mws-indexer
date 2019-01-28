package src

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

// PostUpdateHooks runs post-update hooks
func PostUpdateHooks() {
	if dockerLabel != "" {
		restartDockerWithLabel(dockerLabel)
	}
}

func restartDockerWithLabel(label string) {
	fmt.Printf("Restarting docker container(s) with label %q. \n", label)

	// create a new docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.26"))
	if err != nil {
		panic(err)
	}

	// filter it by label
	filters := filters.NewArgs()
	filters.Add("label", label)

	// list all containers with the label
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		Filters: filters,
	})

	if err != nil {
		panic(err)
	}

	minute := 15 * time.Second

	// restart all the docker containers
	for _, container := range containers {
		fmt.Printf("Restarting container %q\n", container.Names[0])

		err := cli.ContainerRestart(context.Background(), container.ID, &minute)
		if err != nil {
			panic(err)
		}
	}
}
