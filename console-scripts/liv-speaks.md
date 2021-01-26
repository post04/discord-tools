# Description

This is how every female on the planet types.

# Code

```js
window.Utils = {Mods:{}}
window.Utils.Mods.hookMethod = (targetLocation, functionChange, change) => {
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
function genNum(toCheck, limit){
    var rand = Math.ceil(Math.random() * limit)
    if(rand == toCheck){
        if(limit > toCheck) return toCheck++
        return toCheck--
    }
    return rand
}

function shuffelWord (word){
    var shuffledWord = '';
    word = word.split('');
    var first = genNum(0, word.length-1)
    var second = genNum(first, word.length-1)
    var firstletter = word[first]
    var secondletter = word[second]
    var i = 0
    word.forEach(k => {
            if(i == first){
                k = secondletter
            }else if(i == second){
                k = firstletter
            }
            shuffledWord += k
        i++
    })

    return shuffledWord;
}
window.Utils.Mods.hookMethod(window.Utils.Mods.findModule(m => m.hasOwnProperty('sendMessage')), "sendMessage", async (b) => {
    let message = b.methodArguments[1];


    let newcontent = message.content
    var place = 0
    message.content.split(" ").forEach(word => {
    newcontent = newcontent.replace(newcontent.split(" ")[place], shuffelWord(word))
place++
    })
    message.content = newcontent
    return b.callOriginalMethod(b.methodArguments[0], message);
});
