import { CustomError } from "./custom-error";

export class Internal extends CustomError {
  statusCode = 500;

  constructor(public reason: string = "Unknown error") {
    super(reason);

    Object.setPrototypeOf(this, new.target.prototype);
  }

  serializeErrors() {
    return [
      {
        message: this.reason,
      },
    ];
  }
}
