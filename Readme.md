# Appix

[![Go](https://github.com/leanfra/appix/actions/workflows/go.yml/badge.svg)](https://github.com/leanfra/appix/actions/workflows/go.yml)

# Tag
Appix, Orchestrate, Optimize.

# Description:
Appix is an application-centric CMDB designed for orchestrating microservices. I developed it because I wanted to:

- help to manage servers with different features and used by different teams or products.
- help to manage the applications deployments on kubernetes. Simplify the server selection for products and developer teams.
- help to summarize the cost for products and developer teams from clouds bills, which is at the granularity of servers or other resources.

# System Benefits

- Configuration Center: Serves as the configuration center for deployment systems.
- Labeling: Utilizes standardized labels for identification.
- Grouping: It can standardize the grouping of applications and server groups.
- Cost Statistics: Serves as information center for cost statistics.


# System Features

1. Applications management.
2. developer Teams management.
2. Products management.
3. Hostgroups management. Auto match hostgroup for application by features.
4. Features management. Features are used to describe the application's requirements.
5. Tags management. Tags are used to label applications and hostgroups, which can be used to calculate the cost summary of Products and Teams.
6. Environments management.
7. Datacenters management.
8. Clusters management.
9. Users management.

# VI. Quick Start

## Configure

Modify configs/config.yaml for your environment.

## Run server

```
docker run -it \
    -v `pwd`/configs:/data/conf \
    -v `pwd`/database:/data/database \
    -p 9000:9000 \
    -p 8000:8000 \
    appix:v1
```

## Run cli

```
$ appix-cli -h
```
