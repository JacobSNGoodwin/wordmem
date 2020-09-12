import { Pool, PoolConfig } from "pg";

interface DataSources {
  db: Pool;
}

export const init = async (): Promise<DataSources> => {
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
