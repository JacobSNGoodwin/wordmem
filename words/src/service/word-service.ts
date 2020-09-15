import { v4 } from "uuid";
import { WordRepository } from "./interfaces";
import { Word } from "../model/word";

interface CreateData {
  word: string;
  definition: string;
  refUrl?: string;
  emailReminder?: boolean;
  startDate?: Date;
  uid: string;
  email: string;
}

interface UpdateData {
  id: string;
  word: string;
  uid: string;
  definition: string;
  refUrl: string;
  emailReminder: boolean;
  startDate: Date;
}

export class WordService {
  private wr: WordRepository;

  constructor(r: WordRepository) {
    this.wr = r;
  }

  async addWord(d: CreateData): Promise<Word> {
    const id = v4();
    const createdWord = this.wr.create({
      id,
      word: d.word,
      definition: d.definition,
      refUrl: d.refUrl ?? "",
      emailReminder: d.emailReminder ?? false,
      userId: d.uid,
      startDate: new Date(),
    });

    return createdWord;
  }

  async getWords(userId: string): Promise<Word[]> {
    const words = await this.wr.getByUser(userId);

    return words;
  }

  async deleteWords(wordIds: string[]): Promise<string[]> {
    const deletedIds = await this.wr.deleteByIds(wordIds);

    return deletedIds;
  }

  async updateWord(d: UpdateData): Promise<Word> {
    const updatedWord = this.wr.update({
      id: d.id,
      word: d.word,
      definition: d.definition,
      refUrl: d.refUrl,
      emailReminder: d.emailReminder,
      userId: d.uid, // user for making sure you can't change other users words
      startDate: d.startDate,
    });

    return updatedWord;
  }
}
