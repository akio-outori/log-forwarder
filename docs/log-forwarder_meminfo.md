## log-forwarder meminfo

Return memory metrics as a json string

### Synopsis

Command should get linux / unix memory and return it in a JSON formatted string that can be consumed by fluentd

```
log-forwarder meminfo [flags]
```

### Options

```
  -e, --ex     whether to collect extended memory information
  -h, --help   help for meminfo
  -s, --swap   whether to collect swap information
  -m, --vm     whether to collect memory information
```

### SEE ALSO

* [log-forwarder](log-forwarder.md)	 - Test log-forwarder for fluentd in go

###### Auto generated by spf13/cobra on 17-Aug-2020
