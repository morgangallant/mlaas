# MLaaS

Introducing Memory-Leaks as a Service! MLaaS is a simple, cost-effective strategy to keep your engineers shipping code as fast as possible.

Here's how it works:

- Add the MLaaS library to your existing Go server.
- At startup, simply call `go mlaas.Start()`. This can be done inside of your `init()` or `main()` function.
- Every hour, MLaaS will allocate 100MiB of memory (and of course, not free it!).
- If your engineer's don't deploy new versions of your server code fast enough, then your servers will crash after a little while! How long? Who knows, it's a suprise!

Benefits:

- Establish a culture of "hardcore" engineering within your team.
- Increases motivation throughout your engineering team to ship code quickly and often.
- Increases server load, which means your engineers will have to write better, more efficient code to compensate!

Disclaimer:

This is a joke. Please don't actually use this library in production. It's not a good idea. This was inspired by a memory leak that we've had in our production servers at [Operand](https://operand.ai) for some time now. Our memory graphs may resemble the himalayas, but to be frank, we're not in a huge rush to fix it, as long as we deploy new code every day or two, everything works fine!
