import { PubSub } from "@google-cloud/pubsub";
import { UserService } from "../service/user-service";
import { DecodedMessage, PubSubListener } from "./pub-sub-listener";

interface UserUpdatesData {
  uid: string;
  email: string;
}

interface UserUpdatesListenerOptions {
  userService: UserService;
  pubSub: PubSub;
}

export class UserUpdatesListener extends PubSubListener<UserUpdatesData> {
  readonly topicName = "user-updates";

  private userService: UserService;

  constructor(options: UserUpdatesListenerOptions) {
    super(options.pubSub);
    this.userService = options.userService;
  }

  async onMessage(msg: DecodedMessage<UserUpdatesData>): Promise<void> {
    try {
      await this.userService.createOrUpdateUser({
        id: msg.data.uid,
        email: msg.data.email,
      });
    } catch (err) {
      console.error(
        `Error creating or updating user for userId: ${msg.data.uid}`,
        err
      );
      msg.nack();
    }
    msg.ack();
  }
}
