import { CustomError } from "../errors/custom-error";
import { Word } from "../model/word";

export interface WordRepository {
  create(w: Word): Promise<Word | CustomError>;
}
