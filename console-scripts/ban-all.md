# Description 

This scripts bans everyone that your client has cached in a guild, it's pretty fast but you might have to run it a few times depending on the size of the guild, it fetches members by the cache so make sure you load the member bar on the right side. **DOESN'T BAN OFFLINE MEMBERS**

# Code
 
```js
var findModule = (item) => Object.values(webpackJsonp.push([[],{['']:(_,e,r)=>{e.cache=r.c}},[['']]]).cache).find(m=>m.exports&&m.exports.default&&m.exports.default[item]!==void 0).exports.default;

var serverId = prompt('Server Id');
var members = findModule('getMembers').getMembers(serverId);
var { banUser } = findModule('banUser');
var count = 0;

while (true) {
  if ((count % 40) === 0) await new Promise(_ => setTimeout(_, 2000));
  banUser(serverId, members[count].userId);
  count++;
}