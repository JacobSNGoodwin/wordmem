import express, {
  json,
  Express,
  Request,
  Response,
  NextFunction,
} from "express";

import { authUser } from "./middleware/auth-user";

import { createWordRouter } from "./handler/word-router";
import { errorHandler } from "./middleware/error-handler";
import { NotFoundError } from "./errors/not-found-error";

// we use get app, otherwise our dependency injection
// will no be ready as app is imported at the top of the file
const createApp = (): Express => {
  const app = express();

  app.use(json());
  app.use(authUser);

  app.use("/api/words", createWordRouter());

  // if using async function, must pass error to next
  // we probably can omit asyn and just throw the error
  app.all("/api/words/*", () => {
    throw new NotFoundError();
  });
  app.use(errorHandler);

  return app;
};

export default createApp;
