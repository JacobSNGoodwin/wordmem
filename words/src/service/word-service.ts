import { WordRepository } from "./interfaces";
import { CustomError } from "../errors/custom-error";
import { Word } from "../model/word";
import { NotAuthorizedError } from "../errors/not-authorized-error";
import { nextTick } from "process";

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

  async addWord(word: WordInput, user: UserInput): Promise<Word> {
    return {
      id: "abc123",
      email: "bob@bob.com",
      emailReminder: false,
      userId: "123",
      word: "A word",
    };
  }
}
