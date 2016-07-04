= pongoc

== About

Console tool for compile template from stdin to stdout with env variables as context.

== Install

```
$ go go github.com/miolini/pongoc
```

== Run

```
$ ES_CLUSTER_NAME=skynet_logs pongoc < template.txt > result.txt
```
