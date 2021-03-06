# peer version

The `peer version` command displays the version information of the peer. It displays version, Go version, OS/architecture,
if experimental features are turned on, and chaincode information. For example:

```
 peer:
   Version: 1.1.0-beta-snapshot-a6c3447e
   Go version: go1.9.2
   OS/Arch: linux/amd64
   Experimental features: true
   Chaincode:
    Base Image Version: 0.4.5
    Base Docker Namespace: hyperledger
    Base Docker Label: org.hyperledger.udo
    Docker Namespace: hyperledger
```

## Syntax

```
Print current version of the udo peer server.

Usage:
  peer version [flags]

Flags:
  -h, --help   help for version

Global Flags:
      --logging-level string   Default logging level and overrides, see core.yaml for full syntax
```


<a rel="license" href="http://creativecommons.org/licenses/by/4.0/"><img alt="Creative Commons License" style="border-width:0" src="https://i.creativecommons.org/l/by/4.0/88x31.png" /></a><br />This work is licensed under a <a rel="license" href="http://creativecommons.org/licenses/by/4.0/">Creative Commons Attribution 4.0 International License</a>.
