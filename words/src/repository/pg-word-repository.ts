import { WordRepository } from "../service/interfaces";
import { Word } from "../model/word";
import { CustomError } from "../errors/custom-error";
import { Pool } from "pg";
import { InternalError } from "../errors/internal-error";
import { create } from "domain";

export class PGWordRepository implements WordRepository {
  private client: Pool;

  constructor(client: Pool) {
    this.client = client;
  }

  async create(w: Word): Promise<Word> {
    const text =
      "INSERT INTO words (id, userId, email, word, ref_url, email_reminder, start_date) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *";
    const values = [
      w.id,
      w.userId,
      w.email,
      w.word,
      w.refUrl,
      w.emailReminder,
      w.startDate,
    ];

    try {
      const queryRes = await this.client.query({
        text,
        values,
      });

      const createdWord = queryRes.rows[0];

      return {
        id: createdWord.id,
        email: createdWord.email,
        emailReminder: createdWord.email_reminder,
        refUrl: createdWord.ref_url,
        startDate: createdWord.start_date,
        userId: createdWord.userid,
        word: createdWord.word,
      };
    } catch (e) {
      console.debug("Error inserting word into database: ", e);
      throw new InternalError();
    }
  }
}
