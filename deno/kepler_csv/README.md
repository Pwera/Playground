# Deno
First class TypeScript
ES Modules
Security first
Decentralized modules
Standard library
Built in tooling
Browser Compatible API
Single Executable
Async returns Promises
Opinionated Modules
 ``` javascript
deno run --allow-net deno.ts  Helloc // allow net permission
deno run --reload deno.ts // ignore cache and download deps
deno info  deno.ts     
deno cache --lock=lock.json --lock-write deps.ts // Locking deps
deno upgrade   // upgrade deno itself
deno bundle input.ts output.ts
``````

Deno requires files extension (.ts/ .js)
Ecmascript module can be used as Deno module eg. import * as _ from "https://deno.land/x/lodash@4.17.15-es/lodash.js"