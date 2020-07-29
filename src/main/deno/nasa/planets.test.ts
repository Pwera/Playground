import { assertEquals, assertNotEquals } from "./test_deps.ts";
import { log } from "./deps.ts";
import { filterPlantets } from "./models/planets.ts";

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
