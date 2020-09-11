import express from "express";
import app from "../app";
import { authUser } from "../middleware/auth-user";

const appRouter = express.Router();

appRouter.use(authUser);

appRouter.get("/", (req, res) => {
  res.json(req.body);
});

export default appRouter;
