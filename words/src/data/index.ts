import { Pool } from "pg";

export interface DataSources {
  db: Pool;
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

  return {
    db,
  };
};
