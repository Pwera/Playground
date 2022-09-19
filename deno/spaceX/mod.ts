import * as log from "https://deno.land/std/log/mod.ts";
import * as _ from "https://deno.land/x/lodash@4.17.15-es/lodash.js";

// log setup
await log.setup({
  handlers: {
    console: new log.handlers.ConsoleHandler("DEBUG"),
  },
  loggers: {
    default: {
      level: "DEBUG",
      handlers: ["console", "file"],
    },
  },
});

interface Launch {
  flightNumber: number;
  mission: string;
  rocket: string;
  customers: Array<string>;
}
const launches = new Map<number, Launch>();

async function downloadLaunchData() {
  log.debug("debug log");
  log.warning("warning log");
  const response = await fetch("http://api.spacexdata.com/v3/launches", {
    method: "GET",
  });

  if (!response.ok) {
    log.warning("Problem downloading launch data.");
    throw new Error("Launch data download failed.");
  }

  const launchData = await response.json();
  if (false) {
    console.log(launchData);
  }
  for (const launch of launchData) {
    const payloads = launch["rocket"]["second_stage"]["payloads"];
    const customers = _.flatMap(payloads, (payload: any) => {
      return payload["customers"];
    });
    const flightData = {
      flightNumber: launch["flight_number"],
      mission: launch["mission_name"],
      rocket: launch["rocket"]["rocket_name"],
      customers: customers,
    };
    launches.set(flightData.flightNumber, flightData);
    log.info(JSON.stringify(flightData));
    log.info("ok");
  }
}

if (import.meta.main) {
  await downloadLaunchData();
  log.info(`Downloaded data for ${launches.size} SpaceX launches.`);
  log.info(import.meta);
}
