import express, { json } from "express";

import { authUser } from "./middleware/auth-user";
import { errorHandler } from "./middleware/error-handler";
import { wordRouter } from "./handler/routes";

const app = express();

app.use(json());
app.use(authUser);

app.use("/api/words", wordRouter);

app.use(errorHandler);

export default app;
