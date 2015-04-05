Kasoku (加速)
=====

Implements a groupcache backed FUSE layer to provide AP caching capabilities to all kinds of fun stuff.

This project was previously called 'hayai' (速い) but I think kasoku is a better name. Awesome! That's the hard stuff out of the way.

Usage
======

```
Usage of /home/aj/go/bin/kasoku:
  -cacheLimit=1048576000000: Default maximum key size limit, in bytes
  -groupName="kasoku": The name of the groupcache cluster
  -me="127.0.0.1:710": The bind address for the groupcache endpoint
  -mountpoint="/kasoku": Where to mount the FUSE groupcache storage
  -peerlist="": The remote peers to attempt to form Voltron and activate incredible speed power
  -target="/var/cache/cookbooks": The source from which groupcache fetches from
```
