# Description 

Closes all dms on users account.

# Code 

```js
var findModule = (item) => Object.values(webpackJsonp.push([[],{['']:(_,e,r)=>{e.cache=r.c}},[['']]]).cache).find(m=>m.exports&&m.exports.default&&m.exports.default[item]!==void 0).exports.default
var dms = findModule("getPrivateChannels").getPrivateChannels()

for(var k in dms){
findModule("closePrivateChannel").closePrivateChannel(k)
}
