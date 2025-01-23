# OpsPillar

[![Go](https://github.com/leanfra/opspillar/actions/workflows/go.yml/badge.svg)](https://github.com/leanfra/opspillar/actions/workflows/go.yml)

# Tag
opspillar, Orchestrate, Optimize.

# Description:
OpsPillar is an application-centric CMDB designed for orchestrating microservices. I developed it because I wanted to:

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

# Quick Start

## Concepts

- Application and Hostgroup are the forcus items managed by OpsPillar.
- Features are used to describe the application's requirements and hostgroups can provide.
- Application's matched Hostgroups must have the same Team and Product, and Application's Features must be subset of Hostgroup's Features. 
- Teams, Products, Environments, Datacenters, Clusters are used to describe the Application and Hostgroups.
- Teams and Products are special tags which must be set because they are used everywhere, such as, to calculate the cost summary.
- Tags are used to label applications and hostgroups.

## Configure

Modify configs/config.yaml for your environment.

## Run server

```
docker run -it \
    -v `pwd`/configs:/data/conf \
    -v `pwd`/database:/data/database \
    -p 9000:9000 \
    -p 8000:8000 \
    opspillar:v1
```

## Run cli

```
$ opspillar-cli -h
```

## examples

### Application

- get all applications

```
❯ ./opspillar-cli get apps
+----+------+----------------------+-------+----------+---------+-------+-----------+----------+------------+
| ID | Name |     Description      | Owner | Stateful | Product | Team  | Features  |   Tags   | Hostgroups |
+----+------+----------------------+-------+----------+---------+-------+-----------+----------+------------+
|  2 | app1 | app description gag1 | gag1  | false    | meta    | infra | cpu:intel | sla:9999 | meta-intel |
+----+------+----------------------+-------+----------+---------+-------+-----------+----------+------------+

```

### Hostgroup

- get all hostgroups

```
❯ ./opspillar-cli get hg
+----+------------+---------------------+---------+--------------+-----+---------+-------+-----------+------+---------------+------------+
| ID |    Name    |     Description     | Cluster |  Datacenter  | Env | Product | Team  | Features  | Tags | ShareProducts | ShareTeams |
+----+------------+---------------------+---------+--------------+-----+---------+-------+-----------+------+---------------+------------+
|  4 | meta-intel | for meta with intel | k8s-0   | aws-us-west2 | prd | meta    | infra | cpu:intel |      | meta, myin    |            |
+----+------------+---------------------+---------+--------------+-----+---------+-------+-----------+------+---------------+------------+

```

# Development

## Scaffold

- [kratos](https://go-kratos.dev/)
- [gorm](https://gorm.io/)
- [corbra](https://github.com/spf13/cobra)

## Build

Refer to Makefile for details.

- build server

```
make build
```
- build cli

```
make cli
```
