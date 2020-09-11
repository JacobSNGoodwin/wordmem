import { Word } from "../model/word";

export interface IWordService {
  getWords(): Promise<Word>;
}
