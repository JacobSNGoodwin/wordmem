import express, { json, Express } from "express";

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

  app.all("*", async () => {
    throw new NotFoundError();
  });
  app.use(errorHandler);

  return app;
};

export default createApp;
