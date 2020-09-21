import { User } from "../model/user";
import { Word } from "../model/word";

export interface WordListResponse {
  words: Word[];
  count: number;
}

// Could try responses with algebraic data types in the future,
// but really our errors "expand" the layers, so I'm not terribly
// concerned.
export interface WordRepository {
  create(w: Word): Promise<Word>;
  getByUser(options: {
    uid: string;
    limit: number;
    offset: number;
  }): Promise<WordListResponse>;
  getFiboByUser(options: {
    uid: string;
    limit: number;
    offset: number;
  }): Promise<WordListResponse>;
  deleteByIds(wordId: string[]): Promise<string[]>;
  update(w: Word): Promise<Word>;
}

export interface UserRepository {
  upsert(u: User): Promise<User>;
}
