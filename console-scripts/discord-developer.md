# Description 

Makes discord client think you're a developer giving you beta features early.

# Code

```js
var FindModule = (item) => Object.values(webpackJsonp.push([[],{['']:(_,e,r)=>{e.cache=r.c}},[['']]]).cache).find(m=>m.exports&&m.exports.default&&m.exports.default[item]!==void 0).exports.default;

Object.defineProperty(FindModule('getExperimentDescriptor').__proto__, 'isDeveloper', { get: () => true });
