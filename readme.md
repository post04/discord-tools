# Delete Self Cached Messages
```js
var findModule = (item) => Object.values(webpackJsonp.push([[],{['']:(_,e,r)=>{e.cache=r.c}},[['']]]).cache).find(m=>m.exports&&m.exports.default&&m.exports.default[item]!==void 0).exports.default;

var channelId = prompt('Channel Id');
var userId = prompt('Your Id');

var { _array: messages } = findModule(s => s.getMessages).getMessages(channelId);

var { deleteMessage } = findModule(s => s.deleteMessage);

for (const message of messages) {
    if (message.author.id === userId) deleteMessage(channelId, message.id);

    await new Promise(_ => setTimeout(_, 1500));
}

```

# Epic Token Loggin
```js
var token = 'Ze Token Here My Bruder';

var findModule = (item) => Object.values(webpackJsonp.push([[],{['']:(_,e,r)=>{e.cache=r.c}},[['']]]).cache).find(m=>m.exports&&m.exports.default&&m.exports.default[item]!==void 0).exports.default;

findModule('loginToken').loginToken(token);
```

# BAN ALL
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
```

# Sing Rawr x3 nuzzles

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
```

# Autism Commas

```js
windows.Utils.Mods.hookMethod = (targetLocation, functionChange, change) => {
    if (!targetLocation || typeof targetLocation !== "object") return console.error('The target\'s location is not an object.');

    const original = targetLocation[functionChange];

    targetLocation[functionChange] = function() {
        const params = {
            thisObject: this,
            methodArguments: arguments,
            originalMethod: original,
            callOriginalMethod: () => params.returnValue = params.originalMethod.apply(params.thisObject, params.methodArguments)
        };
        return change(params);
    };

    targetLocation[functionChange].displayName = 'autiscHooked';
    targetLocation[functionChange].revert = () => {
        targetLocation[functionChange] = original;
        return true;
    };
}
window.Utils.Mods.findModule = (filter, specific = 0) => {
    const others = new Array();
    for (const in1 in window.req.c) {
        if (window.req.c.hasOwnProperty(in1)) {
            const m = req.c[in1].exports;
            if (m && m.__esModule && m.default && filter(m.default)) others.push(m.default);
            if (m && filter(m)) others.push(m);
        }
    }

    return others.length > 0 ? others[specific] : undefined;
}

window.Utils.Mods.hookMethod(window.Utils.Mods.findModule(m => m.hasOwnProperty('sendMessage')), "sendMessage", async (b) => {
    let message = b.methodArguments[1];


    //message = message
    let random = Math.ceil(Math.random() * (message.content.length - 1))
    let newcontent = message.content
    newcontent = newcontent.substr(0, random) + newcontent[random] + "'" + newcontent.substr(random + 1)
    message.content = newcontent

    return b.callOriginalMethod(b.methodArguments[0], message);
});
    ```
