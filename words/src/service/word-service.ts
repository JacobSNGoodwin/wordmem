import { WordRepository } from "./interfaces";
import { CustomError } from "../errors/custom-error";
import { Word } from "../model/word";
import { NotAuthorizedError } from "../errors/not-authorized-error";

interface WordInput {
  word: string;
  refUrl?: string;
  emailReminder: boolean;
}

interface UserInput {
  id: string;
  email: string;
}

export class WordService {
  private r: WordRepository;

  constructor(r: WordRepository) {
    this.r = r;
  }

  async addWord(): Promise<Word | CustomError> {
    return new NotAuthorizedError();
  }
}
