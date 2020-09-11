import { CustomError } from "./custom-error";

export class NotAuthorizedError extends CustomError {
  statusCode = 401;

  constructor() {
    super("Not authorized");

    Object.setPrototypeOf(this, new.target.prototype);
  }

  serializeErrors() {
    return [
      {
        message: "Not authorized",
      },
    ];
  }
}
