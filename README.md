## Crawler

### Greeting

Hello Monzo team! Nice to meet you :smile:

### Requirements 

go1.15

### Installation

```
$ go install
```

### Running

```
$ crawler -h
Usage:
  crawler [OPTIONS]

Application Options:
  -d, --domain=   domain to crawl
  -n, --workers=  number of concurrent workers (default: 1)
  -t, --timeout=  Maximum dial timeout duration in seconds (default: 15)
      --loglevel= log level {trace, debug, info, error, critical} (default: info)

Help Options:
  -h, --help      Show this help message

```

### Notes
 
The overall design of this could've been different if there were any constraints in the task. Please, consider this.
This crawler works with only one domain provided with a `--domain` flag.
Limited time = limited functionality :ok_hand:
