## Crawler

Toy crawler for whatever you want to use it.

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

---

MIT License, 2020. Yarik Bratashchuk
