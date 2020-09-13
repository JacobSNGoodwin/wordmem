import { CustomError } from "./custom-error";

export class ConflictError extends CustomError {
  statusCode = 409;

  constructor(public field: string) {
    super(`${field} already exists`);

    Object.setPrototypeOf(this, new.target.prototype);
  }

  serializeErrors(): { message: string; field?: string | undefined }[] {
    return [
      {
        message: "The resource already exists",
        field: this.field,
      },
    ];
  }
}
