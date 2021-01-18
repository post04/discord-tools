# Description

Sings (this song)[https://www.youtube.com/watch?v=h6DNdop6pD8&list=FLOBR-3lyUoKw89XUxV-TFZg] in the channel you're currently looking at.

# Code 

```js
var findModule = (item) => Object.values(webpackJsonp.push([[],{['']:(_,e,r)=>{e.cache=r.c}},[['']]]).cache).find(m=>m.exports&&m.exports.default&&m.exports.default[item]!==void 0).exports.default;

var song = `Okay, I know this is a really bad idea but
I'm already here so
Here we fuckin’ go
Rawr
​x3 nuzzles, pounces on you, uwu you so warm (Ooh)
Couldn't help but notice your bulge from across the floor
Nuzzles your necky wecky-tilde murr-tilde, hehe
Unzips your baggy ass pants, oof baby you so musky
Take me home, pet me, and make me yours and don't forget to stuff me
See me wag my widdle baby tail all for your bulgy-wulgy
Kissies and lickies your neck (Mmh)
I hope daddy likies
Nuzzles and wuzzles your chest (Yuh)
I be (Yeah) gettin’ thirsty
Hey, I got a little itch, you think you can help me?
Only seven inches long, uwu, please adopt me
Paws on your bulge as I lick my lips (UwU, punish me please)
'Bout to hit 'em with this furry shit (He don't see it comin')`.split("\n")
var i = 0
song.forEach(lyric => {
i++
setTimeout(() => {
findModule("sendMessage").sendMessage(findModule('getChannelId').getChannelId(), {content: lyric })
}, i * 1000)
})