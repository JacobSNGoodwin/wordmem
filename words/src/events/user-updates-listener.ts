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

  onMessage(msg: DecodedMessage<UserUpdatesData>): void {
    console.log("Decoded message:");

    console.log(msg.type);
    console.log(msg.data);

    msg.ack();
  }
}
