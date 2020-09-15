import dotenv from "dotenv";
import path from "path";
import crypto from "crypto";
import { initDS, DataSources } from "./data";
import { serviceContainer } from "./injection";
import createApp from "./app";
import { UserUpdatesListener } from "./events/user-updates-listener";

const startup = async () => {
  /*
   * Load environment vars
   */
  const result = dotenv.config({
    path: path.resolve(process.cwd(), ".env.dev"),
  });

  if (result.error) {
    console.error(
      `Unable to load environment variables. Reason:\n${result.error}`
    );
    process.exit();
  }

  console.info("Successfully loaded environment variables!");

  /*
   * Initialize data sources (just postgres so far)
   */

  let ds: DataSources;

  try {
    // note that we still use port 5432 since we're in the world of containers
    ds = await initDS();
  } catch (err) {
    console.error(err);
    process.exit();
  }

  console.info("Successfully initialized data sources!");
  // console.info(ds);

  /*
   * Inject concrete repository implementations into services
   */
  serviceContainer.init(ds);

  console.info("Service container initialized");

  const app = createApp();

  app.listen(process.env.PORT, () => {
    console.log(`Listening on port ${process.env.PORT}`);
  });

  // create listener for pubsub
  const listener = new UserUpdatesListener({
    userService: serviceContainer.services.userService,
    pubSub: ds.pubSubClient,
  });

  await listener.init("word-app", {
    ackDeadline: 30,
  });
  listener.listen();

  process.on("SIGINT", async () => await shutdown());
  process.on("SIGTERM", async () => await shutdown());

  const shutdown = async () => {
    ds.pubSubClient.close();
    await listener.subscription?.delete();
  };
};

startup();
