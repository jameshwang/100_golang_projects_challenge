### Lessons Learned

1. Writer and Readers can be files or network connections.  Or it could be
std out/in.  What other things out there are writers and readers and how can
I used them in different scenarios since they're all the same thing with the
same interface?
1. Stack trace stuff can be hard. Use `runtime` lib to get the stack trace
and store them in a buffer and send them over the network.  Thinking how sentry
does this with ruby.
