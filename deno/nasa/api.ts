import { Router } from "./deps.ts";
import * as planets from "./models/planets.ts";
import * as launches from "./models/launches.ts";

const router = new Router();

router.get("/", (ctx) => {
  ctx.response.body = "NASA";
});

router.get("/planets", (ctx) => {
  const response = planets.getAllPlanets();
  ctx.response.body = response;
  //log.info(`/planets returned ${response.length} rows`);
});

router.get("/launches", (ctx) => {
  const response = launches.getAllLaunches();
  ctx.response.body = response;
  //log.info(`/launches returned ${response.length} rows`);
});

router.delete("/launches/:id", (ctx) => {
  if (ctx.params?.id) {
    const result = launches.remove(Number(ctx.params.id));
    ctx.response.body = { success: result };
  }
});

router.get("/launches/:id", (ctx) => {
  if (ctx.params?.id) {
    const launchList = launches.get(Number(ctx.params.id));
    if (launchList) {
      ctx.response.body = launchList;
    } else {
      ctx.throw(400, "Launch with that ID doesn't exist");
    }
  }
});

router.post("/launches", async (ctx) => {
  const result = await ctx.request.body();
  const data = await result.value;
  launches.add(data);
  ctx.response.body = { success: true };
  ctx.response.status = 201;
});

router.get("/error", (ctx) => {
  ctx.throw(
    501,
    "Some server error, this log message is not visible on client side!",
  );
});

export default router;
