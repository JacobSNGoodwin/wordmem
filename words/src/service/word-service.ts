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
  word: string;
  definition: string;
  refUrl: string;
  emailReminder: boolean;
  startDate: Date;
  email: string;
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
      email: d.email,
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

  async updateWord(wordId: string, d: UpdateData): Promise<Word> {
    const updatedWord = this.wr.update({
      id: wordId,
      word: d.word,
      definition: d.definition,
      refUrl: d.refUrl,
      emailReminder: d.emailReminder,
      userId: "", // will not be changed
      startDate: d.startDate,
      email: d.email,
    });

    return updatedWord;
  }
}
