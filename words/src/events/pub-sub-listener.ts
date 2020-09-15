import {
  CreateSubscriptionOptions,
  Message,
  PubSub,
  Subscription,
} from "@google-cloud/pubsub";
import { InternalError } from "../errors/internal-error";

export interface DecodedMessage<T> {
  type: string;
  data: T;
  ack(): void;
  nack(): void;
}

export abstract class PubSubListener<T> {
  abstract topicName: string;
  abstract onMessage(msg: DecodedMessage<T>): void; // needs to receive message for acking

  protected pubSubClinet: PubSub;
  private _subscriptionName?: string;
  subscription?: Subscription;

  constructor(pubSubClient: PubSub) {
    this.pubSubClinet = pubSubClient;
  }

  // init initizes s subscription
  // it checks if the desired subscription exists, and creates
  // it otherwise
  async init(subscriptionName: string, options?: CreateSubscriptionOptions) {
    this._subscriptionName = subscriptionName;

    // first array element is boolean, why google did this. who the hell knows?
    const [exists] = await this.pubSubClinet
      .subscription(subscriptionName)
      .exists();

    if (exists) {
      this.subscription = this.pubSubClinet.subscription(
        this._subscriptionName
      );
    } else {
      const [subscription] = await this.pubSubClinet
        .topic(this.topicName)
        .createSubscription(this._subscriptionName, options);
      this.subscription = subscription;
    }
  }

  listen() {
    if (!this.subscription) {
      throw new InternalError(
        "You must initialize a subscription for this process"
      );
    }

    this.subscription.on("message", (rawMsg: Message) => {
      let parsedData: T;

      try {
        parsedData = JSON.parse(rawMsg.data.toString()) as T;
      } catch (err) {
        console.log("Error receiving and parsing incoming data: ", err);
        throw new InternalError("Unable to receive incoming message");
      }

      this.onMessage({
        type: rawMsg.attributes.type,
        data: parsedData,
        ack: () => rawMsg.ack(), // passing function reference no worky
        nack: () => rawMsg.nack(),
      });
    });
  }
}
