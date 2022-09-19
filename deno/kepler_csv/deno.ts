import "https://deno.land/std/examples/welcome.ts";
import { denode } from "./arch.ts";
import "./deps.ts";

console.log(Deno.build);
console.log(Deno.env);

if (Deno.args[0] === "Hello") {
  console.log("😀");
} else {
  console.log("...");
}

console.table(Deno.metrics());
setTimeout(() => {
  console.log("async");
  console.table(Deno.metrics());
}, 1000);
denode("??");
