import { UserRepository } from "./interfaces";

export class UserService {
  private wr: UserRepository;

  constructor(r: UserRepository) {
    this.wr = r;
  }
}
