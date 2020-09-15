import { UserRepository } from "./interfaces";

interface CreateOrUpdateUserInput {
  id: string;
  email: string;
}

export class UserService {
  private wr: UserRepository;

  constructor(r: UserRepository) {
    this.wr = r;
  }

  async createOrUpdateUser(input: CreateOrUpdateUserInput) {
    const createdUser = this.wr.upsert({
      id: input.id,
      email: input.email,
    });

    return createdUser;
  }
}
