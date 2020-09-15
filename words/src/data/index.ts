import { Pool } from "pg";
import { PubSub } from "@google-cloud/pubsub";

export interface DataSources {
  db: Pool;
  pubSubClient: PubSub;
}

export const initDS = async (): Promise<DataSources> => {
  const db = new Pool();

  // test connection
  try {
    const client = await db.connect();

    client.release();
  } catch (err) {
    throw new Error(`Unable to connect to postgres. Reason: ${err}`);
  }

  // init pubsub client
  const pubSubClient = new PubSub();

  return {
    db,
    pubSubClient,
  };
};
