import { assertEquals, assertNotEquals } from "./test_deps.ts";
import { log } from "./deps.ts";
import { filterPlantets } from "./models/planets.ts";
import * as Fae from "https://deno.land/x/fae/mod.ts";

const HABITABLE_PLANET = {
  koi_disposition: "CONFIRMED",
  koi_prad: "1",
  koi_srad: "1",
  koi_smass: "1",
};
const NOT_CONFIRMED = {
  koi_disposition: "FALSE_POSITIVE",
  koi_prad: "1",
  koi_srad: "1",
  koi_smass: "1",
};
const TOO_LARGE_RAD = {
  koi_disposition: "CONFIRMED",
  koi_prad: "1.5",
  koi_srad: "1",
  koi_smass: "1",
};

Deno.test("filter only habitable planets", () => {
  // given
  const data = [
    HABITABLE_PLANET,
    NOT_CONFIRMED,
    TOO_LARGE_RAD,
  ];

  // when
  const result = filterPlantets(data);

  // then
  assertEquals(result, [HABITABLE_PLANET]);
});

Deno.test("some test", () => {});

Deno.test({
  name: "ignored test",
  ignore: Deno.build.os === "windows",
  fn() {
  },
});

Deno.test("Fae.addIndex", () => {
  assertEquals(
    Fae.addIndex(Fae.map)(Fae.add)([10, 20, 30]),
    [10, 21, 32],
  );
});

Deno.test("Fae.map.filter.test", () => {
  const array = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
  const transformer = Fae.pipe(
    Fae.map(Fae.inc),
    Fae.take(3),
  );

  assertEquals(
    transformer(array),
    [2, 3, 4],
  );
});

Deno.test("Fae.range", () => {
  const limit = 4;
  Fae.range(0, limit).map((index) => {
    console.log(index);
  });

  assertEquals(
    Fae.addIndex(Fae.map)(Fae.add)([10, 20, 30]),
    [10, 21, 32],
  );
});
