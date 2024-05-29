# log-analyser

A terminal program that accepts lines of json on stdin and produces an analysis of those json objects.

```sh
cat file.txt | log-analyzer --interactive
```

By default the program will start off by rendering a TUI that displays the current internal model of the stream of logs. This is does while its reading from stdin and updating its internal model. At a fixed interval the TUI will be updated to reflect the internal model. Once all rows from stdin have been read, the 

Flags:

* `--interactive` disable the TUI
