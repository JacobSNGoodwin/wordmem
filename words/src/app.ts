// For creating the express app
import express, { json } from "express";
import appRouter from "./handler/routes";

const app = express();

app.use(json());
app.use("/api/words", appRouter);

export default app;
