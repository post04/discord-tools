# Description

Remove every friend a user has.

# Code 

```js
var findModule = (item) => Object.values(webpackJsonp.push([[],{['']:(_,e,r)=>{e.cache=r.c}},[['']]]).cache).find(m=>m.exports&&m.exports.default&&m.exports.default[item]!==void 0).exports.default;

var friends = findModule("getRelationships").getRelationships()

for(var k in friends){
findModule("removeRelationship").removeRelationship(k)
}
