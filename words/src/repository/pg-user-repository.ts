import { PubSub } from "@google-cloud/pubsub";
import { Pool } from "pg";
import { User } from "../model/user";
import { UserRepository } from "../service/interfaces";

export class PGUserRepository implements UserRepository {
  private client: Pool;

  constructor(client: Pool) {
    this.client = client;
  }

  create(u: User): Promise<User> {
    throw new Error("Method not implemented.");
  }
  updateUser(u: User): Promise<User> {
    throw new Error("Method not implemented.");
  }
}
