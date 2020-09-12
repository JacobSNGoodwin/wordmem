import dotenv from "dotenv";
import path from "path";
import app from "./app";
import { init } from "./data_sources";
import { PoolConfig } from "pg";

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

  try {
    // note that we still use port 5432 since we're in the world of containers
    await init();
  } catch (err) {
    console.error(err);
    process.exit();
  }

  console.info("Successfully initialized data sources!");

  app.listen(process.env.PORT, () => {
    console.log(`Listening on port ${process.env.PORT}`);
  });
};

startup();
