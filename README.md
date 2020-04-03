Krucible CLI
============

This is the official commandline interface for
[Krucible](https://usekrucible.com), the platform for creating ephemeral
Kubernetes clusters optimised for testing and development.

Installation
------------

Download the relevant binary for your OS from the [latest
release](https://github.com/Krucible/krucible-cli/releases/latest) and put it
on your PATH. Simple.

Usage
-----

### Creating a cluster
```
$ krucible create cluster --display-name $desiredClusterName
```
A cluster can be created with the above command. The `--display-name` flag is
mandatory.

```
$ krucible create cluster --display-name $desiredClusterName --configure-kubectl
```
Optionally, `krucible` can configure your `kubectl` context so that you can
immediately connect to your cluster.

### Retrieving a cluster
```
$ krucible get cluster $clusterUUID
```
A cluster can be retrieved with the above comand. The argument provided should
be the UUID of the cluster.

### Connecting to a cluster
```
$ krucible configure-kubectl $clusterUUID
```
`krucible` can configure your `kubectl` context so that you can immediately
connect to your cluster.
