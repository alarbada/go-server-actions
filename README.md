### Experimental implementation of nextjs server actions in go

Stack:

- My fork of gomponents as a template engine
- gin router for grouping actions
- That's it!


### Why?

Whenever using htmx with go, I always felt like repeating the url in both the view and the server definition was not ideal. It felt fragile. This package is an experiment of bringing almost the same benefits of react server actions to golang + gin + htmx + gomponents.

The implementation is so simple that you could really take this approach to other frameworks / languages. This is just an implementation.
