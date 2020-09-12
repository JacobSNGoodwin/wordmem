import { WordRepository } from "../service/interfaces";
import { Word } from "../model/word";
import { CustomError } from "../errors/custom-error";
import { Pool } from "pg";

export class PGWordRepository implements WordRepository {
  private client: Pool;

  constructor(client: Pool) {
    this.client = client;
  }

  create(w: Word): Promise<Word> {
    throw new Error("Method not implemented.");
  }
}
