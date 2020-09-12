import { Word } from "../model/word";

// Could try responses with algebraic data types in the future,
// but really our errors "expand" the layers, so I'm not terribly
// concerned.
export interface WordRepository {
  create(w: Word): Promise<Word>;
}
