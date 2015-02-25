# RabbitMQ Cluster Migration Tool

This repository contains the RabbitMQ Cluster Migration Tool to be used in conjunction with [`cf-rabbitmq-release`](https://github.com/pivotal-cf/cf-rabbitmq-release).

## Installation

```
go get -u github.com/pivotal-cf/rabbitmq-cluster-migration-tool
```

## Usage

To use the tool, you need to pass in the path to the directory of a migration directory that contains your old `erl_inetrc` file, as well as the path to a file which contains an array of mappings of old node names to new node names.
```
rabbitmq-cluster-migration-tool <migration_dir_path> <cluster_state_file_path>
```

