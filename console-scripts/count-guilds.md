# Description

Counts the number of guilds the client is in.

# Code 

```js
var findModule = (item) => Object.values(webpackJsonp.push([[],{['']:(_,e,r)=>{e.cache=r.c}},[['']]]).cache).find(m=>m.exports&&m.exports.default&&m.exports.default[item]!==void 0).exports.default;

var guilds = findModule("getGuilds").getGuilds()
var i = 0
for(var k in guilds) i++
console.log(i)