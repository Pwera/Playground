import { flatMap, log } from "../deps.ts";

export interface Launch {
  flightNumber: number;
  mission: string;
  rocket: string;
  customers: Array<string>;
  launchDate: number;
  upcoming: boolean;
  success?: boolean;
  target?: string;
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
    const customers = flatMap(payloads, (payload: any) => {
      return payload["customers"];
    });
    const flightData = {
      flightNumber: launch["flight_number"],
      mission: launch["mission_name"],
      rocket: launch["rocket"]["rocket_name"],
      launchDate: launch["launch_date_unix"],
      upcoming: launch["upcoming"],
      success: launch["launch_success"],
      customers: customers,
    };
    launches.set(flightData.flightNumber, flightData);
    // log.info(JSON.stringify(flightData));
  }
}

await downloadLaunchData();
log.info(`Downloaded data for ${launches.size} SpaceX launches.`);

export function getAllLaunches() {
  log.info(`size of launches in getAllLaunches ${launches.size}`);
  return Array.from(launches.values());
}
export function get(id: number) {
  if (launches.has(id)) {
    return launches.get(id);
  }
  return null;
}

export function add(data: Launch) {
  Object.assign(data, {
    upcoming: true,
    customers: ["Some Customer"],
  });
  launches.set(data.flightNumber, data);
}

export function remove(id: number) {
  launches.delete(id);
  const aborted = launches.get(id);
  if (aborted) {
    aborted.upcoming = false;
    aborted.success = false;
  }
  return aborted;
}
