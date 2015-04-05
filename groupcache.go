package main

import "github.com/golang/groupcache"

var filecache *groupcache.Group

//func init() {
//	peers := groupcache.NewHTTPPool(me)
// peers.Set(peerlist...)

// filecache = groupcache.NewGroup("filecache", int64(limit)<<20, groupcache.GetterFunc(
// 	func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
// 		contents, err := ioutil.ReadFile(key)
// 		dest.SetBytes(contents)
// 		return err
// 	}))
// }
