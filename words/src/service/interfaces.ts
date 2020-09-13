import { Word } from "../model/word";

// Could try responses with algebraic data types in the future,
// but really our errors "expand" the layers, so I'm not terribly
// concerned.
export interface WordRepository {
  create(w: Word): Promise<Word>;
  getByUser(uid: string): Promise<Word[]>;
  deleteByIds(wordId: string[]): Promise<string[]>;
  update(w: Word): Promise<Word>;
}
