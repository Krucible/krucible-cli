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
$ krucible create cluster --name my-cluster --duration 1
```
A cluster can be created with the above command. The `--name` and `--duration`
flags are mandatory.

The value for `--duration` represents the number of hours that the cluster
should run for. This can be any integer between 1 and 6 (inclusive) or
"permanent". A permanent cluster will only be deleted when deletion is
explicitly requested.

Optionally, `krucible` can configure your `kubectl` context so that you can
immediately connect to your cluster:
```
$ krucible create cluster --display-name my-cluster --duration --configure-kubectl
```

### Retrieving a cluster
```
$ krucible get cluster c-1234567
```
A cluster can be retrieved with the above comand. The argument provided should
be the ID of the cluster.

### Connecting to a cluster
```
$ krucible configure-kubectl c-1234567
```
`krucible` can configure your `kubectl` context so that you can immediately
connect to your cluster.

### Running kubectl
```
krucible kubectl --cluster c-1234567 -- get pods
```
`krucible` is also capable of running kubectl commands directly.

### Deleting a cluster
```
krucible delete cluster c-1234567
```
A cluster can be deleted with the above command. The argument provided should
be the ID of the cluster.
