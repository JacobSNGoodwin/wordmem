import { PubSub } from "@google-cloud/pubsub";
import { Pool } from "pg";
import { InternalError } from "../errors/internal-error";
import { User } from "../model/user";
import { UserRepository } from "../service/interfaces";

export class PGUserRepository implements UserRepository {
  private client: Pool;

  constructor(client: Pool) {
    this.client = client;
  }

  async upsert(u: User): Promise<User> {
    const text = `
        INSERT INTO users (id, email) 
        VALUES ($1, $2) 
        ON CONFLICT (id)
        DO
        UPDATE SET email=EXCLUDED.email
        RETURNING *;
      `;
    const values = [u.id, u.email];

    try {
      const queryRes = await this.client.query({
        text,
        values,
      });

      const createdUser = queryRes.rows[0];

      return userFromData(createdUser);
    } catch (e) {
      console.debug("Error inserting word into database: ", e);
      throw new InternalError();
    }
  }
}

const userFromData = (dataObj: any): User => ({
  id: dataObj.id,
  email: dataObj.email,
});
