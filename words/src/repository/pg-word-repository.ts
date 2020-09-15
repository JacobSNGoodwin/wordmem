import { WordRepository } from "../service/interfaces";
import { Word } from "../model/word";
import { Pool } from "pg";
import { InternalError } from "../errors/internal-error";

export class PGWordRepository implements WordRepository {
  private client: Pool;

  constructor(client: Pool) {
    this.client = client;
  }

  async create(w: Word): Promise<Word> {
    const text = `
        INSERT INTO words (id, userId, word, definition, ref_url, email_reminder, start_date) 
        VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *
      `;
    const values = [
      w.id,
      w.userId,
      w.word,
      w.definition,
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

      return wordFromData(createdWord);
    } catch (e) {
      console.debug("Error inserting word into database: ", e);
      throw new InternalError();
    }
  }

  async getByUser(uid: string): Promise<Word[]> {
    const text = `
      SELECT * FROM words 
      WHERE userid=$1 
      ORDER BY lower(word);
    `;
    const values = [uid];

    try {
      const queryRes = await this.client.query({
        text,
        values,
      });

      const fetchedWords = queryRes.rows;

      return fetchedWords.map((word) => wordFromData(word));
    } catch (e) {
      console.debug("Error retrieving words for user: ", e);
      throw new InternalError();
    }
  }

  // returns list of deleted words
  async deleteByIds(wordIds: string[]): Promise<string[]> {
    // generate $1, $2, ..., $len(wordIds)
    const params = wordIds.map((_, index) => `\$${index + 1}`).join(",");

    // delete words with wordIds and return their ids only
    const text = `
      DELETE FROM words
      WHERE id IN (${params})
      RETURNING (id);
    `;

    try {
      const queryRes = await this.client.query({
        text,
        values: wordIds,
      });

      const deletedWords = queryRes.rows;

      console.debug("Deleted words: ", deletedWords);

      return deletedWords.map((dataObj) => dataObj.id);
    } catch (e) {
      console.debug("Error deleting words from database: ", e);
      throw new InternalError();
    }
  }

  async update(w: Word): Promise<Word> {
    // current word does not allow for optional parameters, so we must receive the whole shebang of a word
    // no need for coalescing if value is not present
    const text = `
        UPDATE words 
        SET word=$1,
        SET definition=$2,
        SET ref_url=$3,
        SET email_reminder=$4,
        SET start_date=$5
        WHERE id=$6
        RETURNING *;
      `;
    const values = [
      w.word,
      w.definition,
      w.refUrl,
      w.emailReminder,
      w.startDate,
      w.id,
    ];

    try {
      const queryRes = await this.client.query({
        text,
        values,
      });

      const updatedWord = queryRes.rows[0];

      return wordFromData(updatedWord);
    } catch (e) {
      console.debug("Error updating word in database: ", e);
      throw new InternalError();
    }
  }
}

const wordFromData = (dataObj: any): Word => ({
  id: dataObj.id,
  emailReminder: dataObj.email_reminder,
  refUrl: dataObj.ref_url,
  startDate: dataObj.start_date,
  userId: dataObj.user_id,
  word: dataObj.word,
  definition: dataObj.definition,
});
