# Description

Adds a random comma in a random place of your text before you send a message, this is 100% useless by every metric.

# Code

```js
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

window.Utils.Mods.hookMethod(window.Utils.Mods.findModule(m => m.hasOwnProperty('sendMessage')), "sendMessage", async (b) => {
    let message = b.methodArguments[1];


    //message = message
    let random = Math.ceil(Math.random() * (message.content.length - 1))
    let newcontent = message.content
    newcontent = newcontent.substr(0, random) + newcontent[random] + "'" + newcontent.substr(random + 1)
    message.content = newcontent

    return b.callOriginalMethod(b.methodArguments[0], message);
});
