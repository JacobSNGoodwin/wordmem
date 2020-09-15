import {
  CreateSubscriptionOptions,
  Message,
  PubSub,
  Subscription,
} from "@google-cloud/pubsub";
import { type } from "os";
import { InternalError } from "../errors/internal-error";

export interface DecodedMessage<T> {
  attributes: {
    [type: string]: string;
  };
  data: T;
}

export abstract class PubSubListener<T> {
  abstract topicName: string;
  abstract onMessage(msg: DecodedMessage<T>): void;

  protected pubSubClinet: PubSub;
  private _subscriptionName?: string;
  subscription?: Subscription;

  constructor(pubSubClient: PubSub) {
    this.pubSubClinet = pubSubClient;
  }

  async init(subscriptionName: string, options?: CreateSubscriptionOptions) {
    this._subscriptionName = subscriptionName;
    const subscriptionResp = await this.pubSubClinet
      .topic(this.topicName)
      .createSubscription(this._subscriptionName, options);
    this.subscription = subscriptionResp[0];
  }

  listen() {
    if (!this.subscription) {
      throw new InternalError(
        "You must initialize a subscription for this process"
      );
    }

    this.subscription.on("message", (rawMsg: Message) => {
      console.log(`${this._subscriptionName} received a message`);

      let parsedData: T;

      try {
        parsedData = JSON.parse(rawMsg.data.toString()) as T;
      } catch (err) {
        console.log("Error receiving and parsing incoming data: ", err);
        throw new InternalError("Unable to receive incoming message");
      }

      this.onMessage({
        attributes: rawMsg.attributes,
        data: parsedData,
      });
    });
  }
}
