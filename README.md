Witter - Minimalistic Microblogging-server
==============================================

Witter is a minimalistic microblogging-server. Right now all messages (“weets”)
are saved only in-memory and there is only one user “axel” and all in all there
are relitely few things working - it's mostly a minimally extended version of
[Gary Burd's Websocket-Example](http://gary.beagledreams.com/page/go-websocket-chat.html).

All similarities with existing (possibly even trademarked)
microblogging-services are purely coincidental.

Installation
============

`go get github.com/FachschaftMathPhys/witter` is theoretically working, but not
practically, because static resources are served from the current working
directory. Sometime maybe this will not just be a crude hack, maybe it will be
possible to install it like that then.
