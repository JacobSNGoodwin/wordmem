import { v4 } from "uuid";
import { WordRepository } from "./interfaces";
import { Word } from "../model/word";

interface WordData {
  word: string;
  refUrl?: string;
  emailReminder?: boolean;
}

interface UserData {
  id: string;
  email: string;
}

export class WordService {
  private r: WordRepository;

  constructor(r: WordRepository) {
    this.r = r;
  }

  async addWord(w: WordData, u: UserData): Promise<Word> {
    const id = v4();
    const createdWord = this.r.create({
      id,
      word: w.word,
      refUrl: w.refUrl ?? "",
      emailReminder: w.emailReminder ?? false,
      email: u.email,
      userId: u.id,
      startDate: new Date(),
    });

    return createdWord;
  }
}
