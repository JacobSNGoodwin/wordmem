import { DataSources } from "../data";
import { WordService } from "../service/word-service";
import { PGWordRepository } from "../repository/pg-word-repository";

export interface Services {
  wordService: WordService;
}

class ServiceContainer {
  private _services?: Services;

  init(dataSources: DataSources) {
    console.log("Initializing services");
    const wordRepository = new PGWordRepository(dataSources.db);
    const wordService = new WordService(wordRepository);

    this._services = {
      wordService,
    };
  }

  get services(): Services {
    if (!this._services) {
      throw new Error(
        "Cannot access services before instantiating with init method"
      );
    }

    return this._services;
  }
}

// makes this accessible at top level after initializing
export const serviceContainer = new ServiceContainer();
