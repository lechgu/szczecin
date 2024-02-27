## Szczecin

`szczecin` is an efficient and reliable tcp port scanner.

It detects tcp ports that are open on given target hosts out of specified list or range.

### Usage example

```
szczecin scan -t scanme.org -p 1-6535 --workers 1000 --progress
```

The `-t` flag is the host target to scan. More than one host target can be specified.

The `-p` flag indicates the port range and/or list of ports to scan, separated by commas.

The `--workers` flag specifies how many concurrent workers to have for the port scanning operation.

The `--progress` allows the scanner to display scan progress in the console. If omitted, the scan will be silent until it is complete.
