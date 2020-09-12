import express, { json, Express } from "express";

import { authUser } from "./middleware/auth-user";

import { Services } from "./injection";
import { createWordRouter } from "./handler/word-router";
import { errorHandler } from "./middleware/error-handler";

// we use get app, otherwise our dependency injection
// will no be ready as app is imported at the top of the file
const createApp = (services: Services): Express => {
  const app = express();

  app.use(json());
  app.use(authUser);

  app.use("/api/words", createWordRouter(services.wordService));

  app.use(errorHandler);

  return app;
};

export default createApp;
